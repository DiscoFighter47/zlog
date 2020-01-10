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
func Info(msg string) {
	logger.Info().Msg(msg)
}

// Infof ...
func Infof(format string, v ...interface{}) {
	logger.Info().Msgf(format, v...)
}

// Debug ...
func Debug(msg string) {
	logger.Debug().Msg(msg)
}

// Debugf ...
func Debugf(format string, v ...interface{}) {
	logger.Debug().Msgf(format, v...)
}

// Warn ...
func Warn(msg string) {
	logger.Warn().Msg(msg)
}

// Warnf ...
func Warnf(format string, v ...interface{}) {
	logger.Warn().Msgf(format, v...)
}

// Error ...
func Error(msg string) {
	logger.Error().Msg(msg)
}

// Errorf ...
func Errorf(format string, v ...interface{}) {
	logger.Error().Msgf(format, v...)
}

// Panic ...
func Panic(msg string) {
	logger.Panic().Msg(msg)
}

// Panicf ...
func Panicf(format string, v ...interface{}) {
	logger.Panic().Msgf(format, v...)
}

// Fatal ...
func Fatal(msg string) {
	logger.Fatal().Msg(msg)
}

// Fatalf ...
func Fatalf(format string, v ...interface{}) {
	logger.Fatal().Msgf(format, v...)
}
