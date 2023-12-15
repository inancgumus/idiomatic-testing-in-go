package link

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/inancgumus/gobyexample/bite"
	"github.com/inancgumus/gobyexample/bite/sqlx"
)

// ------------------------------------------------------------
// NOTE TO THE READER:
//
// See store.prev.go for the in-memory implementation of Store.
// ------------------------------------------------------------

var (
	ErrLinkExists   = fmt.Errorf("link %w", bite.ErrExists)
	ErrLinkNotExist = fmt.Errorf("link %w", bite.ErrNotExist)
)

// Store persists and retrieves links.
type Store struct {
	// db is the database for the link store.
	db *sqlx.DB
}

// NewStore creates and returns a new [Store].
func NewStore(db *sqlx.DB) *Store {
	return &Store{db: db}
}

// Create persists the given [Link]. It returns [bite.ErrInvalidRequest]
// if the link is invalid. Or it returns an error if the link cannot
// be created.
func (s *Store) Create(ctx context.Context, link Link) error {
	if err := validateNewLink(link); err != nil {
		return fmt.Errorf("%w: %w", bite.ErrInvalidRequest, err)
	}

	const query = `
		INSERT INTO links (
			short_key, uri
		) VALUES (
			?, ?
		)`

	url := sqlx.Base64String(link.URL)

	_, err := s.db.ExecContext(ctx, query, link.Key, url)
	if sqlx.IsPrimaryKeyViolation(err) {
		return ErrLinkExists
	}
	if err != nil {
		return fmt.Errorf("creating link: %w", err)
	}

	return nil
}

// Retrieve returns a [Link] from the given key. It returns [bite.ErrInvalidRequest]
// if the key is invalid. Or it returns [bite.ErrNotExist] if the link does not exist.
// Or it returns an error if the link cannot be retrieved.
func (s *Store) Retrieve(ctx context.Context, key string) (Link, error) {
	if err := validateLinkKey(key); err != nil {
		return Link{}, fmt.Errorf("%w: %w", bite.ErrInvalidRequest, err)
	}

	const query = `
		SELECT uri
		FROM links
		WHERE short_key = ?`

	var url sqlx.Base64String

	err := s.db.QueryRowContext(ctx, query, key).Scan(&url)
	if errors.Is(err, sql.ErrNoRows) {
		return Link{}, ErrLinkNotExist
	}
	if err != nil {
		return Link{}, fmt.Errorf("retrieving link by key %q: %w", key, err)
	}

	return Link{
		Key: key,
		URL: url.String(),
	}, nil
}
