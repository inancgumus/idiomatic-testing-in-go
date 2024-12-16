# Listing 4.9: Writing sub-benchmarks

## [testing](https://github.com/inancgumus/gobyexample/blob/215d3786f7a7f3af04fd340192053cd7143b1714/testing) / [url](https://github.com/inancgumus/gobyexample/blob/215d3786f7a7f3af04fd340192053cd7143b1714/testing/url) / [url_test.go](https://github.com/inancgumus/gobyexample/blob/215d3786f7a7f3af04fd340192053cd7143b1714/testing/url/url_test.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package url

import (
	"fmt"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	const uri = "https://go.dev/play"

	got, err := Parse(uri)
	if err != nil {
		t.Fatalf("Parse(%q) err = %q, want <nil>", uri, err)
	}
	want := &URL{
		Scheme: "https", Host: "go.dev", Path: "play",
	}
	if *got != *want {
		t.Errorf("Parse(%q)\ngot  %#v\nwant %#v", uri, got, want)
	}
}

func TestURLString(t *testing.T) {
	tests := []struct {
		name string
		uri  *URL
		want string
	}{
		{name: "nil", uri: nil, want: ""},
		{name: "empty", uri: &URL{}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.uri.String()
			if got != tt.want {
				t.Errorf("\ngot  %q\nwant %q\nfor  %#v", got, tt.want, tt.uri)
			}
		})
	}
}

func TestParseWithoutPath(t *testing.T) {
	const uri = "https://github.com"

	got, err := Parse(uri)
	if err != nil {
		t.Fatalf("Parse(%q) err = %q, want <nil>", uri, err)
	}

	want := &URL{
		Scheme: "https", Host: "github.com", Path: "",
	}
	if *got != *want {
		t.Errorf("Parse(%q)\ngot  %#v\nwant %#v", uri, got, want)
	}
}

var parseTests = []struct {
	name string
	uri  string
	want *URL
}{
	{
		name: "with_data_scheme",
		uri:  "data:text/plain;base64,R28gYnkgRXhhbXBsZQ==",
		want: &URL{Scheme: "data"},
	},
	{
		name: "full",
		uri:  "https://go.dev/play",
		want: &URL{Scheme: "https", Host: "go.dev", Path: "play"},
	},
	{
		name: "without_path",
		uri:  "http://github.com",
		want: &URL{Scheme: "http", Host: "github.com", Path: ""},
	},
	/* many more test cases can be easily added */
}

func TestParseTable(t *testing.T) {
	for _, tt := range parseTests {
		t.Logf("run %s", tt.name)

		got, err := Parse(tt.uri)
		if err != nil {
			t.Fatalf("Parse(%q) err = %v, want <nil>", tt.uri, err)
		}
		if *got != *tt.want {
			t.Errorf("Parse(%q)\ngot  %#v\nwant %#v", tt.uri, got, tt.want)
		}
	}
}

func TestParseSubtests(t *testing.T) {
	// common test setup and teardown logic can be put here
	for _, tt := range parseTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.uri)
			if err != nil {
				t.Fatalf("Parse(%q) err = %v, want <nil>", tt.uri, err)
			}
			if *got != *tt.want {
				t.Errorf("Parse(%q)\ngot  %#v\nwant %#v", tt.uri, got, tt.want)
			}
		})
	}
}

func TestParseError(t *testing.T) {
	tests := []struct {
		name string
		uri  string
	}{
		{name: "without_scheme", uri: "go.dev"},
		{name: "empty_scheme", uri: "://go.dev"},

		/* we'll add more tests soon */
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Parse(tt.uri)
			if err == nil {
				t.Errorf("Parse(%q)=nil; want an error", tt.uri)
			}
		})
	}
}

func BenchmarkURLString(b *testing.B) {
	u := &URL{
		Scheme: "https",
		Host:   "go.dev",
		Path:   "play",
	}
	for range b.N {
		_ = u.String()
	}
}

func BenchmarkURLStringLong(b *testing.B) {
	for _, n := range []int{1, 10, 100, 1_000} {
		u := &URL{
			Scheme: strings.Repeat("x", n),
			Host:   strings.Repeat("y", n),
			Path:   strings.Repeat("z", n),
		}
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for range b.N {
				_ = u.String()
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
@@ -1,6 +1,10 @@
 package url
 
-import "testing"
+import (
+	"fmt"
+	"strings"
+	"testing"
+)
 
 func TestParse(t *testing.T) {
 	const uri = "https://go.dev/play"
@@ -134,3 +138,18 @@ func BenchmarkURLString(b *testing.B) {
 		_ = u.String()
 	}
 }
+
+func BenchmarkURLStringLong(b *testing.B) {
+	for _, n := range []int{1, 10, 100, 1_000} {
+		u := &URL{
+			Scheme: strings.Repeat("x", n),
+			Host:   strings.Repeat("y", n),
+			Path:   strings.Repeat("z", n),
+		}
+		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
+			for range b.N {
+				_ = u.String()
+			}
+		})
+	}
+}
```

