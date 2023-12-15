package parallel

import (
	"testing"
	"time"
)

func TestPar1(t *testing.T) {
	t.Parallel()
	time.Sleep(5 * time.Second)
	// the rest of the test code goes here...
}

func TestPar2(t *testing.T) {
	t.Parallel()
	time.Sleep(5 * time.Second)
	// ...
}

func TestPar3(t *testing.T) {
	t.Parallel()
	t.Run("subpar1", func(t *testing.T) {
		t.Parallel()
		time.Sleep(5 * time.Second)
		// ...
	})
	t.Run("subpar2", func(t *testing.T) {
		t.Parallel()
		time.Sleep(5 * time.Second)
		// ...
	})
}

func TestSeq(t *testing.T) {
	// ...
}
