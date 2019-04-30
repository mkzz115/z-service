package logic

import (
	"io/ioutil"
	"net/http"

	"github.com/mkzz115/z-service.git/common/log"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	log.Info("hello is called, data")
	data, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Info("hello is called, data[%v]", string(data))
	w.Write([]byte("{errno:0}"))
}
