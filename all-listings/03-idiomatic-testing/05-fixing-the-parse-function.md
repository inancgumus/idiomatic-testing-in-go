# Listing 3.5: Fixing the Parse function

## [testing](https://github.com/inancgumus/gobyexample/blob/4b1d456bf50b0c577669407309d5a40099a0ddb5/testing) / [url](https://github.com/inancgumus/gobyexample/blob/4b1d456bf50b0c577669407309d5a40099a0ddb5/testing/url) / [url.go](https://github.com/inancgumus/gobyexample/blob/4b1d456bf50b0c577669407309d5a40099a0ddb5/testing/url/url.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package url

import (
	"errors"
	"fmt"
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
	scheme, rest, ok := strings.Cut(rawURL, "://")
	if !ok {
		return nil, errors.New("missing scheme")
	}

	host, path, _ := strings.Cut(rest, "/")

	u := &URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}

	return u, nil
}

// String reassembles the URL into a URL string.
func (u *URL) String() string {
	return fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, u.Path)
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -1,6 +1,10 @@
 package url
 
-import "fmt"
+import (
+	"errors"
+	"fmt"
+	"strings"
+)
 
 // A URL represents a parsed URL.
 type URL struct {
@@ -11,13 +15,20 @@ type URL struct {
 
 // Parse parses a raw url into a URL structure.
 func Parse(rawURL string) (*URL, error) {
+	scheme, rest, ok := strings.Cut(rawURL, "://")
+	if !ok {
+		return nil, errors.New("missing scheme")
+	}
+
+	host, path, _ := strings.Cut(rest, "/")
+
 	u := &URL{
-		Scheme: "https",
-		Host:   "go.dev",
-		Path:   "play",
+		Scheme: scheme,
+		Host:   host,
+		Path:   path,
 	}
 
 	return u, nil
 }
 
 // String reassembles the URL into a URL string.
```

