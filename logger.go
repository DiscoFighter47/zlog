package zlog

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

var logger = zerolog.New(zerolog.ConsoleWriter{
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

// Info ...
var Info = logger.Info().Msg

// Infof ...
var Infof = logger.Info().Msgf

// Debug ...
var Debug = logger.Debug().Msg

// Debugf ...
var Debugf = logger.Debug().Msgf

// Warn ...
var Warn = logger.Warn().Msg

// Warnf ...
var Warnf = logger.Warn().Msgf

// Error ...
var Error = logger.Error().Msg

// Errorf ...
var Errorf = logger.Error().Msgf

// Fatal ...
var Fatal = logger.Fatal().Msg

// Fatalf ...
var Fatalf = logger.Fatal().Msgf

// Panic ...
var Panic = logger.Panic().Msg

// Panicf ...
var Panicf = logger.Panic().Msgf
