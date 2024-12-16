package link

import (
	"context"
	"fmt"

	"github.com/inancgumus/gobyexample/bite"
)

// Store persists and retrieves [Link] values in an in-memory map.
type Store struct {
	links map[string]Link
}

// Create persists a [Link] in the store.
// It returns [bite.ErrInvalidRequest] if the [Link] is invalid
// or [bite.ErrExists] if the [Link] already exists.
func (s *Store) Create(_ context.Context, link Link) error {
	if err := validateNewLink(link); err != nil {
		return fmt.Errorf("%w: %w", bite.ErrInvalidRequest, err)
	}

	if _, ok := s.links[link.Key]; ok {
		return bite.ErrExists
	}
	if s.links == nil {
		s.links = map[string]Link{}
	}
	s.links[link.Key] = link

	return nil
}
