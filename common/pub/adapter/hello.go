package adapter

import (
	"time"

	"github.com/mkzz115/z-service.git/common/confutil"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/mkzz115/z-service.git/common/pub/idl/gen-go/hello"
)

func GetHelloAdapter(req *hello.HelloReq, conf *confutil.ClientConfig) (*hello.HelloRes, error) {

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport, err := thrift.NewTSocketTimeout(conf.Addr, time.Duration(conf.Timeout)*time.Second)
	if err != nil {
		return nil, err
	}
	useTransport := transportFactory.GetTransport(transport)
	err = useTransport.Open()
	defer useTransport.Close()
	if err != nil {
		return nil, nil
	}
	c := hello.NewHelloServiceClientFactory(useTransport, protocolFactory)
	res, err := c.GetHello(req)
	return res, err
}
