package processor

import (
	"github.com/mkzz115/z-service.git/common/log"
	"github.com/mkzz115/z-service.git/common/pub/idl/gen-go/hello"
	"github.com/mkzz115/z-service.git/common/serve"
)

var dps serve.DependsServer

func NewHelloServe() *HelloServe {
	return &HelloServe{}
}

func InitHandle(dp serve.DependsServer) error {
	_ = log.Init("InitHandle called")
	dps = dp
	return nil
}

type HelloServe struct {
	dp serve.DependsServer
}

func (h *HelloServe) GetHello(req *hello.HelloReq) (r *hello.HelloRes, err error) {
	log.Info("HelloServe.GetHello called, req.uid[%d]", req.GetUID())
	data := req.GetUID()
	res := hello.NewHelloRes()
	res.Ack = &data
	var uid int64 = 0
	err = dps.DbInstance().Get(&uid, "select uid from user_account_t where account=?", "zz")
	if err != nil {
		log.Info("select db err:%s", err.Error())
		return res, err
	}
	log.Info("account:%d", uid)
	res.Ack = &uid
	return res, nil
}
