package zlog

import (
	"fmt"

	"github.com/rs/zerolog"
)

type ridHook struct{}

func (h ridHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if id := GetReqID(); id != "" && level != zerolog.NoLevel {
		e.Str("request_id", fmt.Sprintf("%v", id))
	}
}
