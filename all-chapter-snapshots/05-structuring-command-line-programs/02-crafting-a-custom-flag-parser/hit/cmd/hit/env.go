package main

import (
	"fmt"
	"strconv"
	"strings"
)

type config struct {
	url string // url to send requests
	n   int    // n is the number of requests
	c   int    // c is the concurrency level
	rps int    // rps is the requests per second
}

type parseFunc func(string) error

func stringVar(p *string) parseFunc {
	return func(s string) error {
		*p = s
		return nil
	}
}

func intVar(p *int) parseFunc {
	return func(s string) error {
		var err error
		*p, err = strconv.Atoi(s)
		return err
	}
}

func parseArgs(c *config, args []string) error {
	flagSet := map[string]parseFunc{
		"url": stringVar(&c.url),
		"n":   intVar(&c.n),
		"c":   intVar(&c.c),
		"rps": intVar(&c.rps),
	}
	for _, arg := range args {
		fn, fv, ok := strings.Cut(arg, "=")
		if !ok {
			continue // wrong flag format
		}
		fn = strings.TrimPrefix(fn, "-")
		parseValue, ok := flagSet[fn]
		if !ok {
			continue // not in flagSet
		}
		if err := parseValue(fv); err != nil {
			return fmt.Errorf("invalid value %q for flag %s: %w", fv, fn, err)
		}
	}
	return nil
}
