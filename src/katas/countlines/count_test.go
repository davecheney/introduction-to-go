package count

import "testing"

func TestCountLines(t *testing.T) {
	got := CountLines("testdata/alice.txt")
	want := 3736
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}
