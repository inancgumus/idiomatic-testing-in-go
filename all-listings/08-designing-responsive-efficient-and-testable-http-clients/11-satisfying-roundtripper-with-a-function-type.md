# Listing 8.11: Satisfying RoundTripper with a function type

## [hit](https://github.com/inancgumus/gobyexample/blob/fd1f11149a066dd093c0c526eb83d4851e1aeb76/hit) / [hit_test.go](https://github.com/inancgumus/gobyexample/blob/fd1f11149a066dd093c0c526eb83d4851e1aeb76/hit/hit_test.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package hit

import "net/http"

// roundTripperFunc is an adapter to allow the use of ordinary functions as an
// [http.RoundTripper]. If f is a function with the appropriate signature,
// roundTripperFunc(f) is an [http.RoundTripper] that calls f.
type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}
```

