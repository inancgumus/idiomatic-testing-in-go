# Listing 9.7: Protecting the map access

## [bite](https://github.com/inancgumus/gobyexample/blob/b300f48619f76873b0fdfb531d0ac057c91b4cc2/bite) / [link](https://github.com/inancgumus/gobyexample/blob/b300f48619f76873b0fdfb531d0ac057c91b4cc2/bite/link) / [store.go](https://github.com/inancgumus/gobyexample/blob/b300f48619f76873b0fdfb531d0ac057c91b4cc2/bite/link/store.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
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
	s.muLinks.Lock()
	defer s.muLinks.Unlock()

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
	s.muLinks.RLock()
	defer s.muLinks.RUnlock()

	lnk, ok := s.links[key]
	if !ok {
		return Link{}, bite.ErrNotExist
	}
	return lnk, nil
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -17,27 +17,33 @@ type Store struct {
 // Create persists a [Link] in the store.
 // It returns [bite.ErrInvalidRequest] if the [Link] is invalid
 // or [bite.ErrExists] if the [Link] already exists.
 func (s *Store) Create(_ context.Context, lnk Link) error {
+	s.muLinks.Lock()
+	defer s.muLinks.Unlock()
+
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
+	s.muLinks.RLock()
+	defer s.muLinks.RUnlock()
+
 	lnk, ok := s.links[key]
 	if !ok {
 		return Link{}, bite.ErrNotExist
 	}
 	return lnk, nil
 }
```

