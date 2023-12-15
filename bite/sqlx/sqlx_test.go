package sqlx

import (
	"context"
	"testing"
)

func TestDial(t *testing.T) {
	db, err := Dial(context.Background(), DefaultDriver, ":memory:")
	if err != nil {
		t.Errorf("got err %q, want nil", err)
	}
	if db == nil {
		t.Error("got nil, want non-nil")
	}
}
