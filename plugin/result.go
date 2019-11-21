package plugin

import (
	"fmt"
	"strings"

	"github.com/akerl/prospectus/expectations"
)

// Result defines the results of executing a Attribute
type Result struct {
	Actual    string               `json:"actual"`
	Expected  expectations.Wrapper `json:"expected"`
	Attribute Attribute            `json:"check"`
}

// ResultSet defines a group of Results
type ResultSet []Result

// String returns the Result as a human-readable string
func (r Result) String() string {
	return fmt.Sprintf(
		"%s: %s / %s",
		r.Attribute,
		r.Actual,
		r.Expected,
	)
}

// String returns the ResultSet as a human-readable string
func (rs ResultSet) String() string {
	var b strings.Builder
	for _, item := range rs {
		b.WriteString(item.String())
		b.WriteString("\n")
	}
	return b.String()
}

// Changed filters a ResultSet to only Results which do not match
func (rs ResultSet) Changed() ResultSet {
	var newResultSet ResultSet
	for _, item := range rs {
		if !item.Matches() {
			newResultSet = append(newResultSet, item)
		}
	}
	return newResultSet
}

// Matches returns true if the Expected and Actual values of the Result match
func (r Result) Matches() bool {
	return r.Expected.Matches(r.Actual)
}

// Fix attempts to resolve a mismatched expectation
func (r Result) Fix() Result {
	newResult := Result{}
	err := call(r.Attribute.File, "fix", r, &newResult)
	if err != nil {
		newResult = NewErrorResult(fmt.Sprintf("fix error: %s", err), r.Attribute)
	}
	newResult.Attribute = r.Attribute
	return newResult
}

// Fix attempts to fix all results in a ResultSet
func (rs ResultSet) Fix() ResultSet {
	newResultSet := make(ResultSet, len(rs))
	for index, item := range rs {
		if item.Matches() {
			newResultSet[index] = item
		} else {
			newResultSet[index] = item.Fix()
		}
	}
	return newResultSet
}

// NewErrorResult creates an error result from a given string
func NewErrorResult(msg string, c Attribute) Result {
	return Result{
		Actual: "error",
		Expected: expectations.Wrapper{
			Type: "error",
			Data: map[string]string{"msg": msg},
		},
		Attribute: c,
	}
}
