package log

import (
	"context"
	"github.com/segmentio/ksuid"
	"sync"
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
	ctx := context.Background()
	ctx = context.WithValue(context.Background(), "rid", ksuid.New().String())
	New("./log/access.log", opt...)

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		Debug(ctx, "test debug: %d", 1008612)
	}()
	go func() {
		defer wg.Done()
		Info(ctx, "test debug: %d", 1008613)
	}()
	go func() {
		defer wg.Done()
		Warn(ctx, "test debug: %d", 1008614)
	}()
	wg.Wait()
}

func TestInfo(t *testing.T) {
	opt := []Option{
		WithMaxAge(time.Hour * 7 * 24),
		WithFileFormat(".%Y-%m-%d"),
	}
	New("./log/access.log", opt...)
	Info(context.Background(), "info data: %s", "hi")
}

func TestError(t *testing.T) {
	opt := []Option{
		WithMaxAge(time.Hour * 7 * 24),
		WithFileFormat(".%Y-%m-%d"),
		WithDistWarnLevel(true),
		WithCallerSkip(0),
	}
	New("./log/access.log", opt...)
	Error(context.Background(), "error data: %s", time.Now().Format("2006-01-02 15:04:05"))
}

func TestWarn(t *testing.T) {
	opt := []Option{
		WithMaxAge(time.Hour * 7 * 24),
		WithFileFormat(".%Y-%m-%d"),
		WithDistWarnLevel(true),
	}
	New("./log/access.log", opt...)
	Warn(context.Background(), "warn data: %s", "hi")
}

func TestFatal(t *testing.T) {
	opt := []Option{
		WithMaxAge(time.Hour * 7 * 24),
		WithFileFormat(".%Y-%m-%d"),
	}
	New("./log/access.log", opt...)
	Fatal(context.Background(), "fatal data")
}
