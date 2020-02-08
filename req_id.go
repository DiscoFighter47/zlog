package zlog

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/tylerb/gls"
)

// ReqIDTag ...
const ReqIDTag = "request_id"

// SetReqID ...
func SetReqID() {
	gls.Set(ReqIDTag, uuid.NewV4().String())
}

// GetReqID ...
func GetReqID() string {
	if v := gls.Get(ReqIDTag); v != nil {
		return fmt.Sprintf("%v", v)
	}
	return ""
}
