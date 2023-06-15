package main

import (
	"ir-web/conf"
	mlog "ir-web/internal/wrapper/mlog"
	"ir-web/middleware"
	"ir-web/runner"
)

func main() {
	config, err := runner.MustReadConfigFromCmdFlags()
	if err != nil {
		panic("init config error: " + err.Error())
	}
	middleware.InitEngineWithConfig(config)
	mlog.InitWithConfig(config)
	conf.InitGlobalConfig(config)
	go runner.InitBackgournTaskWithConfig(config)
	runner.StartWebApp(config)
}
