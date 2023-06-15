package runner

import (
	conf "ir-web/conf"
	"ir-web/web"
)

type Runner struct {
	Conf *conf.IrConfig
}

func New(conf *conf.IrConfig) *Runner {
	return &Runner{
		Conf: conf,
	}
}

// 启动 app
func StartWebApp(conf *conf.IrConfig) {
	web.New().Run(conf)
}
