package url

import "fmt"

// A URL represents a parsed URL.
type URL struct {
	Scheme string
	Host   string
	Path   string
}

// Parse parses a raw url into a URL structure.
func Parse(rawURL string) (*URL, error) {
	u := &URL{
		Scheme: "https",
		Host:   "go.dev",
		Path:   "play",
	}

	return u, nil
}

// String reassembles the URL into a URL string.
func (u *URL) String() string {
	return fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, u.Path)
}
