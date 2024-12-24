# Listing 6.8: Adding a shared test case type

## [hit](https://github.com/inancgumus/gobyexample/blob/6d7995393d61764b6244b1741644dc79f62e658c/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/6d7995393d61764b6244b1741644dc79f62e658c/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/6d7995393d61764b6244b1741644dc79f62e658c/hit/cmd/hit) / [env_test.go](https://github.com/inancgumus/gobyexample/blob/6d7995393d61764b6244b1741644dc79f62e658c/hit/cmd/hit/env_test.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

type parseArgsTest struct {
	name string
	args []string
	want config
}
```

