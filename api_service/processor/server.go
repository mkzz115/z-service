package processor

import (
	"github.com/mkzz115/z-service.git/common/confutil"
	"github.com/mkzz115/z-service.git/common/pub/adapter"
	"github.com/mkzz115/z-service.git/common/serve"
)

//type ApiServer struct {
//	conf   *confutil.Config
//	router *httprouter.Router
//}
//
//func NewApiServer(cfg *confutil.Config) serve.Serve {
//	return &ApiServer{
//		conf: cfg,
//	}
//}
//func (s *ApiServer) Init() error {
//	log.Info("ApiServer.Init called")
//	// install driver
//	// init adapter
//	err := adapter.Init(&s.conf.Client)
//	if err != nil {
//		log.Error("init adapter err:%s", err.Error())
//		return err
//	}
//	s.router = NewApiRouter(s.conf)
//	return nil
//}
//
//func (s *ApiServer) Start() error {
//	serv := &http.Server{
//		Addr:              s.conf.Server.Addr,
//		Handler:           s.router,
//		ReadHeaderTimeout: 5 * time.Second,
//	}
//	err := serv.ListenAndServe()
//	if err != nil {
//		log.Error("ApiServer.Start error[%s]", err.Error())
//	}
//	return nil
//}

func NewHttpProcess(cfg *confutil.Config) serve.Processor {
	return &httpProcess{
		cfg: cfg,
	}
}

type httpProcess struct {
	cfg *confutil.Config
}

func (h *httpProcess) Init() error {
	// init adapter
	err := adapter.Init(&h.cfg.Client)
	return err
}

func (h *httpProcess) Driver() (string, interface{}) {
	r := NewApiRouter(h.cfg)
	return serve.PROCESSOR_HTTP, r
}
