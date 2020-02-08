package zlog_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tylerb/gls"

	"github.com/DiscoFighter47/zlog"
)

func TestSetReqID(t *testing.T) {
	t.Run("initiat", func(t *testing.T) {
		zlog.SetReqID()
		v := gls.Get(zlog.ReqIDTag)
		assert.NotNil(t, v)
	})

	t.Run("no initiat", func(t *testing.T) {
		v := gls.Get(zlog.ReqIDTag)
		assert.Nil(t, v)
	})
}

func TestGetReqID(t *testing.T) {
	t.Run("not initiated", func(t *testing.T) {
		assert.Empty(t, zlog.GetReqID())
	})

	t.Run("initiated", func(t *testing.T) {
		gls.Set(zlog.ReqIDTag, "Hello")
		assert.NotEmpty(t, zlog.GetReqID())
	})
}
