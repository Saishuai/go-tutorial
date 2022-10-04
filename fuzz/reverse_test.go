package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testCases {
		rev := Reverse(tc.in)
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}

func FuzzReverse(f *testing.F) {
	testCases := []string{"Hello world", " ", "!12345"}
	for _, tc := range testCases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, source string) {
		rev, err := ReverseFix(source)
		if err != nil {
			return
		}
		revRev, errRev := ReverseFix(rev)
		if errRev != nil {
			return
		}
		if source != revRev {
			t.Errorf("Before: %q, after: %q", source, revRev)
		}
		t.Logf("Number of runes: source=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(source), utf8.RuneCountInString(rev), utf8.RuneCountInString(revRev))
		if utf8.ValidString(source) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
