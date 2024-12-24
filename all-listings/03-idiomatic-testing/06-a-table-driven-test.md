# Listing 3.6: A table-driven test

## [testing](https://github.com/inancgumus/gobyexample/blob/b2456706240cea8014a4edf24fddc8292d12d94f/testing) / [url](https://github.com/inancgumus/gobyexample/blob/b2456706240cea8014a4edf24fddc8292d12d94f/testing/url) / [url_test.go](https://github.com/inancgumus/gobyexample/blob/b2456706240cea8014a4edf24fddc8292d12d94f/testing/url/url_test.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package url

import "testing"

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
	u := &URL{
		Scheme: "https",
		Host:   "go.dev",
		Path:   "play",
	}

	got := u.String()
	want := "https://go.dev/play"
	if got != want {
		t.Errorf("String() = %q, want %q", got, want)
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
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -46,3 +46,21 @@ func TestParseWithoutPath(t *testing.T) {
 		t.Errorf("Parse(%q)\ngot  %#v\nwant %#v", uri, got, want)
 	}
 }
+
+var parseTests = []struct {
+	name string
+	uri  string
+	want *URL
+}{
+	{
+		name: "full",
+		uri:  "https://go.dev/play",
+		want: &URL{Scheme: "https", Host: "go.dev", Path: "play"},
+	},
+	{
+		name: "without_path",
+		uri:  "http://github.com",
+		want: &URL{Scheme: "http", Host: "github.com", Path: ""},
+	},
+	/* many more test cases can be easily added */
+}
```

