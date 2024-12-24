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
