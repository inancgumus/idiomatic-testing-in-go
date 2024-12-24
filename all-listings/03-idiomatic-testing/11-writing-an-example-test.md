# Listing 3.11: Writing an example test

## [testing](https://github.com/inancgumus/gobyexample/blob/dc32c85b013a9210934266ae00a340c5a1c0cedb/testing) / [url](https://github.com/inancgumus/gobyexample/blob/dc32c85b013a9210934266ae00a340c5a1c0cedb/testing/url) / [example_test.go](https://github.com/inancgumus/gobyexample/blob/dc32c85b013a9210934266ae00a340c5a1c0cedb/testing/url/example_test.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package url_test

import (
	"fmt"
	"log"

	"github.com/inancgumus/gobyexample/testing/url"
)

func ExampleParse() {
	u, err := url.Parse("http://go.dev/play")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
	// Output:
	// http://go.dev/play
}
```

