package url

import (
	"errors"
	"strings"
)

// A URL represents a parsed URL.
type URL struct {
	Scheme string
	Host   string
	Path   string
}

// Parse parses a raw url into a URL structure.
func Parse(rawURL string) (*URL, error) {
	scheme, rest, ok := strings.Cut(rawURL, ":")
	if !ok || scheme == "" {
		return nil, errors.New("missing scheme")
	}

	if !strings.HasPrefix(rest, "//") {
		return &URL{Scheme: scheme}, nil
	}
	host, path, _ := strings.Cut(rest[2:], "/")

	u := &URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}

	return u, nil
}

// String reassembles the URL into a URL string.
func (u *URL) String() string {
	if u == nil {
		return ""
	}
	var s string
	if sc := u.Scheme; sc != "" {
		s += sc + "://"
	}
	if h := u.Host; h != "" {
		s += h
	}
	if p := u.Path; p != "" {
		s += "/" + p
	}
	return s
}
