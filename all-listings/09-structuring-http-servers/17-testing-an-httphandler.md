# Listing 9.17: Testing an http.Handler

## [bite](https://github.com/inancgumus/gobyexample/blob/ce77c111fad17f4466a22d0fa05b7470d24d88f6/bite) / [link](https://github.com/inancgumus/gobyexample/blob/ce77c111fad17f4466a22d0fa05b7470d24d88f6/bite/link) / [server_test.go](https://github.com/inancgumus/gobyexample/blob/ce77c111fad17f4466a22d0fa05b7470d24d88f6/bite/link/server_test.go)

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
	w := httptest.NewRecorder()
	srv.Health(w, httptest.NewRequest(http.MethodGet, "/", nil))

	if w.Code != http.StatusOK {
		t.Errorf("got status code = %d, want %d", w.Code, http.StatusOK)
	}
	if got := w.Body.String(); !strings.Contains(got, "OK") {
		t.Errorf("got body = %s\twant contains %s", got, "OK")
	}
}
```

