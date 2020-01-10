package zlog

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/tylerb/gls"
)

type ridHook struct{}

func (h ridHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if level != zerolog.NoLevel {
		e.Str("request_id", fmt.Sprintf("%v", gls.Get("request_id")))
	}
}
