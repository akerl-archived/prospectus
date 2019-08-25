package expectations

import (
	"fmt"
	"regexp"
)

type regexExpectation struct {
	regex regexp.Regexp
	raw   string
}

func newRegexExpectation(data map[string]string) Expectation {
	raw := data["pattern"]
	pattern, err := regexp.Compile(raw)
	if err != nil {
		return newErrorExpectation(map[string]string{
			"msg": fmt.Sprintf("invalid regex: %s (%s)", raw, err),
		})
	}
	return regexExpectation{pattern: pattern, raw: raw}
}

// Matches returns true if the actual value exists in the expected regex
func (r regexExpectation) Matches(actual string) bool {
	return r.regex.MatchString(actual)
}

// String returns the original string with separators intact
func (r regexExpectation) String() string {
	return r.raw
}
