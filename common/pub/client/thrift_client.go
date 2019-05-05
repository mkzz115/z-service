package client

import (
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/mkzz115/z-service.git/common/confutil"
)

func NewThriftClient(cfg *confutil.ClientConfig) *ThriftClient {
	return &ThriftClient{
		conf: cfg,
	}
}

type ThriftClient struct {
	conf *confutil.ClientConfig
}

func (t *ThriftClient) Connect() error {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport, err := thrift.NewTSocketTimeout(t.conf.Addr, time.Duration(t.conf.Timeout)*time.Second)
	if err != nil {
		return err
	}
	useTransport := transportFactory.GetTransport(transport)
	err = useTransport.Open()
	defer useTransport.Close()
	if err != nil {
		return nil
	}
}
