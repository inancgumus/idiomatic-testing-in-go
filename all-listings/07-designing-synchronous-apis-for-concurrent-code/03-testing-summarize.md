# Listing 7.3: Testing Summarize

## [hit](https://github.com/inancgumus/gobyexample/blob/13177088e620148048b4de5da87128701fba3447/hit) / [result_test.go](https://github.com/inancgumus/gobyexample/blob/13177088e620148048b4de5da87128701fba3447/hit/result_test.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
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
```

