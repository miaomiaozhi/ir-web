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

	port := config.GetInt("web.port", 8080)
	mlog.Info("app listening port:", port)
	app.Listen(":" + strconv.Itoa(int(port)))
}
