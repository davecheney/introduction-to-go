package simplestrings

import "testing"

const weekdays = "Monday Tuesday Wednesday Thursday Friday"

func TestIsTuesdayAWeekeday(t *testing.T) {
	i := Index(weekdays, "Tuesday")
	if i == -1 {
		t.Fatal("Tuesday not found")
	}
}

func TestSundayIsNotAWeekday(t *testing.T) {
	i := Index(weekdays, "Sunday")
	if i != -1 {
		t.Fatal("Sunday found unexpectedly")
	}
}

func TestEmptyStringReturnsZero(t *testing.T) {
	i := Index(weekdays, "")
	if i != 0 {
		t.Fatal("expected i == 0")
	}
}

func TestMondayNotFoundInEmptyString(t *testing.T) {
	i := Index("", "Monday")
	if i != -1 {
		t.Fatal("found Monday in empty string")
	}
}
