package link

import "net/http"

// Client is a URL shortener client.
//
// # Exercise #1
//
// Implement an HTTP client for the URL shortener server.
//
//   - Implement the client for the all endpoints of [Server].
//   - The client should be able to shorten and resolve URLs
//     by sending HTTP requests to the server.
//
// Tip: You can see the client's proposed API in this file.
//
//   - Bonus: Add a timeout option and support.
//   - Bonus: Add a retry mechanism.
//
// Warning:
//
// Avoid using the server's HTTP handler functions in the
// client. These exercises are about implementing a client
// for the server using the standard library's net/http
// package.
//
// # Exercise #2
//
// Create a program called shortlink that uses [Client]
// with the following command-line interface:
//
//   - shortlink -s <url> # shortens the given URL
//   - shortlink -r <key> # resolves the short URL
//   - shortlink health   # checks the health of the server
//
// Optional flags:
//
//   - -addr: The address of the server. Default: localhost:8080
//   - -timeout: The timeout for the client. Default: 10 seconds
//   - -retry: The number of retries for the client. Default: 3
//   - -retry-wait: The wait time between retries. Default: 1 second
//   - -log: The log level. Default: info
type Client struct{ client *http.Client } //nolint:unused

// NewClient creates and returns a new [Client].
func NewClient() *Client { return nil }

// Shorten shortens the given URL and returns a [Link].
func (c *Client) Shorten(url string) (Link, error) { return Link{}, nil }

// Resolve resolves the given short URL and returns a [Link].
func (c *Client) Resolve(key string) (Link, error) { return Link{}, nil }

// Health checks the health of the server.
func (c *Client) Health() error { return nil }
