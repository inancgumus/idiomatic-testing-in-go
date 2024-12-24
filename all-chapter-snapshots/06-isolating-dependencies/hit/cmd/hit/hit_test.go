package main

import (
	"bytes"
	"testing"
)

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

func TestRunValidInput(t *testing.T) {
	t.Parallel()

	e, err := testRun("http://go.dev")
	if err != nil {
		t.Fatalf("got %q;\nwant nil err", err)
	}
	if n := e.stdout.Len(); n == 0 {
		t.Errorf("stdout = 0 bytes; want >0")
	}
	if n, out := e.stderr.Len(), e.stderr.String(); n != 0 {
		t.Errorf("stderr = %d bytes; want 0; stderr:\n%s", n, out)
	}
}

func TestRunInvalidInput(t *testing.T) {
	t.Parallel()

	e, err := testRun("-c=2", "-n=1", "invalid-url")
	if err == nil {
		t.Fatalf("got nil; want err")
	}
	if n := e.stderr.Len(); n == 0 {
		t.Error("stderr = 0 bytes; want >0")
	}
}
