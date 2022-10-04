package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	input := "Turing Test"
	rev := Reverse(input)
	revRev := Reverse(rev)
	fmt.Printf("origin: %q\n", input)
	fmt.Printf("reverse: %q\n", rev)
	fmt.Printf("reverse again: %q\n", revRev)
}

func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		// t := b[i]
		// b[j] = b[i]
		// b[i] = t
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func ReverseFix(s string) (string, error) {
	fmt.Printf("\nString valid check: %q: %v\n", s, utf8.ValidString(s))
	if !utf8.ValidString(s) {
		return "", errors.New("input string is invalid utf-8")
	}
	b := []rune(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		// t := b[i]
		// b[j] = b[i]
		// b[i] = t
		b[i], b[j] = b[j], b[i]
	}
	return string(b), nil
}
