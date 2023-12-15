package link

// ---------------------------------------------------------
// Exercise: Implement an HTTP client for the URL shortener
// server.
//
// - Implement a client for the server's all endpoints.
// - The client should be able to shorten and resolve URLs
//   by sending HTTP requests to the server.
//
// Tip: You can see the client's proposed API in this file.
//
// Bonus: Test the client (reach at least 70% test coverage).
// Bonus: Add a timeout option and support.
// Bonus: Add a retry mechanism.
//
// Warning: Avoid using the server's HTTP handler functions in the client.
// These exercises are about implementing a client for the server using the
// standard library's net/http package.
// ---------------------------------------------------------

// ---------------------------------------------------------
// Exercise: Create a program called link that uses the
// client you implemented in the previous exercise to shorten
// and resolve URLs.
//
// Implement the following operations:
//
// 	- link -s <url> # shortens the given URL
// 	- link -r <key> # resolves the short URL
// 	- link health   # checks the health of the server
//
// Optional flags:
// - -addr: The address of the server. Default: localhost:8080
// - -timeout: The timeout for the client. Default: 10 seconds
// - -retry: The number of retries for the client. Default: 3
// - -retry-wait: The wait time between retries. Default: 1 second
// - -log: The log level. Default: info
// ---------------------------------------------------------

// Client is a URL short HTTP client.
type Client struct {
	// *http.Client
}

// NewClient creates and returns a new Client.
func NewClient() *Client {
	return nil
}

// Shorten shortens the given URL and returns a [Link].
func (c *Client) Shorten(url string) (Link, error) {
	return Link{}, nil
}

// Resolve resolves the given short URL and returns a [Link].
func (c *Client) Resolve(key string) (Link, error) {
	return Link{}, nil
}

// Health checks the health of the server.
func (c *Client) Health() error {
	return nil
}
