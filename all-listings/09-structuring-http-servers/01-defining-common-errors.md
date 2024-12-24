# Listing 9.1: Defining common errors

## [bite](https://github.com/inancgumus/gobyexample/blob/456db62cd0afbc94dc44e4b82c67bfed7c7c9d88/bite) / [error.go](https://github.com/inancgumus/gobyexample/blob/456db62cd0afbc94dc44e4b82c67bfed7c7c9d88/bite/error.go)

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
	ErrInternal       = errors.New("internal error")
)

// other shared types—structs, interfaces, etc., maybe in different files.
```

