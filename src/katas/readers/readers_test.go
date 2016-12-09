package readers

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestCombineReaders(t *testing.T) {
	a := strings.NewReader("first reader\n")
	b := strings.NewReader("second reader\n")
	r := Combine(a, b)

	var w bytes.Buffer
	io.Copy(&w, r)
	got := w.String()
	want := "first reader\nsecond reader\n"
	if got != want {
		t.Fatalf("combine: got %q, want %q", got, want)
	}
}

func TestAReader10(t *testing.T) {
	r := AReader(10)
	var w bytes.Buffer
	io.Copy(&w, r)
	got := w.String()
	want := "AAAAAAAAAA"
	if got != want {
		t.Fatalf("AReader(10): got %q, want %q", got, want)
	}
}

func TestAReader20(t *testing.T) {
	r := AReader(20)
	var w bytes.Buffer
	io.Copy(&w, r)
	got := w.String()
	want := "AAAAAAAAAAAAAAAAAAAA"
	if got != want {
		t.Fatalf("AReader(20): got %q, want %q", got, want)
	}
}
