package zaplog

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
)

type Logger struct {
	level  zap.AtomicLevel
	logger *zap.SugaredLogger
}

func NewLogger(logger *zap.Logger, level zap.AtomicLevel, calldepth int) *Logger {
	return &Logger{
		level:  level,
		logger: logger.WithOptions(zap.AddCallerSkip(calldepth + 1)).Sugar(),
	}
}

func (p *Logger) GetLogLevel() string {
	return p.level.String()
}

func (p *Logger) SetLogLevel(level string) error {
	switch strings.ToLower(level) {
	case zap.DebugLevel.String():
		p.level.SetLevel(zap.DebugLevel)
	case zap.InfoLevel.String():
		p.level.SetLevel(zap.InfoLevel)
	case zap.WarnLevel.String():
		p.level.SetLevel(zap.WarnLevel)
	case zap.ErrorLevel.String():
		p.level.SetLevel(zap.ErrorLevel)
	case zap.DPanicLevel.String():
		p.level.SetLevel(zap.DPanicLevel)
	case zap.PanicLevel.String():
		p.level.SetLevel(zap.PanicLevel)
	case zap.FatalLevel.String():
		p.level.SetLevel(zap.FatalLevel)
	default:
		return fmt.Errorf("%q level is unknown", level)
	}
	return nil
}

func (p *Logger) Debug(args ...interface{}) {
	p.logger.Debug(args...)
}

func (p *Logger) Debugf(format string, args ...interface{}) {
	p.logger.Debugf(format, args...)
}

func (p *Logger) Debugw(message string, kvs ...interface{}) {
	p.logger.Debugw(message, kvs...)
}

func (p *Logger) Trace(args ...interface{}) {
	p.logger.Debug(args...)
}

func (p *Logger) Tracef(format string, args ...interface{}) {
	p.logger.Debugf(format, args...)
}

func (p *Logger) Tracew(message string, kvs ...interface{}) {
	p.logger.Debugw(message, kvs...)
}

func (p *Logger) Info(args ...interface{}) {
	p.logger.Info(args...)
}

func (p *Logger) Infof(format string, args ...interface{}) {
	p.logger.Infof(format, args...)
}

func (p *Logger) Infow(message string, kvs ...interface{}) {
	p.logger.Infow(message, kvs...)
}

func (p *Logger) Warn(args ...interface{}) {
	p.logger.Warn(args...)
}

func (p *Logger) Warnf(format string, args ...interface{}) {
	p.logger.Warnf(format, args...)
}

func (p *Logger) Warnw(message string, kvs ...interface{}) {
	p.logger.Warnw(message, kvs...)
}

func (p *Logger) Error(args ...interface{}) {
	p.logger.Error(args...)
}

func (p *Logger) Errorf(format string, args ...interface{}) {
	p.logger.Errorf(format, args...)
}

func (p *Logger) Errorw(message string, kvs ...interface{}) {
	p.logger.Errorw(message, kvs...)
}

func (p *Logger) Panic(args ...interface{}) {
	p.logger.Panic(args...)
}

func (p *Logger) Panicf(format string, args ...interface{}) {
	p.logger.Panicf(format, args...)
}

func (p *Logger) Panicw(message string, kvs ...interface{}) {
	p.logger.Panicw(message, kvs...)
}

func (p *Logger) Fatal(args ...interface{}) {
	p.logger.Fatal(args...)
}

func (p *Logger) Fatalf(format string, args ...interface{}) {
	p.logger.Fatalf(format, args...)
}

func (p *Logger) Fatalw(message string, kvs ...interface{}) {
	p.logger.Fatalw(message, kvs...)
}
