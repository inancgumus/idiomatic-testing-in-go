# Listing 5.3: Implementing the flag value parsers

## [hit](https://github.com/inancgumus/gobyexample/blob/4e3ba0428a8da17e40f3469d91f0f48ced230396/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/4e3ba0428a8da17e40f3469d91f0f48ced230396/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/4e3ba0428a8da17e40f3469d91f0f48ced230396/hit/cmd/hit) / [env.go](https://github.com/inancgumus/gobyexample/blob/4e3ba0428a8da17e40f3469d91f0f48ced230396/hit/cmd/hit/env.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import "strconv"

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
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -1,8 +1,27 @@
 package main
 
+import "strconv"
+
 type config struct {
 	url string // url to send requests
 	n   int    // n is the number of requests
 	c   int    // c is the concurrency level
 	rps int    // rps is the requests per second
 }
+
+type parseFunc func(string) error
+
+func stringVar(p *string) parseFunc {
+	return func(s string) error {
+		*p = s
+		return nil
+	}
+}
+
+func intVar(p *int) parseFunc {
+	return func(s string) error {
+		var err error
+		*p, err = strconv.Atoi(s)
+		return err
+	}
+}
```

