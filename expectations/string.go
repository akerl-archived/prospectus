package expectations

type stringExpectation struct {
	expected string
}

// Load creates a new string expectation
func (s *stringExpectation) Load(data map[string]string) Expectation {
	s.expected = data["expected"]
	return s
}

// Matches returns true if the expected and actual strings are identical
func (s *stringExpectation) Matches(actual string) bool {
	return s.expected == actual
}

// String returns the expected string
func (s *stringExpectation) String() string {
	return s.expected
}
