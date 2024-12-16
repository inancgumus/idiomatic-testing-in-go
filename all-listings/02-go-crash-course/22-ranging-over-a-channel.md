# Listing 2.22: Ranging over a channel

## [concurrency](https://github.com/inancgumus/gobyexample/blob/93b7866e465851ae50cf70ecd78d59adec021adc/concurrency) / [forrange](https://github.com/inancgumus/gobyexample/blob/93b7866e465851ae50cf70ecd78d59adec021adc/concurrency/forrange) / [main.go](https://github.com/inancgumus/gobyexample/blob/93b7866e465851ae50cf70ecd78d59adec021adc/concurrency/forrange/main.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	results := make(chan int)
	go func() {
		for n := range rand.N(100) + 1 { //nolint:gosec
			results <- max(1, n*2)
		}
		close(results)
	}()
	for result := range results {
		fmt.Print(result, ".")
	}
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -13,11 +13,7 @@ func main() {
 		}
 		close(results)
 	}()
-	for {
-		result, ok := <-results
-		if !ok {
-			break
-		}
+	for result := range results {
 		fmt.Print(result, ".")
 	}
 }
```

