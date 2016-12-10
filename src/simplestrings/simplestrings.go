package simplestrings

// Index returns the index of the first instance of sub
// in s, or -1 if sub is not present in s.
func Index(s, sub string) int {
	for i := range s {
		v := s[i:]
		if len(sub) > len(v) {
			break
		}
		if v[:len(sub)] == sub {
			return i
		}
	}
	return -1
}

// Contains reports whether substr is within s.
func Contains(s, substr string) bool {
	return Index(s, substr) >= 0
}

// HasPrefix tests whether the string s begins with prefix.
func HasPrefix(s, prefix string) bool {
	return Index(s, prefix) == 0
}
