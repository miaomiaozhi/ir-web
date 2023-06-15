package wrapper

import (
	"encoding/json"
	logger "ir-web/internal/wrapper/mlog"
	"strings"

	"github.com/kataras/iris/v12"
)

type Context struct {
	iris.Context

	UserToken *AuthResult
}

type AuthResult struct {
	// TODO no auth needed
}

// Acquire 将 iris 的context转化成自定义的context
// 并且做授权验证
func Acquire(original iris.Context, login bool) *Context {
	// TODO add context to context pool
	ctx := &Context{
		Context: original,
	}
	// set the context to the original one in order to have access to iris's implementation.
	if login {
		// TODO: no auth needed
		// ctx.UserToken = GetAuthResult(original)
		// if ctx.UserToken == nil {
		// 	SendApiUnAuthResponse(ctx, nil, "Token 不合法")
		// 	ctx.StopExecution()
		// } else {
		// 	if err != nil {
		// 		SendApiUnAuthResponse(ctx, nil, "Token 不合法")
		// 		ctx.StopExecution()
		// 	} else {
		// 		userId, _ := claim["UserId"].(float64)
		// 		userName, _ := claim["UserName"].(string)
		// 		priority, _ := claim["Priority"].(float64)

		// 		// 不合法
		// 		if int64(userId) != ctx.UserToken.UserID ||
		// 			userName != ctx.UserToken.UserName ||
		// 			int64(priority) != ctx.UserToken.Priority {
		// 			SendApiUnAuthResponse(ctx, nil, "Token 不合法")
		// 			ctx.StopExecution()
		// 		}
		// 	}
		// }
	}
	return ctx
}

func GetAuthResult(ctx iris.Context) *AuthResult {
	user := ctx.GetHeader("User")
	if user == "" {
		logger.Error("header user invalid")
		return nil
	}
	user = strings.Replace(user, "\\", "", -1)
	logger.Info("header:", user)

	// 得到授权信息
	authInfo := &AuthResult{}
	err := json.Unmarshal([]byte(user), &authInfo)
	if err != nil {
		logger.Error("auth info json unmarshal failed", err.Error())
		return nil
	}
	logger.Debug("user auth token ")
	return authInfo
}

// 处理需要登录请求
func Handler(handler func(*Context)) iris.Handler {
	return func(original iris.Context) {
		ctx := Acquire(original, true)
		if !ctx.IsStopped() { // 请求被终止
			handler(ctx)
		}
		// TODO release
	}
}

// 处理无需登录的请求
func HandlerNotLogin(handle func(*Context)) iris.Handler {
	return func(original iris.Context) {
		ctx := Acquire(original, false)
		if !ctx.IsStopped() { // 请求被终止
			handle(ctx)
		}
		// TODO release
	}
}
