# Listing 6.9: Testing with valid flags

## [hit](https://github.com/inancgumus/gobyexample/blob/f1f7de0cdc5aec8ed10f1cb46e1dc48c83aa546b/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/f1f7de0cdc5aec8ed10f1cb46e1dc48c83aa546b/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/f1f7de0cdc5aec8ed10f1cb46e1dc48c83aa546b/hit/cmd/hit) / [env_test.go](https://github.com/inancgumus/gobyexample/blob/f1f7de0cdc5aec8ed10f1cb46e1dc48c83aa546b/hit/cmd/hit/env_test.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import (
	"io"
	"testing"
)

type parseArgsTest struct {
	name string
	args []string
	want config
}

func TestParseArgsValidInput(t *testing.T) {
	t.Parallel()

	for _, tt := range []parseArgsTest{
		{
			name: "all_flags",
			args: []string{"-n=10", "-c=5", "-rps=5", "http://test"},
			want: config{n: 10, c: 5, rps: 5, url: "http://test"},
		},

		// exercise: test with a mixture of flags
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var got config
			if err := parseArgs(&got, tt.args, io.Discard); err != nil {
				t.Fatalf("parseArgs() error = %v, want no error", err)
			}
			if got != tt.want {
				t.Errorf("flags = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestParseArgsInvalidInput(t *testing.T) {
	t.Parallel()

	for _, tt := range []parseArgsTest{
		{name: "n_syntax", args: []string{"-n=ONE", "http://test"}},
		{name: "n_zero", args: []string{"-n=0", "http://test"}},
		{name: "n_negative", args: []string{"-n=-1", "http://test"}},

		// exercise: test other error conditions
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := parseArgs(&config{}, tt.args, io.Discard)
			if err == nil {
				t.Fatal("parseArgs() = nil, want error")
			}
		})
	}
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -1,7 +1,59 @@
 package main
 
+import (
+	"io"
+	"testing"
+)
+
 type parseArgsTest struct {
 	name string
 	args []string
 	want config
 }
+
+func TestParseArgsValidInput(t *testing.T) {
+	t.Parallel()
+
+	for _, tt := range []parseArgsTest{
+		{
+			name: "all_flags",
+			args: []string{"-n=10", "-c=5", "-rps=5", "http://test"},
+			want: config{n: 10, c: 5, rps: 5, url: "http://test"},
+		},
+
+		// exercise: test with a mixture of flags
+	} {
+		t.Run(tt.name, func(t *testing.T) {
+			t.Parallel()
+
+			var got config
+			if err := parseArgs(&got, tt.args, io.Discard); err != nil {
+				t.Fatalf("parseArgs() error = %v, want no error", err)
+			}
+			if got != tt.want {
+				t.Errorf("flags = %+v, want %+v", got, tt.want)
+			}
+		})
+	}
+}
+
+func TestParseArgsInvalidInput(t *testing.T) {
+	t.Parallel()
+
+	for _, tt := range []parseArgsTest{
+		{name: "n_syntax", args: []string{"-n=ONE", "http://test"}},
+		{name: "n_zero", args: []string{"-n=0", "http://test"}},
+		{name: "n_negative", args: []string{"-n=-1", "http://test"}},
+
+		// exercise: test other error conditions
+	} {
+		t.Run(tt.name, func(t *testing.T) {
+			t.Parallel()
+
+			err := parseArgs(&config{}, tt.args, io.Discard)
+			if err == nil {
+				t.Fatal("parseArgs() = nil, want error")
+			}
+		})
+	}
+}
```

