package processor

import (
	"github.com/mkzz115/z-service.git/account_server/logic"
	"github.com/mkzz115/z-service.git/common/pub/idl/gen-go/account"
	"github.com/mkzz115/z-service.git/common/serve"
)

func NewAccountProcess() serve.Processor {
	return &accountProcess{}
}

type accountProcess struct{}

func (s *accountProcess) Init() error {
	return nil
}

func (s *accountProcess) Driver() (string, interface{}) {
	handler := logic.NewAccountHandle()
	process := account.NewAccountServiceProcessor(handler)
	return serve.PROCESSOR_THRIFT, process
}
