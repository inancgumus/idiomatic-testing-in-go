# Listing 9.2: Implementing the core logic

## [bite](https://github.com/inancgumus/gobyexample/blob/aeaa55c5d3d2d992833fae080513aea9d51d7ef2/bite) / [link](https://github.com/inancgumus/gobyexample/blob/aeaa55c5d3d2d992833fae080513aea9d51d7ef2/bite/link) / [link.go](https://github.com/inancgumus/gobyexample/blob/aeaa55c5d3d2d992833fae080513aea9d51d7ef2/bite/link/link.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package link

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

// Link represents a shortened link.
type Link struct {
	// Key is the shortening key of the link.
	Key LinkKey

	// URL is the original URL that the key points to.
	URL string
}

// Validate validates [Link].
func (lnk Link) Validate() error {
	if err := lnk.Key.Validate(); err != nil {
		return fmt.Errorf("validating link key: %w", err)
	}
	u, err := url.ParseRequestURI(lnk.URL)
	if err != nil {
		return err
	}
	if u.Host == "" {
		return errors.New("empty host")
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("scheme must be http or https")
	}
	return nil
}

// LinkKey represents a unique identifier for a [Link.URL].
type LinkKey string

// Validate validates [LinkKey].
func (key LinkKey) Validate() error {
	if strings.TrimSpace(string(key)) == "" {
		return errors.New("empty key")
	}
	const MaxKeyLen = 16
	if len(string(key)) > MaxKeyLen {
		return fmt.Errorf("key too long (max %d)", MaxKeyLen)
	}
	return nil
}
```

