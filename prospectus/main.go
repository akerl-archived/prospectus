package prospectus

import (
	"fmt"
)

// Results is a set of Result objects
type Results []Result

// Result holds the output of a single state check
type Result struct {
	Check    Check
	Actual   string
	Expected Expected
}

// Matches returns true if the Actual and Expected states match
func (r *Result) Matches() bool {
	return r.Expected.Matches(r.Actual)
}

// String prints the result as a human-readable string
func (r *Result) String() string {
	return fmt.Sprintf(
		"%s::%s: %s / %s",
		r.Check.Dir,
		r.Check.Name,
		r.Actual,
		r.Expected.String(),
	)
}

// Expected defines the interface for expected states, which must be compared
// against Actual state strings
type Expected interface {
	Matches(string) bool
	String() string
}

// Checks is a set of Check objects
type Checks []Check

// Check describes a parsed check that is ready for execution
type Check struct {
	Dir  string
	File string
	Name string
}

// Populate parses checks from a list of directories
func (cs *Checks) Populate(_ []string) error {
	// TODO: execute each dir .prospectus/.prospectus.d to get checks
	return nil
}

// Execute runs the set of checks
func (cs *Checks) Execute() (Results, error) {
	// TODO: execute checks
	// TODO: parallel execution
	return Results{}, nil
}
