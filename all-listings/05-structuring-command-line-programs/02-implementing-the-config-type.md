# Listing 5.2: Implementing the config type

## [hit](https://github.com/inancgumus/gobyexample/blob/10a2656d616ab69398e104a63ce64ab7e026b7a0/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/10a2656d616ab69398e104a63ce64ab7e026b7a0/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/10a2656d616ab69398e104a63ce64ab7e026b7a0/hit/cmd/hit) / [env.go](https://github.com/inancgumus/gobyexample/blob/10a2656d616ab69398e104a63ce64ab7e026b7a0/hit/cmd/hit/env.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

type config struct {
	url string // url to send requests
	n   int    // n is the number of requests
	c   int    // c is the concurrency level
	rps int    // rps is the requests per second
}
```

