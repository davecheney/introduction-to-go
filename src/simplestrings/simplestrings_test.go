package simplestrings

import "testing"

const weekdays = "Monday Tuesday Wednesday Thursday Friday"

func TestIsTuesdayAWeekeday(t *testing.T) {
	i := Index(weekdays, "Tuesday")
	if i == -1 {
		t.Fatal("Tuesday not found")
	}
}

// test that Tuesday is a weekday

// test that Sunday is not a weekday

// test that an empty search string returns 0

// test that the string Monday is not found in the empty string
