package serve

import (
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/mkzz115/z-service.git/api_service/logic"

	"github.com/mkzz115/z-service.git/common/pub/idl/gen-go/hello"
	"github.com/mkzz115/z-service.git/server/processor"
)

func TestZServe_Start(t *testing.T) {
	go test_thrift(t)
	test_http(t)
}

func test_thrift(t *testing.T) {
	fn := func(server DependsServer) error {
		t.Log("fnnnnnnn")
		return nil
	}
	proc := &thriftServe{}
	zserve := NewZServer("test_thrift", "127.0.0.1:12306")

	zserve.Init(nil, fn, proc)
	zserve.Start()
}

type thriftServe struct{}

func (p *thriftServe) Init() error {
	println("thriftServe.Init")
	return nil
}

func (p *thriftServe) Driver() (string, interface{}) {
	handler := processor.NewHelloServe()
	processor := hello.NewHelloServiceProcessor(handler)
	return PROCESSOR_THRIFT, processor
}

func test_http(t *testing.T) {
	fn := func(server DependsServer) error {
		t.Log("fnnnnnnn")
		return nil
	}
	proc := &httpServe{}
	zserve := NewZServer("test_http", "127.0.0.1:12307")

	zserve.Init(nil, fn, proc)
	zserve.Start()

}

type httpServe struct{}

func (p *httpServe) Init() error {
	println("httpServe.Init")
	return nil
}

func (p *httpServe) Driver() (string, interface{}) {
	r := httprouter.New()
	h := logic.NewLogicHandle(nil)
	r.HandlerFunc(http.MethodPost, "/hello", h.Hello)
	return PROCESSOR_HTTP, r
}
