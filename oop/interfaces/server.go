package main

import "time"

type server struct {
	url          string
	responseTime time.Duration
}

func (s *server) check()     { s.responseTime = 3 * time.Second }
func (s *server) slow() bool { return s.responseTime > 2*time.Second }
