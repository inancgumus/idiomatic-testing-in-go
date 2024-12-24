# Listing 9.3: Implementing the in-memory store

## [bite](https://github.com/inancgumus/gobyexample/blob/822d3e9809c5df00674812eaaf86749aa25c655a/bite) / [link](https://github.com/inancgumus/gobyexample/blob/822d3e9809c5df00674812eaaf86749aa25c655a/bite/link) / [store.go](https://github.com/inancgumus/gobyexample/blob/822d3e9809c5df00674812eaaf86749aa25c655a/bite/link/store.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package link

// Store persists and retrieves [Link] values in an in-memory map.
type Store struct {
	links map[LinkKey]Link
}
```

