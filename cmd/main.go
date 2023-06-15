package main

import (
	"ir-web/conf"
	wrapper "ir-web/internal/wrapper/mlog"
	"ir-web/runner"
)

func main() {
	config, err := runner.MustReadConfigFromCmdFlags()
	if err != nil {
		panic("init config error: " + err.Error())
	}
	conf.InitGlobalConfig(config)
	wrapper.InitWithConfig(config)
}
