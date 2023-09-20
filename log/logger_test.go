package log

import (
	"testing"
)

func TestNew(t *testing.T) {
	New("./log/access.log")
}

func TestDebug(t *testing.T) {
	New("./log/test-debug.log")
	Debug("test debug: %d", 1)
}
