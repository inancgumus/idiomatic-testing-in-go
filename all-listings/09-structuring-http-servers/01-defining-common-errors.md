# Listing 9.1: Defining common errors

## [bite](https://github.com/inancgumus/gobyexample/blob/6bab010d876ca61098d3875285e4bc876752625b/bite) / [errors.go](https://github.com/inancgumus/gobyexample/blob/6bab010d876ca61098d3875285e4bc876752625b/bite/errors.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package bite

import "errors"

var (
	ErrExists         = errors.New("already exists")
	ErrNotExist       = errors.New("does not exist")
	ErrInvalidRequest = errors.New("invalid request")
	ErrInternal       = errors.New(
		"internal error: please try again later or contact support",
	)
)

// other shared types—structs, interfaces, etc., maybe in different files.
```

