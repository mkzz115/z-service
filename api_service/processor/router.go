package processor

import (
	"net/http"

	"github.com/mkzz115/z-service.git/common/log"
	"github.com/mkzz115/z-service.git/common/serve"

	"github.com/julienschmidt/httprouter"
	"github.com/mkzz115/z-service.git/api_service/logic"
	"github.com/mkzz115/z-service.git/common/confutil"
	"github.com/mkzz115/z-service.git/common/httputil"
)

var dps serve.DependsServer

func InitRouteHandle(dp serve.DependsServer) error {
	_ = log.Init("InitHandle called")
	dps = dp

	return nil
}

func NewApiRouter(cfg *confutil.Config) *httprouter.Router {
	r := httprouter.New()
	processApiRouter(r, cfg)
	return r
}

func processApiRouter(r *httprouter.Router, cfg *confutil.Config) {
	h := logic.NewLogicHandle(cfg)
	factory := httputil.HttpWrapper(
		httputil.NewAccess(),
	)
	r.HandlerFunc(http.MethodPost, "/hello", factory(h.Hello))
}
