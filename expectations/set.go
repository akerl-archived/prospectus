package expectations

import (
	"strings"
)

type setExpectation struct {
	expected  []string
	separator string
	raw       string
}

// Load creates a new set expectation
func (s *setExpectation) Load(data map[string]string) Expectation {
	s.separator = data["separator"]
	if s.separator == "" {
		s.separator = ","
	}
	s.raw = data["expected"]
	s.expected = strings.Split(s.raw, s.separator)
	return s
}

// Matches returns true if the actual value exists in the expected set
func (s *setExpectation) Matches(actual string) bool {
	for _, item := range s.expected {
		if item == actual {
			return true
		}
	}
	return false
}

// String returns the original string with separators intact
func (s *setExpectation) String() string {
	return s.raw
}
