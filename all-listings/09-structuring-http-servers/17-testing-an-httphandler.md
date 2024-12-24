# Listing 9.17: Testing an http.Handler

## [bite](https://github.com/inancgumus/gobyexample/blob/91c9e6f10fd119ac6877b007ae5fe1dd680da20a/bite) / [link](https://github.com/inancgumus/gobyexample/blob/91c9e6f10fd119ac6877b007ae5fe1dd680da20a/bite/link) / [server_test.go](https://github.com/inancgumus/gobyexample/blob/91c9e6f10fd119ac6877b007ae5fe1dd680da20a/bite/link/server_test.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package link

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServerHealth(t *testing.T) {
	t.Parallel()

	var srv Server
	rec := httptest.NewRecorder()
	srv.Health(rec, httptest.NewRequest(http.MethodGet, "/", nil))

	if rec.Code != http.StatusOK {
		t.Errorf("got status code = %d, want %d", rec.Code, http.StatusOK)
	}
	if got := rec.Body.String(); !strings.Contains(got, "OK") {
		t.Errorf("got body = %s\twant contains %s", got, "OK")
	}
}
```

