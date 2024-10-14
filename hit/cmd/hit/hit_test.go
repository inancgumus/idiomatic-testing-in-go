package main

import "bytes"

type testEnv struct {
	env    env
	stdout bytes.Buffer
	stderr bytes.Buffer
}
