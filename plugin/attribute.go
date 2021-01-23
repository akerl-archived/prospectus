package plugin

import (
	"fmt"
	"strings"
)

// Attribute defines a single check that is ready for execution
type Attribute struct {
	Dir      string            `json:"dir"`
	File     string            `json:"file"`
	Name     string            `json:"name"`
	Metadata map[string]string `json:"metadata"`
}

// AttributeSet defines a group of Attributes
type AttributeSet []Attribute

// String returns the Result as a human-readable string
func (a Attribute) String() string {
	return fmt.Sprintf(
		"%s::%s",
		a.Dir,
		a.Name,
	)
}

// String returns the AttributeSet as a human-readable string
func (as AttributeSet) String() string {
	var b strings.Builder
	for _, item := range as {
		b.WriteString(item.String())
		b.WriteString("\n")
	}
	return b.String()
}

// Check runs the Attribute and returns Results
func (a Attribute) Check() Result {
	r := Result{}
	err := call(a.File, "check", a, &r)
	if err != nil {
		r = NewErrorResult(fmt.Sprintf("execution error: %s", err))
	}
	r.Attribute = a
	return r
}

// Check returns the Results from a AttributeSet by calling Execute on each Attribute
func (as AttributeSet) Check() ResultSet {
	resultSet := make(ResultSet, len(as))
	for index, item := range as {
		resultSet[index] = item.Check()
	}
	return resultSet
}
