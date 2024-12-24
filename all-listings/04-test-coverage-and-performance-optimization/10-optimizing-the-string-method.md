# Listing 4.10: Optimizing the String method

## [testing](https://github.com/inancgumus/gobyexample/blob/d64225d13e86efbdf6a3ff980f5fc8b1eeeb30cc/testing) / [url](https://github.com/inancgumus/gobyexample/blob/d64225d13e86efbdf6a3ff980f5fc8b1eeeb30cc/testing/url) / [url.go](https://github.com/inancgumus/gobyexample/blob/d64225d13e86efbdf6a3ff980f5fc8b1eeeb30cc/testing/url/url.go)

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

	var s strings.Builder

	const (
		lenSchemeSeparator = len("://")
		lenPathSeparator   = len("/")
	)
	lenURL := len(u.Scheme) + lenSchemeSeparator +
		len(u.Host) + lenPathSeparator +
		len(u.Path)
	s.Grow(lenURL)

	if sc := u.Scheme; sc != "" {
		s.WriteString(sc)
		s.WriteString("://")
	}
	if h := u.Host; h != "" {
		s.WriteString(h)
	}
	if p := u.Path; p != "" {
		s.WriteByte('/')
		s.WriteString(p)
	}

	return s.String()
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -36,17 +36,31 @@ func Parse(rawURL string) (*URL, error) {
 // String reassembles the URL into a URL string.
 func (u *URL) String() string {
 	if u == nil {
 		return ""
 	}
-	var s string
+
+	var s strings.Builder
+
+	const (
+		lenSchemeSeparator = len("://")
+		lenPathSeparator   = len("/")
+	)
+	lenURL := len(u.Scheme) + lenSchemeSeparator +
+		len(u.Host) + lenPathSeparator +
+		len(u.Path)
+	s.Grow(lenURL)
+
 	if sc := u.Scheme; sc != "" {
-		s += sc + "://"
+		s.WriteString(sc)
+		s.WriteString("://")
 	}
 	if h := u.Host; h != "" {
-		s += h
+		s.WriteString(h)
 	}
 	if p := u.Path; p != "" {
-		s += "/" + p
+		s.WriteByte('/')
+		s.WriteString(p)
 	}
-	return s
+
+	return s.String()
 }
```

