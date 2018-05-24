package zaplog

import (
	"testing"

	"github.com/ironzhang/x-pearls/zlog"

	"go.uber.org/zap"
)

func NewTestZapLogger() *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.Level.SetLevel(zap.DebugLevel)
	l, err := cfg.Build(zap.AddStacktrace(zap.NewAtomicLevelAt(zap.DPanicLevel)))
	if err != nil {
		panic(err)
	}
	return l
}

func TestLogger(t *testing.T) {
	l := NewLogger(NewTestZapLogger(), 0)
	l.Debug("debug")
	l.Debugf("debug")
	l.Debugw("debug")
	l.Trace("trace")
	l.Tracef("trace")
	l.Tracew("trace")
	l.Info("info")
	l.Infof("info")
	l.Infow("info")
	l.Warn("warn")
	l.Warnf("warn")
	l.Warnw("warn")
	l.Error("error")
	l.Errorf("error")
	l.Errorw("error")

	func() {
		defer func() { recover() }()
		l.Panic("panic")
	}()

	func() {
		defer func() { recover() }()
		l.Panicf("panic")
	}()

	func() {
		defer func() { recover() }()
		l.Panicw("panic")
	}()

	//l.Fatal("fatal")
}

func TestZLog(t *testing.T) {
	l := NewLogger(NewTestZapLogger(), 1)
	zlog.SetLogger(l)

	zlog.Debug("debug")
	zlog.Trace("trace")
	zlog.Info("info")
	zlog.Warn("warn")
	zlog.Error("error")
	//zlog.Panic("panic")
	//zlog.Fatal("fatal")
}
