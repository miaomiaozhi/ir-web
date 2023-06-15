package v1

import (
	v1_controller "ir-web/controllers/v1"
	"ir-web/internal/wrapper"
	mlog "ir-web/internal/wrapper/mlog"
	"net/http"

	"github.com/kataras/iris/v12/core/router"
)

func RegisterHealthRouter(party router.Party) {
	v1 := party.Party("/health")
	{
		v1.Handle(http.MethodGet, "/", wrapper.HandlerNotLogin(v1_controller.HealthController{}.HeartBeat))
	}
	mlog.Info("register health router success")
}
