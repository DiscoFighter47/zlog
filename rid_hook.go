package zlog

import (
	"fmt"

	"github.com/tylerb/gls"

	"github.com/rs/zerolog"
)

type ridHook struct{}

func (h ridHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if id := gls.Get(ReqIDTag); id != nil && level != zerolog.NoLevel {
		e.Str("request_id", fmt.Sprintf("%v", id))
	}
}
