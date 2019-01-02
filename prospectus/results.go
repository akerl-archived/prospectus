package prospectus

// Results is a set of Result objects
type Results []Result

// Result holds the output of a single state check
type Result struct {
	Path     string
	Name     string
	Actual   string
	Expected Expected
}

// Expected defines the interface for expected states, which must be compared
// against Actual state strings
type Expected interface {
	Match(string) bool
}
