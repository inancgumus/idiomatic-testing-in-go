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
