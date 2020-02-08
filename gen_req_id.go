package zlog

import (
	"net/http"

	"github.com/tylerb/gls"
)

// GenReqID ...
func GenReqID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		SetReqID()
		defer gls.Cleanup()
		next.ServeHTTP(w, r)
	})
}
