package plugin

import (
	"fmt"
	"strings"
)

// Result defines the results of executing a Attribute
type Result struct {
	Actual    string    `json:"actual"`
	Expected  string    `json:"expected"`
	Matches   bool      `json:"matches"`
	Attribute Attribute `json:"attribute"`
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

// Fix attempts to resolve a mismatched expectation
func (r Result) Fix() Result {
	if r.Matches {
		return r
	}
	newResult := Result{}
	err := call(r.Attribute.File, "fix", r, &newResult)
	if err != nil {
		newResult = NewErrorResult(fmt.Sprintf("fix error: %s", err))
	}
	newResult.Attribute = r.Attribute
	return newResult
}

// Fix attempts to fix all results in a ResultSet
func (rs ResultSet) Fix() ResultSet {
	newResultSet := make(ResultSet, len(rs))
	for index, item := range rs {
		newResultSet[index] = item.Fix()
	}
	return newResultSet
}

// NewErrorResult creates an error result from a given string
func NewErrorResult(msg string) Result {
	return Result{
		Actual:   "error",
		Expected: msg,
		Matches:  false,
	}
}
