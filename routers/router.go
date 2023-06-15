package routers

import (
	mlog "ir-web/internal/wrapper/mlog"
	v1 "ir-web/routers/api/v1"

	"github.com/kataras/iris/v12"
)

type IrisRouter struct{}

func (IrisRouter) InitApp(app *iris.Application) {
	mlog.Info("init app")
	loadMiddlerware(app)
	appRouter := app.Party("/api/v1")
	{
		// 注册健康检查路由
		v1.RegisterHealthRouter(appRouter)

		// 注册User路由 包括普通用户以及审核用户
		// v1.RegisterUserRouter(appRouter)
	}
	mlog.Info("init app router success")
}

func loadMiddlerware(app *iris.Application) {
	// logger.Info("load middle ware for app")
	//app.Use()
}
