package wrapper

import (
	mlog "ir-web/internal/wrapper/mlog"
	model "ir-web/models/protoresp"

	"github.com/kataras/iris/v12"
)

func makeResponse(ctx *Context, data interface{}, msg string, statusCode int32) {
	resp := model.BaseResponse{}

	if msg == "" || msg == "success" {
		msg = "success"
		statusCode = 200
	} else {
		statusCode = int32(ctx.GetStatusCode())
	}
	resp.StatusCode = statusCode
	resp.StatusMsg = msg
	resp.Data = data

	err := ctx.JSON(resp)
	if err != nil {
		mlog.Error("make response error " + err.Error())
	}
}

// 响应成功 200
func SendApiOKResponse(ctx *Context, data interface{}, msg string) {
	ctx.StatusCode(iris.StatusOK)
	makeResponse(ctx, data, msg, iris.StatusOK)
}

// 参数错误 400
func SendApiBadRequestResponse(ctx *Context, data interface{}, msg string) {
	ctx.StatusCode(iris.StatusBadRequest)
	makeResponse(ctx, data, msg, iris.StatusBadRequest)
}

// 认证错误 403
func SendApiForbiddenResponse(ctx *Context, data interface{}, msg string) {
	ctx.StatusCode(iris.StatusForbidden)
	makeResponse(ctx, data, msg, iris.StatusBadRequest)
}

// 内部错误 500
func SendApiErrorResponse(ctx *Context, data interface{}, msg string) {
	ctx.StatusCode(iris.StatusInternalServerError)
	makeResponse(ctx, data, msg, iris.StatusInternalServerError)
}

// 认证失败 401，需要重新登录
func SendApiUnAuthResponse(ctx *Context, data interface{}, msg string) {
	ctx.StatusCode(iris.StatusUnauthorized)
	makeResponse(ctx, data, msg, iris.StatusUnauthorized)
}
