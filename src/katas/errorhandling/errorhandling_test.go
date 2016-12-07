package errorhandling

import "testing"

func TestCountLines(t *testing.T) {
	got, err := CountLines("testdata/alice.txt")
	if err != nil {
		t.Fatalf("got error %v, expected nil", err)
	}
	want := 3736
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestMissingFile(t *testing.T) {
	got, err := CountLines("testdata/missing")
	want := 0
	if err == nil {
		t.Fatalf("got error nil, expected non nil")
	}
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}
