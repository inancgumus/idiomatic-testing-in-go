// Package sqlx provides a wrapper around database/sql.DB.
//
// To take control of the database driver dependencies, sqlite
// driver import happens in a single package. Also, this is a
// good place to change the driver or support multiple ones.
// Another benefit of this package is that it provides an
// interception point for the database operations. For example,
// to log/trace the queries; provide a simpler API, etc.
package sqlx

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"

	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

// DefaultDriver is the default SQL driver.
const DefaultDriver = "sqlite"

//go:embed schema.sql
var schema string

// DB is a wrapper around [database/sql.DB].
type DB struct {
	*sql.DB
}

// Dial is like [sql.Open] + [sql.Ping] and also migrates the schema.
func Dial(ctx context.Context, driver, dsn string) (*DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, fmt.Errorf(
			"opening database driver %q by %q: %w", driver, dsn, err,
		)
	}
	// it wasn't necessary to ping the database as ExecContext
	// will do it already. kept it here for the sake of the
	// example.
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("pinging database: %w", err)
	}
	if _, err := db.ExecContext(ctx, schema); err != nil {
		return nil, fmt.Errorf("migrating schema: %w", err)
	}

	return &DB{DB: db}, nil
}

// IsPrimaryKeyViolation returns true if the given error is a
// primary key constraint error. For example, this can be used
// to check if a primary key already exists in the database.
func IsPrimaryKeyViolation(err error) bool {
	var serr *sqlite.Error
	if errors.As(err, &serr) {
		return serr.Code() == sqlite3.SQLITE_CONSTRAINT_PRIMARYKEY
	}
	return false
}
