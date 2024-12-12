# Listing 2.14: Notifying with multiple notifiers

## [oop](https://github.com/inancgumus/gobyexample/blob/3ad25609313b45c852d670824558988c797985b0/oop) / [interfaces](https://github.com/inancgumus/gobyexample/blob/3ad25609313b45c852d670824558988c797985b0/oop/interfaces) / [main.go](https://github.com/inancgumus/gobyexample/blob/3ad25609313b45c852d670824558988c797985b0/oop/interfaces/main.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	srv := &server{url: "auth", responseTime: time.Minute}
	srv.check()
	if !srv.slow() {
		return
	}
	notify(
		multiNotifier{new(slackNotifier), new(smsNotifier)},
		fmt.Sprintf("%s server is slow: %s", srv.url, srv.responseTime),
	)
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -11,7 +11,8 @@ func main() {
 	if !srv.slow() {
 		return
 	}
-	msg := fmt.Sprintf("%s server is slow: %s", srv.url, srv.responseTime)
-	notify(new(slackNotifier), msg)
-	notify(new(smsNotifier), msg)
+	notify(
+		multiNotifier{new(slackNotifier), new(smsNotifier)},
+		fmt.Sprintf("%s server is slow: %s", srv.url, srv.responseTime),
+	)
 }
```

