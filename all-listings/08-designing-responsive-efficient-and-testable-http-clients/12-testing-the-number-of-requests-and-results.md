# Listing 8.12: Testing the number of requests and results

## [hit](https://github.com/inancgumus/gobyexample/blob/5b79521ddfae7e173149f23ef86f35a5a30477a7/hit) / [hit_test.go](https://github.com/inancgumus/gobyexample/blob/5b79521ddfae7e173149f23ef86f35a5a30477a7/hit/hit_test.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package hit

import (
	"context"
	"net/http"
	"sync/atomic"
	"testing"
)

// roundTripperFunc is an adapter to allow the use of ordinary functions as an
// [http.RoundTripper]. If f is a function with the appropriate signature,
// roundTripperFunc(f) is an [http.RoundTripper] that calls f.
type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func TestSendNRequestResultCount(t *testing.T) {
	t.Parallel()

	const wantRequests = 10

	var totalRequests atomic.Int64
	countingRoundTripper := roundTripperFunc(
		func(*http.Request) (*http.Response, error) {
			totalRequests.Add(1)
			return &http.Response{StatusCode: http.StatusOK}, nil
		},
	)

	req, err := http.NewRequest(http.MethodGet, "/any", http.NoBody)
	if err != nil {
		t.Fatalf("http.NewRequest() err = %q, want <nil>", err)
	}
	results, err := SendN(context.Background(), wantRequests, req, &Options{
		Client:      &http.Client{Transport: countingRoundTripper},
		Concurrency: 2,
	})
	if err != nil {
		t.Fatalf("SendN() err = %q, want <nil>", err)
	}

	var totalResults int
	for range results {
		totalResults++
	}

	if got := totalRequests.Load(); got != wantRequests {
		t.Fatalf("totalRequests = %d, want %d", got, wantRequests)
	}
	if totalResults != wantRequests {
		t.Fatalf("totalResults = %d, want %d", totalResults, wantRequests)
	}
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -1,6 +1,11 @@
 package hit
 
-import "net/http"
+import (
+	"context"
+	"net/http"
+	"sync/atomic"
+	"testing"
+)
 
 // roundTripperFunc is an adapter to allow the use of ordinary functions as an
 // [http.RoundTripper]. If f is a function with the appropriate signature,
@@ -10,3 +15,41 @@ type roundTripperFunc func(*http.Request) (*http.Response, error)
 func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
 	return f(r)
 }
+
+func TestSendNRequestResultCount(t *testing.T) {
+	t.Parallel()
+
+	const wantRequests = 10
+
+	var totalRequests atomic.Int64
+	countingRoundTripper := roundTripperFunc(
+		func(*http.Request) (*http.Response, error) {
+			totalRequests.Add(1)
+			return &http.Response{StatusCode: http.StatusOK}, nil
+		},
+	)
+
+	req, err := http.NewRequest(http.MethodGet, "/any", http.NoBody)
+	if err != nil {
+		t.Fatalf("http.NewRequest() err = %q, want <nil>", err)
+	}
+	results, err := SendN(context.Background(), wantRequests, req, &Options{
+		Client:      &http.Client{Transport: countingRoundTripper},
+		Concurrency: 2,
+	})
+	if err != nil {
+		t.Fatalf("SendN() err = %q, want <nil>", err)
+	}
+
+	var totalResults int
+	for range results {
+		totalResults++
+	}
+
+	if got := totalRequests.Load(); got != wantRequests {
+		t.Fatalf("totalRequests = %d, want %d", got, wantRequests)
+	}
+	if totalResults != wantRequests {
+		t.Fatalf("totalResults = %d, want %d", totalResults, wantRequests)
+	}
+}
```

