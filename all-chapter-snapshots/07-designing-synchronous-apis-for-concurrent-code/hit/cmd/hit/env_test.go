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
