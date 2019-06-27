package logic

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mkzz115/z-service.git/common/confutil"
	"github.com/mkzz115/z-service.git/common/pub/adapter"
	"github.com/mkzz115/z-service.git/common/pub/idl/gen-go/hello"

	"github.com/mkzz115/z-service.git/common/log"
)

func NewLogicHandle(cfg *confutil.Config) *logicInfo {
	return &logicInfo{
		cfg: cfg,
	}
}

type logicInfo struct {
	cfg *confutil.Config
}

func (m *logicInfo) Hello(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Info("hello is called, data[%v]", string(buf))
	req := &hello.HelloReq{}
	err = json.Unmarshal(buf, req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := adapter.GetHelloAdapter(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resData, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resData)
}
