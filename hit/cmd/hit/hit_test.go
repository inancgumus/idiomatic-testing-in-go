package main

import (
	"bytes"
	"testing"
)

// testEnv is a test environment for running the program
// in a controlled environment.
type testEnv struct {
	env    env          // the environment
	stdout bytes.Buffer // captured standard output
	stderr bytes.Buffer // captured standard error
}

// testRun runs the tool with the given arguments and returns
// the environment and error. Use stderr and stdout to inspect
// the output.
func testRun(args ...string) (*testEnv, error) {
	var t testEnv // use bytes.Buffer zero values
	t.env = env{
		args:   append([]string{"hit"}, args...), // prepend the command
		stdout: &t.stdout,                        // capture the output
		stderr: &t.stderr,                        // capture the errors
		dryRun: true,                             // don't make requests
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
