package expectations

import (
	"fmt"
)

type errorExpectation struct {
	typeName string
}

func newErrorExpectation(data map[string]string) Expectation {
	return errorExpectation{typeName: data["type"]}
}

// Matches returns false in all cases
func (e errorExpectation) Matches(_ string) bool {
	return false
}

// String returns a type error message
func (e errorExpectation) String() string {
	return fmt.Sprintf("error: expectation type not known: %s", e.typeName)
}
