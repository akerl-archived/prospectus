package runner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/akerl/prospectus/v3/plugin"
	"github.com/ghodss/yaml"
)

// Value is the string result of the plugin, or error
type Value struct {
	Output string
	Error  error
}

// Result defines an individual state comparison
type Result struct {
	Name     string
	Expected Value
	Actual   Value
}

// ResultSet defines a group of state comparisons
type ResultSet []Result

// Plugin defines a partial struct that captures plugin type
type Plugin struct {
	Type string          `json:"type"`
	Args json.RawMessage `json:"args"`
	Meta plugin.Meta     `json:"-"`
}

// Check defines a set of expected/actual plugins
type Check struct {
	Type           string          `json:"type"`
	Args           json.RawMessage `json:"args"`
	ExpectedPlugin Plugin          `json:"expected"`
	ActualPlugin   Plugin          `json:"actual"`
}

// Expected returns either the merged plugin or expected plugin
func (c *Check) Expected() Plugin {
	meta := plugin.Meta{
		Mode: plugin.Expected,
	}
	if c.Type == "" {
		c.ExpectedPlugin.Meta = meta
		return c.ExpectedPlugin
	}
	return Plugin{
		Type: c.Type,
		Args: c.Args,
		Meta: meta,
	}
}

// Actual returns either the merged plugin or actual plugin
func (c *Check) Actual() Plugin {
	meta := plugin.Meta{
		Mode: plugin.Actual,
	}
	if c.Type == "" {
		c.ActualPlugin.Meta = meta
		return c.ActualPlugin
	}
	return Plugin{
		Type: c.Type,
		Args: c.Args,
		Meta: meta,
	}
}

// Runner defines a set of named Checks
type Runner struct {
	Items map[string]Check `json:"items"`
}

// New creates a new runner based on the configuration in the current directory
func New() (Runner, error) {
	r := Runner{}

	file, err := ioutil.ReadFile(".prospectus")
	if err != nil {
		return r, err
	}

	err = yaml.Unmarshal(file, &r)
	return r, err
}

// Check runs the plugins for each item to get its status
func (r Runner) Check() ResultSet {
	rs := ResultSet{}

	for name, check := range r.Items {
		res := Result{Name: name}
		res.Expected = check.Expected().Run()
		res.Actual = check.Actual().Run()
		rs = append(rs, res)
	}

	return rs
}

// Run executes an individual plugin
func (p Plugin) Run() Value {
	if p.Type == "" {
		return Value{Error: fmt.Errorf("plugin not set")}
	}

	pluginCommand := fmt.Sprintf("prospectus-%s", p.Type)
	cmd := exec.Command(pluginCommand)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return Value{Error: err}
	}

	pi := plugin.Input{
		Meta: p.Meta,
		Args: p.Args,
	}
	input, err := json.Marshal(pi)
	stdin.Write(input)
	stdin.Close()

	var stdoutBytes bytes.Buffer
	var stderrBytes bytes.Buffer

	cmd.Stdout = &stdoutBytes
	cmd.Stderr = &stderrBytes
	err = cmd.Run()
	if err != nil {
		return Value{Error: fmt.Errorf("%s: %s", err, strings.TrimSpace(stderrBytes.String()))}
	}

	return Value{Output: strings.TrimSpace(stdoutBytes.String())}
}

// Matches checks if a result ran successfully and if its expected and actual values match
func (r Result) Matches() bool {
	if r.Expected.Error != nil || r.Actual.Error != nil || r.Expected.Output != r.Actual.Output {
		return false
	}
	return true
}

// String converts the result to a human readable format
func (r Result) String() string {
	var expected, actual string
	if r.Expected.Error != nil {
		expected = r.Expected.Error.Error()
	} else {
		expected = r.Expected.Output
	}
	if r.Actual.Error != nil {
		actual = r.Actual.Error.Error()
	} else {
		actual = r.Actual.Output
	}
	return fmt.Sprintf("%s: %s / %s", r.Name, actual, expected)
}

// String converts the result set to a human readable format
func (rs ResultSet) String() string {
	var b strings.Builder
	for _, item := range rs {
		b.WriteString(item.String())
		b.WriteString("\n")
	}
	return b.String()
}
