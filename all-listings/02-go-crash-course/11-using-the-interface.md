# Listing 2.11: Using the interface

## [oop](https://github.com/inancgumus/gobyexample/blob/08dcd85b1002f54fdc9f503475eb8a87d5260955/oop) / [interfaces](https://github.com/inancgumus/gobyexample/blob/08dcd85b1002f54fdc9f503475eb8a87d5260955/oop/interfaces) / [notify.go](https://github.com/inancgumus/gobyexample/blob/08dcd85b1002f54fdc9f503475eb8a87d5260955/oop/interfaces/notify.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import "fmt"

// Concrete notifier types.

type slackNotifier struct{ apiKey string }

func (s *slackNotifier) notify(msg string) { fmt.Println("slack:", msg) }
func (s *slackNotifier) disconnect()       { fmt.Println("slack: bye!") }

type smsNotifier struct{ gatewayIP string }

func (s *smsNotifier) notify(msg string) { fmt.Println("sms:", msg) }

// Abstract notifier interface to represent the notification behavior.

type notifier interface {
	notify(message string)
}

// notify sends a message to the given notifier.
// It doesn't matter what the concrete type is.
func notify(n notifier, msg string) {
	n.notify(msg)
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -18,3 +18,9 @@ func (s *smsNotifier) notify(msg string) { fmt.Println("sms:", msg) }
 type notifier interface {
 	notify(message string)
 }
+
+// notify sends a message to the given notifier.
+// It doesn't matter what the concrete type is.
+func notify(n notifier, msg string) {
+	n.notify(msg)
+}
```

