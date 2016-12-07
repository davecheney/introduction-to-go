package main

import (
	"strings"
	"testing"
)

func TestReadAll(t *testing.T) {
	const lines = `Line 1
Line 2
Line 3
Line 4
Line 5`

	r := strings.NewReader(lines)
	got := ReadAll(r)
	want := lines
	if got != want {
		t.Fatalf("got %q, expected %q", got, want)
	}
}
