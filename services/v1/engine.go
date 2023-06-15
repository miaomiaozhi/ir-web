package v1

import (
	"ir-web/internal/wrapper"
	mlog "ir-web/internal/wrapper/mlog"
	engine "ir-web/middleware"
	v1_req "ir-web/models/protoreq/v1"
	v1_resp "ir-web/models/protoresp/v1"
)

func Query(ctx *wrapper.Context, reqBody interface{}) error {
	mlog.Info("handle query now")
	req := reqBody.(v1_req.EngineRequest)
	engine.GetEngine().QueryIndexListByToken(req.Token)

	resp := v1_resp.EngineResponse{
		Title: []string{"hello"},
	}
	wrapper.SendApiOKResponse(ctx, resp, "查询成功")
	return nil
}
