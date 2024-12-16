# Listing 6.6: Declaring the testRun function

## [hit](https://github.com/inancgumus/gobyexample/blob/5671b2765c3551dd798947e36ba53a41f9fcf5bb/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/5671b2765c3551dd798947e36ba53a41f9fcf5bb/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/5671b2765c3551dd798947e36ba53a41f9fcf5bb/hit/cmd/hit) / [hit_test.go](https://github.com/inancgumus/gobyexample/blob/5671b2765c3551dd798947e36ba53a41f9fcf5bb/hit/cmd/hit/hit_test.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
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
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -7,3 +7,14 @@ type testEnv struct {
 	stdout bytes.Buffer
 	stderr bytes.Buffer
 }
+
+func testRun(args ...string) (*testEnv, error) {
+	var t testEnv
+	t.env = env{
+		args:   append([]string{"hit"}, args...),
+		stdout: &t.stdout,
+		stderr: &t.stderr,
+		dryRun: true,
+	}
+	return &t, run(&t.env)
+}
```

