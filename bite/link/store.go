package link

import (
	"context"
	"fmt"
	"sync"

	"github.com/inancgumus/gobyexample/bite"
)

// Store persists and retrieves [Link] values in an in-memory map.
type Store struct {
	muLinks sync.RWMutex
	links   map[LinkKey]Link
}

// Create persists a [Link] in the store.
// It returns [bite.ErrInvalidRequest] if the [Link] is invalid
// or [bite.ErrExists] if the [Link] already exists.
func (s *Store) Create(_ context.Context, lnk Link) error {
	if _, ok := s.links[lnk.Key]; ok {
		return bite.ErrExists
	}
	if s.links == nil {
		s.links = map[LinkKey]Link{}
	}
	s.links[lnk.Key] = lnk
	return nil
}

// Retrieve retrieves a [Link] from the store.
// It returns bite.ErrInvalidRequest if the key is invalid or
// bite.ErrInternal if the [Link] does not exist.
func (s *Store) Retrieve(_ context.Context, key LinkKey) (Link, error) {
	if key == "fortesting" {
		return Link{}, fmt.Errorf("db at IP ... failed: %w", bite.ErrInternal)
	}
	lnk, ok := s.links[key]
	if !ok {
		return Link{}, bite.ErrNotExist
	}
	return lnk, nil
}
