package hit

import (
	"slices"
	"testing"
	"time"
)

func TestSummarize(t *testing.T) {
	results := []Result{
		{Duration: 2 * time.Second /*, other fields */},
		{Duration: 5 * time.Second /*, other fields */},
	}
	sum := Summarize(Results(slices.Values(results)))
	if sum.Fastest != 2*time.Second {
		t.Errorf("Fastest=%v; want 2s", sum.Fastest)
	}
	// other tests
}

func TestSummarizeNilResults(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("should not panic: %v", err)
		}
	}()
	_ = Summarize(nil)
}
