# Listing 2.24: Listening to multiple channels

## [concurrency](https://github.com/inancgumus/gobyexample/blob/7f9efb4a0ed5a91cb507b25fb8c127652549b6f2/concurrency) / [barrier](https://github.com/inancgumus/gobyexample/blob/7f9efb4a0ed5a91cb507b25fb8c127652549b6f2/concurrency/barrier) / [main.go](https://github.com/inancgumus/gobyexample/blob/7f9efb4a0ed5a91cb507b25fb8c127652549b6f2/concurrency/barrier/main.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import (
	"fmt"

	"github.com/inancgumus/gobyexample/concurrency/syncx"
)

func main() {
	var sg syncx.SafeGroup

	wait := make(chan struct{})
	stop := make(chan struct{})
	for range 10 {
		sg.Go(func() {
			select {
			case <-wait:
			case <-stop:
				fmt.Print("stopped.")
				return
			}
			fmt.Print("go!")
		})
	}
	// do other work
	close(stop)
	// Either close the stop channel or the wait channel
	// depending on your program's logic.
	sg.Wait()
	// do other work
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -10,14 +10,22 @@ func main() {
 	var sg syncx.SafeGroup
 
 	wait := make(chan struct{})
+	stop := make(chan struct{})
 	for range 10 {
 		sg.Go(func() {
-			<-wait
+			select {
+			case <-wait:
+			case <-stop:
+				fmt.Print("stopped.")
+				return
+			}
 			fmt.Print("go!")
 		})
 	}
 	// do other work
-	close(wait)
+	close(stop)
+	// Either close the stop channel or the wait channel
+	// depending on your program's logic.
 	sg.Wait()
 	// do other work
 }
```

