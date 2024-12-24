# Listing 2.12: Gluing everything together

## [oop](https://github.com/inancgumus/gobyexample/blob/9d5ddfdbb6234b772885b4382f8ab5ce28db88ac/oop) / [interfaces](https://github.com/inancgumus/gobyexample/blob/9d5ddfdbb6234b772885b4382f8ab5ce28db88ac/oop/interfaces) / [main.go](https://github.com/inancgumus/gobyexample/blob/9d5ddfdbb6234b772885b4382f8ab5ce28db88ac/oop/interfaces/main.go)

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
	msg := fmt.Sprintf("%s server is slow: %s", srv.url, srv.responseTime)
	notify(new(slackNotifier), msg)
	notify(new(smsNotifier), msg)
}
```

## [oop](https://github.com/inancgumus/gobyexample/blob/9d5ddfdbb6234b772885b4382f8ab5ce28db88ac/oop) / [interfaces](https://github.com/inancgumus/gobyexample/blob/9d5ddfdbb6234b772885b4382f8ab5ce28db88ac/oop/interfaces) / [server.go](https://github.com/inancgumus/gobyexample/blob/9d5ddfdbb6234b772885b4382f8ab5ce28db88ac/oop/interfaces/server.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import "time"

type server struct {
	url          string
	responseTime time.Duration
}

func (s *server) check()     { s.responseTime = 3 * time.Second }
func (s *server) slow() bool { return s.responseTime > 2*time.Second }
```

