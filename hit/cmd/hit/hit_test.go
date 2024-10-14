package main

import "bytes"

type testEnv struct {
	env    env
	stdout bytes.Buffer
	stderr bytes.Buffer
}

func testRun(args ...string) (*testEnv, error) {
	var t testEnv
	t.env = env{
		args:   append([]string{"hit"}, args...),
		stdout: &t.stdout,
		stderr: &t.stderr,
		dryRun: true,
	}
	return &t, run(&t.env)
}
