package v1

import (
	"ir-web/internal/wrapper"
)

func HeartBeat(ctx *wrapper.Context, reqBody interface{}) error {
	wrapper.SendApiOKResponse(ctx, nil, "health")
	return nil
}
