package expectations

import (
	"strings"
)

type setExpectation struct {
	expected []string
	raw      string
}

func newSetExpectation(data map[string]string) Expectation {
	separator := data["separator"]
	if separator == "" {
		separator = ","
	}
	raw := data["expected"]
	parts = strings.Split(raw, separator)
	return setExpectation{expected: parts, raw: raw}
}

// Matches returns true if the actual value exists in the expected set
func (s setExpectation) Matches(actual string) bool {
	for _, item := range s.expected {
		if item == actual {
			return true
		}
	}
	return false
}

// String returns the original string with separators intact
func (s setExpectation) String() string {
	return s.raw
}
