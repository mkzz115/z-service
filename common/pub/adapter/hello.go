package adapter

import (
	"time"

	"github.com/mkzz115/z-service.git/common/confutil"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/mkzz115/z-service.git/common/pub/idl/gen-go/hello"
)

var (
	defaultTimeout = time.Duration(time.Second * 5)
	hostPort       = ""
)

func Init(conf *confutil.ClientConfig) error {
	hostPort = conf.Addr
	defaultTimeout = time.Duration(conf.Timeout) * time.Second
	return nil
}

func rpc(timeout time.Duration, fn func(*hello.HelloServiceClient) error) error {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport, err := thrift.NewTSocketTimeout(hostPort, timeout)
	if err != nil {
		return err
	}
	useTransport := transportFactory.GetTransport(transport)
	err = useTransport.Open()
	defer useTransport.Close()
	if err != nil {
		return err
	}
	c := hello.NewHelloServiceClientFactory(useTransport, protocolFactory)
	return fn(c)
}

func GetHelloAdapter(req *hello.HelloReq) (*hello.HelloRes, error) {
	var res *hello.HelloRes

	err := rpc(defaultTimeout, func(c *hello.HelloServiceClient) error {
		r, e := c.GetHello(req)
		res = r
		return e
	})
	return res, err
}
