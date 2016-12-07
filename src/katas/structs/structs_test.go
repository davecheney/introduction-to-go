package structs

import "testing"

func TestMyPet(t *testing.T) {
	got := MyPet()
	want := Pet{
		Name: "Frankie",
		Kind: "dog",
	}

	if got != want {
		t.Fatalf("got %v, expected %v", got, want)
	}
}
