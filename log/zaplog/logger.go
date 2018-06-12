package zaplog

import "go.uber.org/zap"

type Logger struct {
	logger *zap.SugaredLogger
}

func NewLogger(logger *zap.Logger, calldepth int) *Logger {
	return &Logger{logger: logger.WithOptions(zap.AddCallerSkip(calldepth + 1)).Sugar()}
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
