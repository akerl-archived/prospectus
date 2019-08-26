package expectations

type stringExpectation struct {
	expected string
}

func newStringExpectation(data map[string]string) Expectation {
	return stringExpectation{expected: data["expected"]}
}

// Matches returns true if the expected and actual strings are identical
func (s stringExpectation) Matches(actual string) bool {
	return s.expected == actual
}

// String returns the expected string
func (s stringExpectation) String() string {
	return s.expected
}
