package processor

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/mkzz115/z-service.git/api_service/config"
	"github.com/mkzz115/z-service.git/api_service/logic"
	"github.com/mkzz115/z-service.git/common/httputil"
	"github.com/mkzz115/z-service.git/common/log"
	"github.com/mkzz115/z-service.git/common/serve"
)

type ApiServer struct {
	conf   *config.Config
	router *httprouter.Router
}

func NewApiServer(cfg *config.Config) serve.Serve {
	return &ApiServer{
		conf: cfg,
	}
}
func (s *ApiServer) Init() error {
	log.Info("ApiServer.Init called")
	// install driver
	factory := httputil.HttpWrapper(
		httputil.NewAccess(),
	)
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/", factory(logic.Hello))
	s.router = router
	return nil
}

func (s *ApiServer) Start() error {
	serv := &http.Server{
		Addr:              ":10230",
		Handler:           s.router,
		ReadHeaderTimeout: 5 * time.Second,
	}
	err := serv.ListenAndServe()
	if err != nil {
		log.Error("ApiServer.Start error[%s]", err.Error())
	}
	return nil
}
