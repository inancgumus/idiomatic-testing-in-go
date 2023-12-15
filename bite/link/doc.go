// Package link provides a link management service.
//
// # Endpoints
//
// The service provides two endpoints:
//   - /shorten for shortening URLs.
//   - /r/{key} for resolving shortened URLs.
//   - /health for health checking the service.
//
// # Curl examples
//
// Shorten a URL:
//
//	$ curl localhost:8080/shorten -d '{"key":"go", "url":"https://go.dev"}'
//
// Resolve a shortened URL:
//
//	$ curl localhost:8080/r/go
//
// Health check:
//
//	$ curl localhost:8080/health
package link
