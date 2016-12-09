package jsonenc

type Person struct {
	Firstname  string // should be encoded as 'first'
	Middlename string // should be encoded as 'middle', and not present if blank
	Lastname   string // should be encoded as 'last'

	SSID int64 // should not be encoded

	City    string // should be encoded as 'city' and not present if missing
	Country string // should be encoded as 'country'

	Telephone int64 // should be encoded as 'tel', the value should be a string, not a number
}
