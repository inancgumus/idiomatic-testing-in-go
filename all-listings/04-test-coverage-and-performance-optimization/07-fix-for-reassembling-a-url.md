# Listing 4.7: Fix for reassembling a URL

## [testing](https://github.com/inancgumus/gobyexample/blob/e6af5b0765d6b50525ea559ce2c0d21dae7edef8/testing) / [url](https://github.com/inancgumus/gobyexample/blob/e6af5b0765d6b50525ea559ce2c0d21dae7edef8/testing/url) / [url.go](https://github.com/inancgumus/gobyexample/blob/e6af5b0765d6b50525ea559ce2c0d21dae7edef8/testing/url/url.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
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
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -2,8 +2,7 @@ package url
 
 import (
 	"errors"
-	"fmt"
 	"strings"
 )
 
 // A URL represents a parsed URL.
@@ -37,7 +36,17 @@ func Parse(rawURL string) (*URL, error) {
 // String reassembles the URL into a URL string.
 func (u *URL) String() string {
 	if u == nil {
 		return ""
 	}
-	return fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, u.Path)
+	var s string
+	if sc := u.Scheme; sc != "" {
+		s += sc + "://"
+	}
+	if h := u.Host; h != "" {
+		s += h
+	}
+	if p := u.Path; p != "" {
+		s += "/" + p
+	}
+	return s
 }
```

