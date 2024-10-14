package main

import "flag"

type config struct {
	url string // url to send requests
	n   int    // n is the number of requests
	c   int    // c is the concurrency level
	rps int    // rps is the requests per second
}

func parseArgs(c *config, args []string) error {
	fs := flag.NewFlagSet("hit", flag.ContinueOnError)

	fs.StringVar(&c.url, "url", "", "HTTP server URL ")
	fs.IntVar(&c.n, "n", c.n, "Number of requests")
	fs.IntVar(&c.c, "c", c.c, "Concurrency level")
	fs.IntVar(&c.rps, "rps", c.rps, "Requests per second")

	return fs.Parse(args)
}
