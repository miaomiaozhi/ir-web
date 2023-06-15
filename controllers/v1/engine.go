package v1

import (
	"ir-web/internal/wrapper"
	v1_req "ir-web/models/protoreq/v1"
	v1 "ir-web/services/v1"
)

type EngineController struct {
}

func (EngineController) Query(ctx *wrapper.Context) {
	wrapper.ApiWrapper(ctx, v1.Query, true, &v1_req.EngineRequest{}, nil)
}
