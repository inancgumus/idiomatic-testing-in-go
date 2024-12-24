# Listing 2.10: Discovering the interface

## [oop](https://github.com/inancgumus/gobyexample/blob/4f09e4cd2830fae66b41df1769ea943592f10b0d/oop) / [interfaces](https://github.com/inancgumus/gobyexample/blob/4f09e4cd2830fae66b41df1769ea943592f10b0d/oop/interfaces) / [notify.go](https://github.com/inancgumus/gobyexample/blob/4f09e4cd2830fae66b41df1769ea943592f10b0d/oop/interfaces/notify.go)

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
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -12,3 +12,9 @@ func (s *slackNotifier) disconnect()       { fmt.Println("slack: bye!") }
 type smsNotifier struct{ gatewayIP string }
 
 func (s *smsNotifier) notify(msg string) { fmt.Println("sms:", msg) }
+
+// Abstract notifier interface to represent the notification behavior.
+
+type notifier interface {
+	notify(message string)
+}
```

