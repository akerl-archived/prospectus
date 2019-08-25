package checks

import (
	"encoding/json"
	"strings"

	"github.com/akerl/go-prospectus/expectations"
)

// TODO: add timber logging
// TODO: add parallelization
// TODO: implement String Expected
// TODO: implement Regex Expected
// TODO: implement Set Expected

// Check defines a single check that is ready for execution
type Check struct {
	Dir      string
	File     string
	Name     string
	Metadata map[string]string
}

// CheckSet defines a group of Checks
type CheckSet []Check

// Result defines the results of executing a Check
type Result struct {
	Actual   string
	Expected Expected
	Check    Check
}

// ResultSet defines a group of Results
type ResultSet []Result

// Expected defines a pluggable interface for matching desired state to actual
type Expected interface {
	Matches(string) bool
	String() string
}

// NewSet returns a CheckSet based on a provided list of directories
func NewSet(dirs []string) (CheckSet, error) {
	// TODO: Actually load checks from directories
	return CheckSet{}, nil
}

// Execute returns the Results from a CheckSet by calling Execute on each Check
func (cs CheckSet) Execute() ResultSet {
	resultSet = make(ResultSet, len(cs))
	for index, item := range cs {
		resultSet[index] = item.Execute()
	}
	return resultSet
}

// Execute runs the Check and returns Results
func (c Check) Execute() Result {
	// TODO: actually run the check
	return Result{}
}

// Changed filters a ResultSet to only Results which do not match
func (rs ResultSet) Changed() ResultSet {
	var newResultSet ResultSet
	for _, item := range rs {
		if !i.Matches() {
			newResultSet = append(newResultSet, item)
		}
	}
}

// Matches returns true if the Expected and Actual values of the Result match
func (r Result) Matches() bool {
	return r.Expected.Matches(r.Actual)
}

// Json returns the ResultsSet as a marshalled JSON string
func (rs ResultSet) Json() (string, err) {
	data, err := json.MarshalIndent(rs, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data)
}

// String returns the ResultsSet as a human-readable string
func (rs ResultSet) String() string {
	var b strings.Builder
	for _, item := range rs {
		b.WriteString(item.String())
		b.WriteString("\n")
	}
	output = b.String()
}

// String returns the Result as a human-readable string
func (r Result) String() string {
	return fmt.Sprintf(
		"%s::%s: %s / %s",
		r.Check.Dir,
		r.Check.Name,
		r.Actual,
		r.Expected.String(),
	)
}
