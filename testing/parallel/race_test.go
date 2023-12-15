package parallel

import "testing"

var counter int //nolint:gochecknoglobals

func incr() { counter++ }

func TestDataRacy(t *testing.T) {
	t.Skip("this is a racy test. remove this line to see the data race")

	t.Parallel()
	t.Run("", func(t *testing.T) {
		t.Parallel()
		incr()
		// ...
	})
	t.Run("", func(t *testing.T) {
		t.Parallel()
		incr()
		// ...
	})
}
