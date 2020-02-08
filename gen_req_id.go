package zlog

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
	"github.com/tylerb/gls"
)

// ReqIDTag ...
const ReqIDTag = "request_id"

// GenReqID ...
func GenReqID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gls.Set(ReqIDTag, uuid.NewV4().String())
		defer gls.Cleanup()
		next.ServeHTTP(w, r)
	})
}
