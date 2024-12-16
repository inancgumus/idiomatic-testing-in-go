# Listing 8.1: Setting a default http.Client

## [hit](https://github.com/inancgumus/gobyexample/blob/4b641e6b36ae586d679a614454f14989ea357382/hit) / [option.go](https://github.com/inancgumus/gobyexample/blob/4b641e6b36ae586d679a614454f14989ea357382/hit/option.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package hit

import "net/http"

// SendFunc is a type of function that sends a request
// and returns a result.
type SendFunc func(*http.Request) Result

// Options defines the options for sending requests.
type Options struct {
	// Concurrency is the number of concurrent requests to send.
	// Default: 1
	Concurrency int

	// RPS is the requests to send per second.
	// Default: 0 (no rate limiting)
	RPS int

	// Send allows customizing the request processing.
	// Default: Uses [Send].
	Send SendFunc

	// Client is the HTTP client to use for sending requests.
	// Default: http.DefaultClient
	Client *http.Client
}

// DefaultOptions returns a new [Options] with the defaults.
func DefaultOptions() *Options {
	return new(Options).setDefaults()
}

// setDefaults sets the defaults.
// It returns a new [Options] if the current one is nil.
func (o *Options) setDefaults() *Options {
	if o == nil {
		o = new(Options)
	}
	if o.Concurrency <= 0 {
		o.Concurrency = 1
	}
	if o.Client == nil {
		o.Client = http.DefaultClient
	}
	if o.Send == nil {
		o.Send = func(r *http.Request) Result {
			return Send(o.Client, r)
		}
	}
	return o
}

// clone returns a shallow copy of the [Options].
// It returns nil if the current [Options] is nil.
func (o *Options) clone() *Options {
	if o == nil {
		return nil
	}
	c := *o
	return &c
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -19,6 +19,10 @@ type Options struct {
 	// Send allows customizing the request processing.
 	// Default: Uses [Send].
 	Send SendFunc
+
+	// Client is the HTTP client to use for sending requests.
+	// Default: http.DefaultClient
+	Client *http.Client
 }
 
 // DefaultOptions returns a new [Options] with the defaults.
@@ -35,8 +39,13 @@ func (o *Options) setDefaults() *Options {
 	if o.Concurrency <= 0 {
 		o.Concurrency = 1
 	}
+	if o.Client == nil {
+		o.Client = http.DefaultClient
+	}
 	if o.Send == nil {
-		o.Send = Send
+		o.Send = func(r *http.Request) Result {
+			return Send(o.Client, r)
+		}
 	}
 	return o
 }
```

