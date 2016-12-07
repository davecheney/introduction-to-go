package methods

import (
	"fmt"
	"testing"
)

func TestPointString(t *testing.T) {
	p := Point{X: 300, Y: 60}
	got := fmt.Sprintf("%v", p)
	want := "point: x=300, y=60"
	if got != want {
		t.Fatalf("got %q, expected %q", got, want)
	}
}
