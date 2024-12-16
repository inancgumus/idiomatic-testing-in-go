# Listing 9.18: Testing with a test helper

## [bite](https://github.com/inancgumus/gobyexample/blob/ab3515bfdff2e130c0a8eb4502b6984f76218262/bite) / [link](https://github.com/inancgumus/gobyexample/blob/ab3515bfdff2e130c0a8eb4502b6984f76218262/bite/link) / [server_test.go](https://github.com/inancgumus/gobyexample/blob/ab3515bfdff2e130c0a8eb4502b6984f76218262/bite/link/server_test.go)

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
	w := httptest.NewRecorder()
	srv.Health(w, newRequest(t, http.MethodGet, "/", nil))

	if w.Code != http.StatusOK {
		t.Errorf("got status code = %d, want %d", w.Code, http.StatusOK)
	}
	if got := w.Body.String(); !strings.Contains(got, "OK") {
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
@@ -1,6 +1,7 @@
 package link
 
 import (
+	"io"
 	"net/http"
 	"net/http/httptest"
 	"strings"
@@ -12,7 +13,7 @@ func TestServerHealth(t *testing.T) {
 
 	var srv Server
 	w := httptest.NewRecorder()
-	srv.Health(w, httptest.NewRequest(http.MethodGet, "/", nil))
+	srv.Health(w, newRequest(t, http.MethodGet, "/", nil))
 
 	if w.Code != http.StatusOK {
 		t.Errorf("got status code = %d, want %d", w.Code, http.StatusOK)
@@ -21,3 +22,16 @@ func TestServerHealth(t *testing.T) {
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

