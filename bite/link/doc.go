// Package link provides a link management [Server] and its [Client].
//
//   - The [Server] shortens URLs and resolves shortened URLs.
//   - The [Client] can shorten and resolve URLs. However, it is not
//     implemented yet and is left as an exercise.
//
// # Endpoints
//
// The service provides three handlers:
//   - [Shorten] - Shortens a URL.
//   - [Resolve] - Resolves a shortened URL.
//   - [Health] - Checks the health of the service.
//
// # Curl examples
//
// Shorten a URL:
//
//	$ curl localhost:8080/shorten -d '{"key":"inanc", "url":"https://x.com/inancgumus"}'
//
// Resolve a shortened URL:
//
//	$ curl localhost:8080/r/inanc
//
// Health check:
//
//	$ curl localhost:8080/health
package link
