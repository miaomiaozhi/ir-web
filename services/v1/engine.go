package v1

import (
	"ir-web/internal/wrapper"
	mlog "ir-web/internal/wrapper/mlog"
	engine "ir-web/middleware"
	v1_req "ir-web/models/protoreq/v1"
	v1_resp "ir-web/models/protoresp/v1"
	"time"
)

func Query(ctx *wrapper.Context, reqBody interface{}) error {
	mlog.Info("handle query now")
	req := reqBody.(*v1_req.EngineRequest)
	mlog.Info("req Token is", req.Token)
	ids := engine.GetEngine().QueryIndexListByToken(req.Token)

	start := time.Now()
	if len(ids) == 0 {
		ids = engine.GetEngine().FuzzyQueryIndexListByToken(req.Token)
	}

	title := make([]string, 0, len(ids))
	urls := make([]string, 0, len(ids))
	for _, v := range ids {
		title = append(title, engine.GetEngine().Title[v])
		urls = append(urls, engine.GetEngine().Urls[v])
	}

	used := time.Since(start)
	resp := v1_resp.EngineResponse{
		Title: title,
		Urls:  urls,
		Time:  int32(used.Milliseconds()),
	}

	mlog.Info("used time", resp.Time, "ms")
	mlog.Info("token is", req.Token)
	mlog.Info("query ids", ids)
	wrapper.SendApiOKResponse(ctx, resp, "查询成功")
	return nil
}
