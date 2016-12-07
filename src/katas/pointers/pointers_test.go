package pointers

import (
	"fmt"
	"testing"
)

func TestPointString(t *testing.T) {
	p := Point{X: 300, Y: 60}
	p.Move(-200, 40)
	got := fmt.Sprintf("%v", p)
	want := "point: x=100, y=100"
	if got != want {
		t.Fatalf("got %q, expected %q", got, want)
	}
}
