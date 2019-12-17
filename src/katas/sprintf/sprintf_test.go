package sprintf

import (
	"fmt"
	"strings"
	"testing"
)

func TestTrainJourney(t *testing.T) {
	train, number, duration := "新幹線", 291, 2.2
	got := fmt.Sprintf("%v number %v duration %v hours", train, number, duration)
	want := "新幹線 number 291 duration 2.2 hours"
	if got != want {
		t.Fatalf("got %q, expected %q", got, want)
	}
}

func TestCities(t *testing.T) {
	cities := []string{"東京", "京都", "大阪市"}

	got := fmt.Sprintf("%s", strings.Join(cities, ", "))
	want := "東京, 京都, 大阪市"
	if got != want {
		t.Fatalf("got %q, expected %q", got, want)
	}
}
