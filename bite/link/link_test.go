package link

import (
	"log/slog"
	"math"
	"testing"
)

func TestMain(m *testing.M) {
	slog.SetLogLoggerLevel(math.MaxInt)
	m.Run()
}
