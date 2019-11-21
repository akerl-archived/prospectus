package plugin

// Check defines a single check that is ready for execution
type Check struct {
	Dir      string            `json:"dir"`
	File     string            `json:"file"`
	Name     string            `json:"name"`
	Metadata map[string]string `json:"metadata"`
}

// CheckSet defines a group of Checks
type CheckSet []Check

// String returns the Result as a human-readable string
func (c Check) String() string {
	return fmt.Sprintf(
		"%s::%s",
		c.Dir,
		c.Name,
	)
}

// String returns the CheckSet as a human-readable string
func (cs CheckSet) String() string {
	var b strings.Builder
	for _, item := range cs {
		b.WriteString(item.String())
		b.WriteString("\n")
	}
	return b.String()
}

// Execute runs the Check and returns Results
func (c Check) Execute() Result {
	r := Result{}
	err := call(c.File, "execute", c, &r)
	if err != nil {
		r = NewErrorResult(fmt.Sprintf("%s error: %s", method, err), c)
	}
	r.Check = c
	return r
}

// Execute returns the Results from a CheckSet by calling Execute on each Check
func (cs CheckSet) Execute() ResultSet {
	resultSet := make(ResultSet, len(cs))
	for index, item := range cs {
		resultSet[index] = item.Execute()
	}
	return resultSet
}
