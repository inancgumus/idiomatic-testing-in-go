# Listing 7.2: Implementing Summarize

## [hit](https://github.com/inancgumus/gobyexample/blob/4ae2822f0f76417d0a6eb0651ff0b3856baa4bca/hit) / [result.go](https://github.com/inancgumus/gobyexample/blob/4ae2822f0f76417d0a6eb0651ff0b3856baa4bca/hit/result.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package hit

import (
	"iter"
	"net/http"
	"time"
)

// Result is performance metrics of a single [http.Request].
//
// Its zero value is useful, allowing it to be directly used.
type Result struct {
	Status   int           // Status is the HTTP status code
	Bytes    int64         // Bytes is the number of bytes transferred
	Duration time.Duration // Duration is the time to complete the request
	Error    error         // Error from the request
}

// Summary is the performance metrics of multiple requests.
//
// Its zero value is useful, allowing it to be directly used.
type Summary struct {
	Started  time.Time     // Started is the time when the requests began
	Requests int           // Requests is the total number of requests made
	Errors   int           // Errors is the total number of failed requests
	Bytes    int64         // Bytes is the total number of bytes transferred
	RPS      float64       // RPS is the requests per second
	Duration time.Duration // Duration is the total time since the requests started
	Fastest  time.Duration // Fastest is the fastest request duration
	Slowest  time.Duration // Slowest is the slowest request duration
}

// Results is an iterator for [Result] values.
type Results iter.Seq[Result]

// Summarize returns a [Summary] of [Results].
func (r Results) Summarize() Summary {
	return Summarize(r)
}

// Summarize returns a [Summary] of [Results].
func Summarize(results Results) Summary {
	var s Summary
	if results == nil {
		return s
	}

	s.Started = time.Now()

	for r := range results {
		s.Requests++
		s.Bytes += r.Bytes

		if r.Error != nil || r.Status != http.StatusOK {
			s.Errors++
		}
		if s.Fastest == 0 {
			s.Fastest = r.Duration
		}
		if r.Duration < s.Fastest {
			s.Fastest = r.Duration
		}
		if r.Duration > s.Slowest {
			s.Slowest = r.Duration
		}
	}

	s.Duration = time.Since(s.Started)
	s.RPS = float64(s.Requests) / s.Duration.Seconds()

	return s
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -1,7 +1,11 @@
 package hit
 
-import "time"
+import (
+	"iter"
+	"net/http"
+	"time"
+)
 
 // Result is performance metrics of a single [http.Request].
 //
 // Its zero value is useful, allowing it to be directly used.
@@ -25,3 +29,44 @@ type Summary struct {
 	Fastest  time.Duration // Fastest is the fastest request duration
 	Slowest  time.Duration // Slowest is the slowest request duration
 }
+
+// Results is an iterator for [Result] values.
+type Results iter.Seq[Result]
+
+// Summarize returns a [Summary] of [Results].
+func (r Results) Summarize() Summary {
+	return Summarize(r)
+}
+
+// Summarize returns a [Summary] of [Results].
+func Summarize(results Results) Summary {
+	var s Summary
+	if results == nil {
+		return s
+	}
+
+	s.Started = time.Now()
+
+	for r := range results {
+		s.Requests++
+		s.Bytes += r.Bytes
+
+		if r.Error != nil || r.Status != http.StatusOK {
+			s.Errors++
+		}
+		if s.Fastest == 0 {
+			s.Fastest = r.Duration
+		}
+		if r.Duration < s.Fastest {
+			s.Fastest = r.Duration
+		}
+		if r.Duration > s.Slowest {
+			s.Slowest = r.Duration
+		}
+	}
+
+	s.Duration = time.Since(s.Started)
+	s.RPS = float64(s.Requests) / s.Duration.Seconds()
+
+	return s
+}
```

