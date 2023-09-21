package log

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	opt := []Option{
		WithMaxAge(time.Hour * 7 * 24),
	}
	New("./log/access.log", opt...)
}

func TestDebug(t *testing.T) {
	opt := []Option{
		WithMaxAge(time.Hour * 7 * 24),
		WithFileFormat(".%Y-%m-%d"),
	}
	New("./log/access.log", opt...)
	Debug("test debug: %d", 1)
}

func TestInfo(t *testing.T) {
	opt := []Option{
		WithMaxAge(time.Hour * 7 * 24),
		WithFileFormat(".%Y-%m-%d"),
	}
	New("./log/access.log", opt...)
	Info("info data: %s", "hi")
}

func TestError(t *testing.T) {
	opt := []Option{
		WithMaxAge(time.Hour * 7 * 24),
		WithFileFormat(".%Y-%m-%d"),
		WithDistWarnLevel(true),
		WithCallerSkip(0),
	}
	New("./log/access.log", opt...)
	Error("error data: %s", time.Now().Format("2006-01-02 15:04:05"))
}

func TestWarn(t *testing.T) {
	opt := []Option{
		WithMaxAge(time.Hour * 7 * 24),
		WithFileFormat(".%Y-%m-%d"),
		WithDistWarnLevel(true),
	}
	New("./log/access.log", opt...)
	Warn("warn data: %s", "hi")
}

func TestFatal(t *testing.T) {
	opt := []Option{
		WithMaxAge(time.Hour * 7 * 24),
		WithFileFormat(".%Y-%m-%d"),
	}
	New("./log/access.log", opt...)
	Fatal("fatal data")
}
