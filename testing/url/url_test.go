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
		t.Fatalf("Parse(%q) err = %q, want nil", uri, err)
	}
	want := &URL{
		Scheme: "https", Host: "go.dev", Path: "play",
	}
	if *got != *want {
		t.Errorf("Parse(%q)\ngot  %#v\nwant %#v", uri, got, want)
	}
}

func TestParseWithoutPath(t *testing.T) {
	const uri = "https://github.com"

	got, err := Parse(uri)
	if err != nil {
		t.Fatalf("Parse(%q) err = %q, want nil", uri, err)
	}
	want := &URL{
		Scheme: "https", Host: "github.com", Path: "",
	}
	if *got != *want {
		t.Errorf("Parse(%q)\ngot  %#v\nwant %#v", uri, got, want)
	}
}

var parseTests = []struct { //nolint:gochecknoglobals
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
			t.Fatalf("Parse(%q) err = %v, want nil", tt.uri, err)
		}
		if *got != *tt.want {
			t.Errorf("Parse(%q)\ngot  %#v\nwant %#v", tt.uri, got, tt.want)
		}
	}
}

func TestParseSubtests(t *testing.T) {
	// put setup or teardown code if necessary here
	for _, tt := range parseTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.uri)
			if err != nil {
				t.Fatalf("Parse(%q) err = %v, want nil", tt.uri, err)
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
		/* more test cases should be added */
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := Parse(tt.uri); err == nil {
				t.Errorf("Parse(%q)=nil; want an error", tt.uri)
			}
		})
	}
}

func TestURLString(t *testing.T) {
	t.SkipNow()
	tests := []struct {
		name string
		uri  *URL
		want string
	}{
		{name: "nil", uri: nil, want: ""},
		{name: "empty", uri: &URL{}, want: ""},
		/* more test cases should be added */
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
		b.ResetTimer()
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for range b.N {
				_ = u.String()
			}
		})
	}
}

func Pow(a, n int) int {
	for range n {
		a *= n
	}
	return a
}

func BenchmarkPow1(b *testing.B) {
	for range b.N {
		Pow(5, 1)
	}
}

var Sink int

func BenchmarkPow2(b *testing.B) {
	var sink int
	for range b.N {
		sink = Pow(5, 1)
	}
	Sink = sink
}
