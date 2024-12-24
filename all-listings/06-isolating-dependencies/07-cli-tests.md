# Listing 6.7: CLI tests

## [hit](https://github.com/inancgumus/gobyexample/blob/126f430c076a69dba7f7d8b46ea733f3abaef2ae/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/126f430c076a69dba7f7d8b46ea733f3abaef2ae/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/126f430c076a69dba7f7d8b46ea733f3abaef2ae/hit/cmd/hit) / [hit_test.go](https://github.com/inancgumus/gobyexample/blob/126f430c076a69dba7f7d8b46ea733f3abaef2ae/hit/cmd/hit/hit_test.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
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
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -1,6 +1,9 @@
 package main
 
-import "bytes"
+import (
+	"bytes"
+	"testing"
+)
 
 type testEnv struct {
 	env    env
@@ -18,3 +21,30 @@ func testRun(args ...string) (*testEnv, error) {
 	}
 	return &t, run(&t.env)
 }
+
+func TestRunValidInput(t *testing.T) {
+	t.Parallel()
+
+	e, err := testRun("http://go.dev")
+	if err != nil {
+		t.Fatalf("got %q;\nwant nil err", err)
+	}
+	if n := e.stdout.Len(); n == 0 {
+		t.Errorf("stdout = 0 bytes; want >0")
+	}
+	if n, out := e.stderr.Len(), e.stderr.String(); n != 0 {
+		t.Errorf("stderr = %d bytes; want 0; stderr:\n%s", n, out)
+	}
+}
+
+func TestRunInvalidInput(t *testing.T) {
+	t.Parallel()
+
+	e, err := testRun("-c=2", "-n=1", "invalid-url")
+	if err == nil {
+		t.Fatalf("got nil; want err")
+	}
+	if n := e.stderr.Len(); n == 0 {
+		t.Error("stderr = 0 bytes; want >0")
+	}
+}
```

