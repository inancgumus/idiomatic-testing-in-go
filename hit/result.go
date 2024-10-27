package hit

import "time"

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
