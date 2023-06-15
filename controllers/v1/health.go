package v1

import (
	"ir-web/internal/wrapper"
	v1 "ir-web/services/v1"
	mlog "ir-web/internal/wrapper/mlog"
)

type HealthController struct {
}

func (HealthController) HeartBeat(ctx *wrapper.Context) {
	mlog.Info("health controller wrapper HeartBeat")
	wrapper.ApiWrapper(ctx, v1.HeartBeat, true, nil, nil)
}
