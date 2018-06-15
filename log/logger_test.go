package log

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func init() {
	PanicFunc = func(v interface{}) {
		fmt.Printf("%v\n", v)
	}

	ExitFunc = func(code int) {
		fmt.Printf("exit: %d\n", code)
	}
}

type T1 struct {
	A int
	B string
}

type T2 struct {
	a int
	b string
}

func (t T2) String() string {
	return fmt.Sprintf("{A:%d, B:%s}", t.a, t.b)
}

type T3 struct {
	a int
	b string
}

func (t T3) Error() string {
	return fmt.Sprintf("{A:%d, B:%s}", t.a, t.b)
}

func TestMarshal(t *testing.T) {
	tests := []struct {
		a interface{}
		s string
	}{
		{
			a: 1,
			s: "1",
		},
		{
			a: T1{A: 1, B: "B"},
			s: `{"A":1,"B":"B"}`,
		},
		{
			a: T2{a: 1, b: "B"},
			s: `"{A:1, B:B}"`,
		},
		{
			a: T3{a: 1, b: "B"},
			s: `"{A:1, B:B}"`,
		},
	}

	for i, tt := range tests {
		if got, want := marshal(tt.a), tt.s; got != want {
			t.Fatalf("%d: got %v, want %v", i, got, want)
		}
	}
}

func TestFmtkvs(t *testing.T) {
	tests := []struct {
		kvs []interface{}
		str string
	}{
		{
			kvs: []interface{}{"x", 1, "y", 2, "err", "error"},
			str: `"x": 1, "y": 2, "err": "error"`,
		},
		{
			kvs: []interface{}{"x", 1, "y", 2, "err"},
			str: `"x": 1, "y": 2, "err": "@none"`,
		},
		{
			kvs: []interface{}{"x", 1, 3, 2, "err"},
			str: `"x": 1, "@2": 2, "err": "@none"`,
		},
	}
	for i, tt := range tests {
		got, want := fmtkvs(tt.kvs), tt.str
		if got != want {
			t.Errorf("%d: %q != %q", i, got, want)
		}
		t.Logf("%d: %s", i, got)
	}
}

func TestSprint(t *testing.T) {
	tests := []struct {
		level Level
		args  []interface{}
		str   string
	}{
		{
			level: DEBUG,
			args:  []interface{}{"1", 2, 3, "4"},
			str:   "[DEBUG] 12 34",
		},
	}
	for _, tt := range tests {
		if got, want := sprint(tt.level, tt.args...), tt.str; got != want {
			t.Errorf("sprint: %v != %v", got, want)
		}
	}
}

func TestSprintf(t *testing.T) {
	tests := []struct {
		level  Level
		format string
		args   []interface{}
		str    string
	}{
		{
			level:  DEBUG,
			format: "%v %v %v %v",
			args:   []interface{}{"1", 2, 3, "4"},
			str:    "[DEBUG] 1 2 3 4",
		},
	}
	for _, tt := range tests {
		if got, want := sprintf(tt.level, tt.format, tt.args...), tt.str; got != want {
			t.Errorf("sprint: %v != %v", got, want)
		}
	}
}

func PrintTestLogs(msg string, l *StdLogger) {
	fmt.Println(msg)

	l.Debug("debug", 1, "2", 3.0)
	l.Debugf("debugf: %v, %v, %v", 1, "2", 3.0)
	l.Debugw("debugw", "x", 1, "y", 1.0)

	l.Trace("trace", 1, "2", 3.0)
	l.Tracef("tracef: %v, %v, %v", 1, "2", 3.0)
	l.Tracew("tracew", "x", 1, "y", 1.0)

	l.Info("info", 1, "2", 3.0)
	l.Infof("infof: %v, %v, %v", 1, "2", 3.0)
	l.Infow("infow", "x", 1, "y", 1.0)

	l.Warn("warn", 1, "2", 3.0)
	l.Warnf("warnf: %v, %v, %v", 1, "2", 3.0)
	l.Warnw("warnw", "x", 1, "y", 1.0)

	l.Error("error", 1, "2", 3.0)
	l.Errorf("errorf: %v, %v, %v", 1, "2", 3.0)
	l.Errorw("errorw", "x", 1, "y", 1.0)

	l.Panic("panic", 1, "2", 3.0)
	l.Panicf("panicf: %v, %v, %v", 1, "2", 3.0)
	l.Panicw("panicw", "x", 1, "y", 1.0)

	l.Fatal("fatal", 1, "2", 3.0)
	l.Fatalf("fatalf: %v, %v, %v", 1, "2", 3.0)
	l.Fatalw("fatalw", "x", 1, "y", 1.0)
}

func TestLogger(t *testing.T) {
	var enable bool
	enable = true
	if enable {
		l := NewStdLogger(log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile), DEBUG, 0)

		l.SetLevel(DEBUG)
		PrintTestLogs("debug level", l)

		l.SetLevel(TRACE)
		PrintTestLogs("trace level", l)

		l.SetLevel(INFO)
		PrintTestLogs("info level", l)

		l.SetLevel(WARN)
		PrintTestLogs("warn level", l)

		l.SetLevel(ERROR)
		PrintTestLogs("error level", l)

		l.SetLevel(FATAL)
		PrintTestLogs("fatal level", l)
	}
}
