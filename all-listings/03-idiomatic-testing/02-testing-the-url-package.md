# Listing 3.2: Testing the url package

## [testing](https://github.com/inancgumus/gobyexample/blob/8a4cda260d4f85a7de2002600589607fb4378c53/testing) / [url](https://github.com/inancgumus/gobyexample/blob/8a4cda260d4f85a7de2002600589607fb4378c53/testing/url) / [url_test.go](https://github.com/inancgumus/gobyexample/blob/8a4cda260d4f85a7de2002600589607fb4378c53/testing/url/url_test.go)

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
```

