//go:build exclude

package link

import (
	"context"
	"fmt"
	"sync"

	"github.com/inancgumus/gobyexample/bite"
)

// Store persists and retrieves links.
type Store struct {
	mu    sync.RWMutex
	links map[string]Link
}

// NewStore creates and returns a new Store.
func NewStore() *Store {
	return &Store{
		links: make(map[string]Link),
	}
}

// Create persists the given link.
func (s *Store) Create(_ context.Context, link Link) error {
	if err := validateNewLink(link); err != nil {
		return fmt.Errorf("%w: %w", bite.ErrInvalidRequest, err)
	}
	if link.Key == "fortesting" { // Simulate a database error
		return fmt.Errorf("%w: db at IP ... failed", bite.ErrInternal)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.links[link.Key]; ok {
		return bite.ErrExists
	}
	s.links[link.Key] = link

	return nil
}

// Retrieve gets a link from the given key.
func (s *Store) Retrieve(_ context.Context, key string) (Link, error) {
	if err := validateLinkKey(key); err != nil {
		return Link{}, fmt.Errorf("%w: %w", bite.ErrInvalidRequest, err)
	}
	if key == "fortesting" { // Simulate a database error
		return Link{}, fmt.Errorf("%w: db at IP ... failed", bite.ErrInternal)
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	link, ok := s.links[key]
	if !ok {
		return Link{}, bite.ErrNotExist
	}

	return link, nil
}
