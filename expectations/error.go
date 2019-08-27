package expectations

import (
	"fmt"
)

type errorExpectation struct {
	msg string
}

// Load creates a new error expectation
func (e *errorExpectation) Load(data map[string]string) Expectation {
	e.msg = data["msg"]
	return e
}

// Matches returns false in all cases
func (e *errorExpectation) Matches(_ string) bool {
	return false
}

// String returns a type error message
func (e *errorExpectation) String() string {
	return fmt.Sprintf("error: %s", e.msg)
}
