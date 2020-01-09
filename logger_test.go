package zlog

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tylerb/gls"
)

func TestLogger(t *testing.T) {
	t.Run("info", func(t *testing.T) {
		Logger.Info().Msg("Hello World!")
		gls.Cleanup()
	})

	t.Run("info", func(t *testing.T) {
		gls.Set("request_id", rand.Int63())
		Logger.Info().Msg("Hello World!")
		gls.Cleanup()
	})

	t.Run("debug", func(t *testing.T) {
		gls.Set("request_id", rand.Int63())
		Logger.Debug().Msg("Hello World!")
		gls.Cleanup()
	})

	t.Run("warn", func(t *testing.T) {
		gls.Set("request_id", rand.Int63())
		Logger.Warn().Msg("Hello World!")
		gls.Cleanup()
	})

	t.Run("error", func(t *testing.T) {
		gls.Set("request_id", rand.Int63())
		Logger.Error().Msg("Hello World!")
		gls.Cleanup()
	})

	t.Run("panic", func(t *testing.T) {
		gls.Set("request_id", rand.Int63())
		assert.Panics(t, func() { Logger.Panic().Msg("Hello World!") })
		gls.Cleanup()
	})
}
