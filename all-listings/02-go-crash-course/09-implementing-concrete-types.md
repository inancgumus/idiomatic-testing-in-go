# Listing 2.9: Implementing concrete types

## [oop](https://github.com/inancgumus/gobyexample/blob/533d230c640d9153052d189df97c7074c6b416da/oop) / [interfaces](https://github.com/inancgumus/gobyexample/blob/533d230c640d9153052d189df97c7074c6b416da/oop/interfaces) / [notify.go](https://github.com/inancgumus/gobyexample/blob/533d230c640d9153052d189df97c7074c6b416da/oop/interfaces/notify.go)

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
```

