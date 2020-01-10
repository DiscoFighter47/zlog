package zlog_test

import (
	"net/http"
	"testing"

	"github.com/DiscoFighter47/zlog"
)

func testHandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		zlog.Info("Hello World!")
	}
	return http.HandlerFunc(fn)
}

func TestGatekeeper(t *testing.T) {
	svr := zlog.GenReqID(testHandler())
	svr.ServeHTTP(nil, nil)
}
