# Listing 6.5: Declaring the testEnv type

## [hit](https://github.com/inancgumus/gobyexample/blob/10041fc57f7038933cf71c353b5228ab1f4904fa/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/10041fc57f7038933cf71c353b5228ab1f4904fa/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/10041fc57f7038933cf71c353b5228ab1f4904fa/hit/cmd/hit) / [hit_test.go](https://github.com/inancgumus/gobyexample/blob/10041fc57f7038933cf71c353b5228ab1f4904fa/hit/cmd/hit/hit_test.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import "bytes"

type testEnv struct {
	env    env
	stdout bytes.Buffer
	stderr bytes.Buffer
}
```

