package zlog_test

import (
	"math/rand"
	"testing"

	"github.com/DiscoFighter47/zlog"
	"github.com/stretchr/testify/assert"

	"github.com/tylerb/gls"
)

func TestLogger(t *testing.T) {
	t.Run("no id", func(t *testing.T) {
		zlog.Info("Hello World! no id")
		gls.Cleanup()
	})

	t.Run("info", func(t *testing.T) {
		gls.Set("request_id", rand.Int63())
		zlog.Info("Hello World!")
		gls.Cleanup()
	})

	t.Run("infof", func(t *testing.T) {
		gls.Set("request_id", rand.Int63())
		zlog.Infof("%s %s", "Hello", "World!")
		gls.Cleanup()
	})

	t.Run("debug", func(t *testing.T) {
		gls.Set("request_id", rand.Int63())
		zlog.Debug("Hello World!")
		gls.Cleanup()
	})

	t.Run("debug", func(t *testing.T) {
		gls.Set("request_id", rand.Int63())
		zlog.Debugf("%s %s", "Hello", "World!")
		gls.Cleanup()
	})

	t.Run("warn", func(t *testing.T) {
		gls.Set("request_id", rand.Int63())
		zlog.Warn("Hello World!")
		gls.Cleanup()
	})
	t.Run("warn", func(t *testing.T) {
		gls.Set("request_id", rand.Int63())
		zlog.Warnf("%s %s", "Hello", "World!")
		gls.Cleanup()
	})

	t.Run("error", func(t *testing.T) {
		gls.Set("request_id", rand.Int63())
		zlog.Error("Hello World!")
		gls.Cleanup()
	})

	t.Run("errorf", func(t *testing.T) {
		gls.Set("request_id", rand.Int63())
		zlog.Errorf("%s %s", "Hello", "World!")
		gls.Cleanup()
	})

	t.Run("panic", func(t *testing.T) {
		gls.Set("request_id", rand.Int63())
		assert.Panics(t, func() { zlog.Panic("Hello World!") })
		gls.Cleanup()
	})

	t.Run("panicf", func(t *testing.T) {
		gls.Set("request_id", rand.Int63())
		assert.Panics(t, func() { zlog.Panicf("%s %s", "Hello", "World!") })
		gls.Cleanup()
	})
}
