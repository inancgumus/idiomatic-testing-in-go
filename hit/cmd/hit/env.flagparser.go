//go:build exclude

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ------------------------------------------------------------
// custom flag parsing example
// ------------------------------------------------------------

// config holds the program's configuration.
type config struct {
	url string // url to send requests
	n   int    // n is the number of requests
	c   int    // c is the concurrency level
	rps int    // rps is the requests per second
}

// parseFlags parses flags from args and updates c.
// Flags are expected to be in the form of -flag=value.
func parseFlags(c *config, args []string) error {
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

// parseFunc is a function that parses a string and updates a value.
type parseFunc func(string) error

// stringVar returns a parseFunc that updates p.
func stringVar(p *string) parseFunc {
	return func(s string) error {
		*p = s
		return nil
	}
}

// intVar returns a parseFunc that converts s to an int and updates p.
func intVar(p *int) parseFunc {
	return func(s string) error {
		var err error
		*p, err = strconv.Atoi(s)
		return err
	}
}
