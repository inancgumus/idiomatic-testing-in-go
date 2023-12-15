package link

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

// Link represents a link.
type Link struct {
	Key string
	URL string
}

// validateNewLink checks a new link's validity.
func validateNewLink(link Link) error {
	if err := validateLinkKey(link.Key); err != nil {
		return err
	}
	u, err := url.ParseRequestURI(link.URL)
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

// MaxKeyLen is the maximum length of a key.
const MaxKeyLen = 16

// validateLinkKey checks the key's validity.
func validateLinkKey(k string) error {
	if strings.TrimSpace(k) == "" {
		return errors.New("empty key")
	}
	if len(k) > MaxKeyLen {
		return fmt.Errorf("key too long (max %d)", MaxKeyLen)
	}
	return nil
}
