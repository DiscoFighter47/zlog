package zlog

import (
	"net/http"
	"testing"
)

func testHandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		Logger.Info().Msg("Hello World!")
	}
	return http.HandlerFunc(fn)
}

func TestGatekeeper(t *testing.T) {
	svr := GenReqID(testHandler())
	svr.ServeHTTP(nil, nil)
}
