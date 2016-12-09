// +build ignore

package jsonenc

type Person struct {
	Firstname  string `json:"first"`
	Middlename string `json:"middle,omitempty"`
	Lastname   string `json:"last"`

	SSID int64 `json:"-"`

	City    string `json:"city,omitempty"`
	Country string `json:"country"`

	Telephone int64 `json:"tel,string"`
}
