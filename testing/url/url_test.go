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
