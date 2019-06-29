package logic

import (
	"github.com/mkzz115/z-service.git/common/log"
	"github.com/mkzz115/z-service.git/common/pub/idl/gen-go/account"
	"github.com/mkzz115/z-service.git/common/serve"
)

var dps serve.DependsServer

func NewAccountHandle() *AccountHandle {
	return &AccountHandle{}
}

func InitHandle(dp serve.DependsServer) error {
	dps = dp
	return nil
}

type AccountHandle struct {
	//dp serve.DependsServer
}

func (AccountHandle) UserLogIn(req *account.LogInReq) (r *account.LogInRes, err error) {
	log.Info("AccountHandle.UserLogIn is called")
	return nil, nil
}

func (AccountHandle) UserLogOut(req *account.LogOutReq) (r *account.LogOutRes, err error) {
	log.Info("AccountHandle.UserLogOut is called")
	return nil, nil
}

func (AccountHandle) UserRegister(req *account.RegisterReq) (r *account.RegisterRes, err error) {
	log.Info("AccountHandle.UserRegister is called")
	return nil, nil
}

func (AccountHandle) UserChangePw(req *account.ChangePwReq) (r *account.ChangePwRes, err error) {
	log.Info("AccountHandle.UserChangePw is called")
	return nil, nil
}
