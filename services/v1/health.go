package v1

import (
	"ir-web/internal/wrapper"
	mlog "ir-web/internal/wrapper/mlog"
)

func HeartBeat(ctx *wrapper.Context, reqBody interface{}) error {
	wrapper.SendApiOKResponse(ctx, nil, "health")
	mlog.Info("Apiwrapper ok")
	return nil
}
