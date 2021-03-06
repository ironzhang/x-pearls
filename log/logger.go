package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

var ExitFunc = os.Exit

var PanicFunc = func(v interface{}) {
	panic(v)
}

type Level int

const (
	DEBUG Level = -2
	TRACE Level = -1
	INFO  Level = 0
	WARN  Level = 1
	ERROR Level = 2
	PANIC Level = 3
	FATAL Level = 4
)

func (l Level) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case PANIC:
		return "PANIC"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

type StdLogger struct {
	logger    *log.Logger
	level     Level
	calldepth int
}

func NewStdLogger(logger *log.Logger, level Level, calldepth int) *StdLogger {
	return &StdLogger{
		logger:    logger,
		level:     level,
		calldepth: calldepth + 2,
	}
}

func (p *StdLogger) SetLogger(l *log.Logger) {
	p.logger = l
}

func (p *StdLogger) GetLevel() Level {
	return p.level
}

func (p *StdLogger) SetLevel(l Level) {
	p.level = l
}

func (p *StdLogger) SetCalldepth(calldepth int) {
	p.calldepth = calldepth
}

func (p *StdLogger) GetLogLevel() string {
	return p.GetLevel().String()
}

func (p *StdLogger) SetLogLevel(level string) error {
	switch strings.ToUpper(level) {
	case DEBUG.String():
		p.SetLevel(DEBUG)
	case TRACE.String():
		p.SetLevel(TRACE)
	case INFO.String():
		p.SetLevel(INFO)
	case WARN.String():
		p.SetLevel(WARN)
	case ERROR.String():
		p.SetLevel(ERROR)
	case PANIC.String():
		p.SetLevel(PANIC)
	case FATAL.String():
		p.SetLevel(FATAL)
	default:
		return fmt.Errorf("%q level is unknown", level)
	}
	return nil
}

func (p *StdLogger) Debug(args ...interface{}) {
	if p.level <= DEBUG {
		p.logger.Output(p.calldepth, sprint(DEBUG, args...))
	}
}

func (p *StdLogger) Debugf(format string, args ...interface{}) {
	if p.level <= DEBUG {
		p.logger.Output(p.calldepth, sprintf(DEBUG, format, args...))
	}
}

func (p *StdLogger) Debugw(message string, kvs ...interface{}) {
	if p.level <= DEBUG {
		p.logger.Output(p.calldepth, sprintkvs(DEBUG, message, kvs...))
	}
}

func (p *StdLogger) Trace(args ...interface{}) {
	if p.level <= TRACE {
		p.logger.Output(p.calldepth, sprint(TRACE, args...))
	}
}

func (p *StdLogger) Tracef(format string, args ...interface{}) {
	if p.level <= TRACE {
		p.logger.Output(p.calldepth, sprintf(TRACE, format, args...))
	}
}

func (p *StdLogger) Tracew(message string, kvs ...interface{}) {
	if p.level <= TRACE {
		p.logger.Output(p.calldepth, sprintkvs(TRACE, message, kvs...))
	}
}

func (p *StdLogger) Info(args ...interface{}) {
	if p.level <= INFO {
		p.logger.Output(p.calldepth, sprint(INFO, args...))
	}
}

func (p *StdLogger) Infof(format string, args ...interface{}) {
	if p.level <= INFO {
		p.logger.Output(p.calldepth, sprintf(INFO, format, args...))
	}
}

func (p *StdLogger) Infow(message string, kvs ...interface{}) {
	if p.level <= INFO {
		p.logger.Output(p.calldepth, sprintkvs(INFO, message, kvs...))
	}
}

func (p *StdLogger) Warn(args ...interface{}) {
	if p.level <= WARN {
		p.logger.Output(p.calldepth, sprint(WARN, args...))
	}
}

func (p *StdLogger) Warnf(format string, args ...interface{}) {
	if p.level <= WARN {
		p.logger.Output(p.calldepth, sprintf(WARN, format, args...))
	}
}

func (p *StdLogger) Warnw(message string, kvs ...interface{}) {
	if p.level <= WARN {
		p.logger.Output(p.calldepth, sprintkvs(WARN, message, kvs...))
	}
}

func (p *StdLogger) Error(args ...interface{}) {
	if p.level <= ERROR {
		p.logger.Output(p.calldepth, sprint(ERROR, args...))
	}
}

func (p *StdLogger) Errorf(format string, args ...interface{}) {
	if p.level <= ERROR {
		p.logger.Output(p.calldepth, sprintf(ERROR, format, args...))
	}
}

func (p *StdLogger) Errorw(message string, kvs ...interface{}) {
	if p.level <= ERROR {
		p.logger.Output(p.calldepth, sprintkvs(ERROR, message, kvs...))
	}
}

func (p *StdLogger) Panic(args ...interface{}) {
	if p.level <= PANIC {
		p.logger.Output(p.calldepth, sprint(PANIC, args...))
	}
	PanicFunc(fmt.Sprint(args...))
}

func (p *StdLogger) Panicf(format string, args ...interface{}) {
	if p.level <= PANIC {
		p.logger.Output(p.calldepth, sprintf(PANIC, format, args...))
	}
	PanicFunc(fmt.Sprintf(format, args...))
}

func (p *StdLogger) Panicw(message string, kvs ...interface{}) {
	if p.level <= PANIC {
		p.logger.Output(p.calldepth, sprintkvs(PANIC, message, kvs...))
	}
	PanicFunc(messagekvs(message, kvs))
}

func (p *StdLogger) Fatal(args ...interface{}) {
	if p.level <= FATAL {
		p.logger.Output(p.calldepth, sprint(FATAL, args...))
	}
	ExitFunc(1)
}

func (p *StdLogger) Fatalf(format string, args ...interface{}) {
	if p.level <= FATAL {
		p.logger.Output(p.calldepth, sprintf(FATAL, format, args...))
	}
	ExitFunc(1)
}

func (p *StdLogger) Fatalw(message string, kvs ...interface{}) {
	if p.level <= FATAL {
		p.logger.Output(p.calldepth, sprintkvs(FATAL, message, kvs...))
	}
	ExitFunc(1)
}

func sprint(l Level, args ...interface{}) string {
	return "[" + l.String() + "] " + fmt.Sprint(args...)
}

func sprintf(l Level, format string, args ...interface{}) string {
	return "[" + l.String() + "] " + fmt.Sprintf(format, args...)
}

func sprintkvs(l Level, message string, kvs ...interface{}) string {
	return "[" + l.String() + "] " + messagekvs(message, kvs)
}

func messagekvs(message string, kvs []interface{}) string {
	return fmt.Sprintf("%s\t{%s}", message, fmtkvs(kvs))
}

func fmtkvs(kvs []interface{}) string {
	var buf bytes.Buffer
	for i := 0; i < len(kvs); i += 2 {
		key, ok := kvs[i].(string)
		if !ok {
			key = fmt.Sprintf("@%d", i)
		}
		var val interface{}
		if i+1 < len(kvs) {
			val = kvs[i+1]
		} else {
			val = "@none"
		}

		if i == 0 {
			fmt.Fprintf(&buf, "%q: %s", key, marshal(val))
		} else {
			fmt.Fprintf(&buf, ", %q: %s", key, marshal(val))
		}
	}
	return buf.String()
}

func marshal(a interface{}) string {
	switch v := a.(type) {
	case error:
		return fmt.Sprintf("%q", v.Error())
	case fmt.Stringer:
		return fmt.Sprintf("%q", v.String())
	default:
		data, _ := json.Marshal(a)
		return string(data)
	}
}
