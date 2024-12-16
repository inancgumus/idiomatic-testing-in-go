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
	// It is used to resolve the original URL.
	// It must be unique, not empty, and at most
	// 16 characters long.
	Key string

	// URL is the original URL that the key points to.
	// It must be a valid URL.
	URL string
}

// validateNewLink checks the link's validity.
func validateNewLink(link Link) error {
	if err := validateLinkKey(link.Key); err != nil {
		return err
	}
	u, err := url.Parse(link.URL)
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

// validateLinkKey checks the link key's validity.
func validateLinkKey(key string) error {
	if strings.TrimSpace(key) == "" {
		return errors.New("empty key")
	}
	// An arbitrary number to keep the keys short.
	const MaxKeyLen = 16
	if len(key) > MaxKeyLen {
		return fmt.Errorf("key too long (max %d)", MaxKeyLen)
	}
	return nil
}
