package main

import (
	"fmt"
)

// slackNotifier notifies a server's status via Slack.
type slackNotifier struct {
	apiKey  string //nolint:unused
	channel string //nolint:unused
	// ..other fields
}

// notify sends a message via Slack.
func (s *slackNotifier) notify(msg string) {
	// simulate sending a message to Slack
	fmt.Println("slack:", msg)
}

// disconnect disconnects from Slack.
func (s *slackNotifier) disconnect() { //nolint:unused
	fmt.Println("slack: disconnecting")
}

// smsNotifier notifies a server's status via SMS.
type smsNotifier struct {
	gatewayIP string //nolint:unused
	// ..other fields
}

// notify sends a message via SMS.
func (s *smsNotifier) notify(msg string) {
	// simulate sending a message via SMS
	fmt.Println("sms:", msg)
}

// --------------------------------------------------------------------------------

// notifier notifies a server's status to an external server.
type notifier interface {
	notify(message string)
}

// notify sends a message via a notifier if the server is slow.
func notify(s *server, n notifier) {
	if !s.slow() {
		return
	}

	msg := fmt.Sprintf(
		"%s server is slow: %s",
		s.url, s.responseTime,
	)
	n.notify(msg)

	// n.disconnect() // incorrect: n is not a slackNotifier
}

// --------------------------------------------------------------------------------

// multiNotifier notifies a server's status via multiple notifiers.
type multiNotifier []notifier

// notify sends a message via multiple notifiers.
func (m multiNotifier) notify(msg string) {
	for _, n := range m {
		n.notify(msg)
	}
}
