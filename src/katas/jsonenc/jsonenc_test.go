package jsonenc

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func encode(t *testing.T, p *Person) string {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	err := enc.Encode(p)
	if err != nil {
		t.Fatal(err)
	}
	return strings.TrimSpace(b.String()) // remove newlines
}

func TestPerson(t *testing.T) {
	p := &Person{
		Firstname: "Bill",
		Lastname:  "Gates",
		SSID:      235897234582,
		City:      "Seattle",
		Country:   "USA",
		Telephone: 18556149487,
	}

	got := encode(t, p)
	want := `{"first":"Bill","last":"Gates","city":"Seattle","country":"USA","tel":"18556149487"}`
	if got != want {
		t.Fatalf("encode(%v): got %s, want %s", p, got, want)
	}
}
