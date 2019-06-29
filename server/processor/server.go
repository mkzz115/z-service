package processor

import (
	"github.com/mkzz115/z-service.git/common/pub/idl/gen-go/hello"
	"github.com/mkzz115/z-service.git/common/serve"
)

//
//func NewThriftServer(cfg *confutil.Config) serve.Serve {
//	return &ThriftServer{
//		conf: cfg,
//	}
//}
//
//type ThriftServer struct {
//	conf *confutil.Config
//}
//
//func (t *ThriftServer) Init() error {
//	log.Info("ThriftServer.Init called")
//	return nil
//}
//
//func (t *ThriftServer) Start() error {
//	log.Info("ThriftServer.Init called")
//	handler := NewHelloServe()
//	process := hello.NewHelloServiceProcessor(handler)
//	transport := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
//	protocol := thrift.NewTBinaryProtocolFactoryDefault()
//	sock, err := thrift.NewTServerSocket(t.conf.Server.Addr)
//	if err != nil {
//		log.Error("ThriftServer.Start create socket error[%s], addr[%s]",
//			err.Error(), t.conf.Server.Addr)
//		return err
//	}
//	server := thrift.NewTSimpleServer4(process, sock, transport, protocol)
//	err = server.Listen()
//	if err != nil {
//		log.Error("ThriftServer.Start listen error[%s], addr[%s]",
//			err.Error(), t.conf.Server.Addr)
//		return err
//	}
//	err = server.Serve()
//	if err != nil {
//		log.Error("serve error[%s]", err.Error())
//	}
//	return nil
//}

func NewThriftProcess() serve.Processor {
	return &thriftProcess{}
}

type thriftProcess struct{}

func (s *thriftProcess) Init() error {
	return nil
}

func (s *thriftProcess) Driver() (string, interface{}) {
	handler := NewHelloServe()
	process := hello.NewHelloServiceProcessor(handler)
	return serve.PROCESSOR_THRIFT, process
}
