package simplestrings

import "testing"

const weekdays = "Monday Tuesday Wednesday Thursday Friday"

func TestIndexFound(t *testing.T) {
	i := Index(weekdays, "Tuesday")
	if i == -1 {
		t.Fatal("Tuesday not found")
	}
}

func TestIndexNotFound(t *testing.T) {
	i := Index(weekdays, "Sunday")
	if i != -1 {
		t.Fatal("Sunday found, not expected")
	}
}

func TestIndexEmptySubString(t *testing.T) {
	i := Index(weekdays, "")
	if i != 0 {
		t.Fatalf(`Index(weekdays, ""): got: %d, want: 0`, i)
	}
}

func TestIndexEmptyString(t *testing.T) {
	i := Index("", "Monday")
	if i != -1 {
		t.Fatalf(`Index("", "Monday"): got: %d, want: -1`, i)
	}
}
