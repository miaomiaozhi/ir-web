package main

import (
	"ir-web/conf"
	mlog "ir-web/internal/wrapper/mlog"
	"ir-web/runner"
)

func main() {
	config, err := runner.MustReadConfigFromCmdFlags()
	if err != nil {
		panic("init config error: " + err.Error())
	}
	go runner.InitBackgournTaskWithConfig(config)
	conf.InitGlobalConfig(config)
	mlog.InitWithConfig(config)
	runner.StartWebApp(config)
}
