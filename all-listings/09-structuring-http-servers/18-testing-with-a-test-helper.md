# Listing 9.18: Testing with a test helper

## [bite](https://github.com/inancgumus/gobyexample/blob/717a701fa0cc926e2b6db1b970143e02c240544c/bite) / [link](https://github.com/inancgumus/gobyexample/blob/717a701fa0cc926e2b6db1b970143e02c240544c/bite/link) / [server_test.go](https://github.com/inancgumus/gobyexample/blob/717a701fa0cc926e2b6db1b970143e02c240544c/bite/link/server_test.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package link

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServerHealth(t *testing.T) {
	t.Parallel()

	var srv Server
	rec := httptest.NewRecorder()
	srv.Health(rec, newRequest(t, http.MethodGet, "/", nil))

	if rec.Code != http.StatusOK {
		t.Errorf("got status code = %d, want %d", rec.Code, http.StatusOK)
	}
	if got := rec.Body.String(); !strings.Contains(got, "OK") {
		t.Errorf("got body = %s\twant contains %s", got, "OK")
	}
}

// newRequest creates a new [http.Request] for testing.
// It fails the calling test if it cannot create the request.
func newRequest(
	tb testing.TB, method string, target string, body io.Reader,
) *http.Request {
	tb.Helper()
	r, err := http.NewRequest(method, target, body)
	if err != nil {
		tb.Fatalf("http.NewRequest() err = %v, want nil", err)
	}
	return r
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -1,8 +1,9 @@
 package link
 
 import (
+	"io"
 	"net/http"
 	"net/http/httptest"
 	"strings"
 	"testing"
 )
@@ -10,14 +11,27 @@ import (
 func TestServerHealth(t *testing.T) {
 	t.Parallel()
 
 	var srv Server
 	rec := httptest.NewRecorder()
-	srv.Health(rec, httptest.NewRequest(http.MethodGet, "/", nil))
+	srv.Health(rec, newRequest(t, http.MethodGet, "/", nil))
 
 	if rec.Code != http.StatusOK {
 		t.Errorf("got status code = %d, want %d", rec.Code, http.StatusOK)
 	}
 	if got := rec.Body.String(); !strings.Contains(got, "OK") {
 		t.Errorf("got body = %s\twant contains %s", got, "OK")
 	}
 }
+
+// newRequest creates a new [http.Request] for testing.
+// It fails the calling test if it cannot create the request.
+func newRequest(
+	tb testing.TB, method string, target string, body io.Reader,
+) *http.Request {
+	tb.Helper()
+	r, err := http.NewRequest(method, target, body)
+	if err != nil {
+		tb.Fatalf("http.NewRequest() err = %v, want nil", err)
+	}
+	return r
+}
```

