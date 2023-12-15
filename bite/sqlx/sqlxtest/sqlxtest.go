// Package sqlxtest provides a test database pool.
package sqlxtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/inancgumus/gobyexample/bite/sqlx"
)

// Dial opens and returns an in-memory test database pool
// that is unique to the test name. The pool is closed when
// the test is finished.
// Avoid calling Dial multiple times in a single test.
func Dial(ctx context.Context, tb testing.TB) *sqlx.DB {
	tb.Helper()

	dsn := fmt.Sprintf("file:%s?mode=memory&cache=shared", tb.Name())

	db, err := sqlx.Dial(ctx, sqlx.DefaultDriver, dsn)
	if err != nil {
		tb.Fatalf("dialing test db: %v", err)
	}
	tb.Cleanup(func() {
		if err := db.Close(); err != nil {
			tb.Log("closing test db:", err)
		}
	})

	return db
}

// You can add another Dial here for opening a non-memory database.
