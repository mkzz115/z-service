package processor

import (
	"github.com/mkzz115/z-service.git/common/log"
	"github.com/mkzz115/z-service.git/common/pub/idl/gen-go/hello"
)

func NewHelloServe() *HelloServe {
	return &HelloServe{}
}

type HelloServe struct{}

func (h *HelloServe) GetHello(req *hello.HelloReq) (r *hello.HelloRes, err error) {
	log.Info("HelloServe.GetHello called, req.uid[%d]", req.GetUID())
	data := req.GetUID()
	res := hello.NewHelloRes()
	res.Ack = &data
	return res, nil
}
