package runner

import (
	conf "ir-web/conf"
)

type Runner struct {
	Conf *conf.IrConfig
}

func New(conf *conf.IrConfig) *Runner {
	return &Runner{
		Conf: conf,
	}
}

// // 启动 app
// func (r *Runner) StartWebApp() {
// 	web.New().Run(r.Conf)
// }
