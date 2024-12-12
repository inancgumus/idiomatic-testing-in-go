# Listing 2.9: Implementing concrete types

## [oop](https://github.com/inancgumus/gobyexample/blob/889e8723420a840cf9cf26c1b4647737d050abac/oop) / [interfaces](https://github.com/inancgumus/gobyexample/blob/889e8723420a840cf9cf26c1b4647737d050abac/oop/interfaces) / [notify.go](https://github.com/inancgumus/gobyexample/blob/889e8723420a840cf9cf26c1b4647737d050abac/oop/interfaces/notify.go)

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

