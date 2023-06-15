package v1

import (
	v1_controller "ir-web/controllers/v1"
	"ir-web/internal/wrapper"
	mlog "ir-web/internal/wrapper/mlog"
	"net/http"

	"github.com/kataras/iris/v12/core/router"
)

func RegisterEngineRouter(party router.Party) {
	v1 := party.Party("/")
	{
		v1.Handle(http.MethodPost, "/query", wrapper.HandlerNotLogin(v1_controller.EngineController{}.Query))
	}
	mlog.Info("register Engine router success")
}
