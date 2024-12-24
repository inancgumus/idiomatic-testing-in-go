# Listing 3.1: The url package

## [testing](https://github.com/inancgumus/gobyexample/blob/f45986c6e0a9db02cfd8725d66107079cd8c4fa0/testing) / [url](https://github.com/inancgumus/gobyexample/blob/f45986c6e0a9db02cfd8725d66107079cd8c4fa0/testing/url) / [url.go](https://github.com/inancgumus/gobyexample/blob/f45986c6e0a9db02cfd8725d66107079cd8c4fa0/testing/url/url.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
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
```

