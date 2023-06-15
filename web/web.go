package web

import (
	"ir-web/conf"
	mlog "ir-web/internal/wrapper/mlog"
	"ir-web/routers"
	"strconv"

	"github.com/kataras/iris/v12"
)

type IrWeb struct {
}

func newApp() *iris.Application {
	return iris.New()
}

func New() *IrWeb {
	return &IrWeb{}
}

func (*IrWeb) Run(config *conf.IrConfig) {
	app := newApp()
	routers.IrisRouter{}.InitApp(app)

	// 注册路由，当访问 /index 路径时，返回 index.html 页面
	app.Get("/index", func(ctx iris.Context) {
		ctx.ServeFile("./view/index.html")
		// ctx.JSON(iris.Map{
		// 	"status": "ok",
		// })
	})

	port := config.GetInt("web.port", 8080)
	mlog.Info("app listening port:", port)
	app.Listen(":" + strconv.Itoa(int(port)))
}
