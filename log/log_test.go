package log_test

import (
	"fmt"
	"testing"

	"github.com/ironzhang/x-pearls/log"
)

func init() {
	log.PanicFunc = func(v interface{}) {
		fmt.Printf("%v\n", v)
	}

	log.ExitFunc = func(code int) {
		fmt.Printf("exit: %d\n", code)
	}
}

func PrintTestZLogs(msg string) {
	fmt.Println(msg)

	log.Debug("debug", 1, "2", 3.0)
	log.Debugf("debugf: %v, %v, %v", 1, "2", 3.0)
	log.Debugw("debugw", "A", 1, "B", "2", "C", 3.0)

	log.Trace("trace", 1, "2", 3.0)
	log.Tracef("tracef: %v, %v, %v", 1, "2", 3.0)
	log.Tracew("tracew", "A", 1, "B", "2", "C", 3.0)

	log.Info("info", 1, "2", 3.0)
	log.Infof("infof: %v, %v, %v", 1, "2", 3.0)
	log.Infow("infow", "A", 1, "B", "2", "C", 3.0)

	log.Warn("warn", 1, "2", 3.0)
	log.Warnf("warnf: %v, %v, %v", 1, "2", 3.0)
	log.Warnw("warnw", "A", 1, "B", "2", "C", 3.0)

	log.Error("error", 1, "2", 3.0)
	log.Errorf("errorf: %v, %v, %v", 1, "2", 3.0)
	log.Errorw("errorw", "A", 1, "B", "2", "C", 3.0)

	log.Panic("panic", 1, "2", 3.0)
	log.Panicf("panicf: %v, %v, %v", 1, "2", 3.0)
	log.Panicw("panicw", "A", 1, "B", "2", "C", 3.0)

	log.Fatal("fatal", 1, "2", 3.0)
	log.Fatalf("fatalf: %v, %v, %v", 1, "2", 3.0)
	log.Fatalw("fatalw", "A", 1, "B", "2", "C", 3.0)
}

func TestZlog(t *testing.T) {
	PrintTestZLogs("default")

	log.Default.SetLevel(log.DEBUG)
	PrintTestZLogs("debug")

	log.Default.SetLevel(log.TRACE)
	PrintTestZLogs("trace")

	log.Default.SetLevel(log.INFO)
	PrintTestZLogs("info")

	log.Default.SetLevel(log.WARN)
	PrintTestZLogs("warn")

	log.Default.SetLevel(log.ERROR)
	PrintTestZLogs("error")

	log.Default.SetLevel(log.FATAL)
	PrintTestZLogs("fatal")
}
