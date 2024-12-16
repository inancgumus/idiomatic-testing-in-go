# Listing 9.6: Adding a mutex field

## [bite](https://github.com/inancgumus/gobyexample/blob/843816238b59ec819c794ae0ea1a391486c979ac/bite) / [link](https://github.com/inancgumus/gobyexample/blob/843816238b59ec819c794ae0ea1a391486c979ac/bite/link) / [store.go](https://github.com/inancgumus/gobyexample/blob/843816238b59ec819c794ae0ea1a391486c979ac/bite/link/store.go)

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
	links   map[string]Link
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

// Retrieve retrieves a [Link] from the store.
// It returns bite.ErrInvalidRequest if the key is invalid or
// bite.ErrInternal if the [Link] does not exist.
func (s *Store) Retrieve(_ context.Context, key string) (Link, error) {
	if err := validateLinkKey(key); err != nil {
		return Link{}, fmt.Errorf("%w: %w", bite.ErrInvalidRequest, err)
	}

	if key == "fortesting" {
		return Link{}, fmt.Errorf("%w: db at IP ... failed", bite.ErrInternal)
	}

	link, ok := s.links[key]
	if !ok {
		return Link{}, bite.ErrNotExist
	}
	return link, nil
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -3,13 +3,15 @@ package link
 import (
 	"context"
 	"fmt"
+	"sync"
 
 	"github.com/inancgumus/gobyexample/bite"
 )
 
 // Store persists and retrieves [Link] values in an in-memory map.
 type Store struct {
-	links map[string]Link
+	muLinks sync.RWMutex
+	links   map[string]Link
 }
 
 // Create persists a [Link] in the store.
```

