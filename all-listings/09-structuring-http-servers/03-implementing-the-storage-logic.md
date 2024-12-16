# Listing 9.3: Implementing the storage logic

## [bite](https://github.com/inancgumus/gobyexample/blob/a91b722b336c6353386ff14bf1c90f9a8ea46d2f/bite) / [link](https://github.com/inancgumus/gobyexample/blob/a91b722b336c6353386ff14bf1c90f9a8ea46d2f/bite/link) / [store.go](https://github.com/inancgumus/gobyexample/blob/a91b722b336c6353386ff14bf1c90f9a8ea46d2f/bite/link/store.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package link

// Store persists and retrieves [Link] values in an in-memory map.
type Store struct {
	links map[string]Link
}
```

