package zlog

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// Logger ...
var Logger zerolog.Logger

func init() {
	Logger = zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
		FormatTimestamp: func(i interface{}) string {
			return fmt.Sprintf("%v", i)
		},
		FormatLevel: func(i interface{}) string {
			return fmt.Sprintf("[%-5v]", i)
		},
		FormatFieldName: func(i interface{}) string {
			return fmt.Sprintf("%v:", i)
		},
		FormatFieldValue: func(i interface{}) string {
			return fmt.Sprintf("%v", i)
		},
	}).With().Timestamp().Logger().Hook(ridHook{})
}
