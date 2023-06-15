package v1

import (
	"ir-web/internal/wrapper"
	v1 "ir-web/services/v1"
)

type HealthController struct {
}

func (HealthController) HeartBeat(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, v1.HeartBeat, true, nil, nil)
}
