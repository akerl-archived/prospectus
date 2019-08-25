package expectations

import (
	"fmt"
)

type errorExpectation struct {
	msg string
}

func newErrorExpectation(data map[string]string) Expectation {
	return errorExpectation{msg: data["msg"]}
}

// Matches returns false in all cases
func (e errorExpectation) Matches(_ string) bool {
	return false
}

// String returns a type error message
func (e errorExpectation) String() string {
	return fmt.Sprintf("error: %s", e.msg)
}
