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
