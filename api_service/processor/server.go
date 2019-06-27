package processor

import (
	"net/http"
	"time"

	"github.com/mkzz115/z-service.git/common/pub/adapter"

	"github.com/julienschmidt/httprouter"
	"github.com/mkzz115/z-service.git/common/confutil"
	"github.com/mkzz115/z-service.git/common/log"
	"github.com/mkzz115/z-service.git/common/serve"
)

type ApiServer struct {
	conf   *confutil.Config
	router *httprouter.Router
}

func NewApiServer(cfg *confutil.Config) serve.Serve {
	return &ApiServer{
		conf: cfg,
	}
}
func (s *ApiServer) Init() error {
	log.Info("ApiServer.Init called")
	// install driver
	// init adapter
	err := adapter.Init(&s.conf.Client)
	if err != nil {
		log.Error("init adapter err:%s", err.Error())
		return err
	}
	s.router = NewApiRouter(s.conf)
	return nil
}

func (s *ApiServer) Start() error {
	serv := &http.Server{
		Addr:              s.conf.Server.Addr,
		Handler:           s.router,
		ReadHeaderTimeout: 5 * time.Second,
	}
	err := serv.ListenAndServe()
	if err != nil {
		log.Error("ApiServer.Start error[%s]", err.Error())
	}
	return nil
}
