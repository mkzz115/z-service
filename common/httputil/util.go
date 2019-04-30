package httputil

import (
	"net/http"
	"strings"
	"time"

	"github.com/mkzz115/z-service.git/common/log"
)

type HandleDecorator func(http.HandlerFunc) http.HandlerFunc

func HttpWrapper(h ...HandleDecorator) HandleDecorator {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		for i := len(h) - 1; i >= 0; i-- {
			handler = h[i](handler)
		}
		return handler
	}
}

func NewAccess() HandleDecorator {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			begin := time.Now()
			forwarded := r.Header.Get("X-FORWARDED-FOR")
			if len(forwarded) > 0 {
				item := strings.Split(forwarded, ",")
				if len(item) > 0 {
					forwarded = item[0]
				}
			}
			log.Info("_com_request_in,method[%s],forwarded[%s],remote[%s],uri[%d],params[%v],Content-Type[%s],time(us)[%d]",
				r.Method, forwarded, r.RemoteAddr, r.URL.Path, r.URL.RawQuery, begin.UnixNano()/1e6)
			f(w, r)
			log.Info("_com_request_out,laency(ms)[%d]", int64(time.Since(begin)/time.Millisecond))
		}
	}
}
