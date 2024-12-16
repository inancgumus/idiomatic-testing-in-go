# Listing 2.1: Implementing the book package

## [packages](https://github.com/inancgumus/gobyexample/blob/d5055fa6ce1c455c50cb14b26a9bb3e5f6a831b0/packages) / [book](https://github.com/inancgumus/gobyexample/blob/d5055fa6ce1c455c50cb14b26a9bb3e5f6a831b0/packages/book) / [book.go](https://github.com/inancgumus/gobyexample/blob/d5055fa6ce1c455c50cb14b26a9bb3e5f6a831b0/packages/book/book.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
// Package book offers information about the Go by Example book.
package book

// Title returns the title of this book.
func Title() string {
	return "Go by Example: " + subtitle()
}

// subtitle returns the subtitle of this book.
func subtitle() string {
	return "Programmer's Guide to Idiomatic and Testable Code"
}
```

