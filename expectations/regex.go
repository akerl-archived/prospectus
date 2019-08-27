package expectations

import (
	"fmt"
	"regexp"
)

type regexExpectation struct {
	regex *regexp.Regexp
	raw   string
}

// Load creates a new regex expectation
func (r *regexExpectation) Load(data map[string]string) Expectation {
	var err error
	r.raw = data["pattern"]
	r.regex, err = regexp.Compile(r.raw)
	if err != nil {
		e := errorExpectation{}
		return e.Load(map[string]string{
			"msg": fmt.Sprintf("invalid regex: %s (%s)", r.raw, err),
		})
	}
	return r
}

// Matches returns true if the actual value exists in the expected regex
func (r *regexExpectation) Matches(actual string) bool {
	return r.regex.MatchString(actual)
}

// String returns the original string with separators intact
func (r *regexExpectation) String() string {
	return r.raw
}
