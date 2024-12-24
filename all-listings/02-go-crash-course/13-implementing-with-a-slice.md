# Listing 2.13: Implementing with a slice

## [oop](https://github.com/inancgumus/gobyexample/blob/5ace4e8a2f08a701f4b1dff597f1c267586d5e9c/oop) / [interfaces](https://github.com/inancgumus/gobyexample/blob/5ace4e8a2f08a701f4b1dff597f1c267586d5e9c/oop/interfaces) / [notify.go](https://github.com/inancgumus/gobyexample/blob/5ace4e8a2f08a701f4b1dff597f1c267586d5e9c/oop/interfaces/notify.go)

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

// multiNotifier can send notifications using a group of notifiers.
type multiNotifier []notifier

func (mn multiNotifier) notify(msg string) {
	for _, n := range mn {
		n.notify(msg)
	}
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -24,3 +24,12 @@ type notifier interface {
 func notify(n notifier, msg string) {
 	n.notify(msg)
 }
+
+// multiNotifier can send notifications using a group of notifiers.
+type multiNotifier []notifier
+
+func (mn multiNotifier) notify(msg string) {
+	for _, n := range mn {
+		n.notify(msg)
+	}
+}
```

