package plugin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

// Mode signals whether the check should return expected or actual value
type Mode int

const (
	// Actual checks the current state of the check
	Actual Mode = iota
	// Expected checks the desired state of the check
	Expected
)

// Meta passes additional information to the check from the runner
type Meta struct {
	Mode Mode `json:"mode"`
}

// Input merges runner metadata with config args
type Input struct {
	Meta Meta            `json:"meta"`
	Args json.RawMessage `json:"args"`
}

// ParseConfig loads the config from stdin
func ParseConfig(cfg interface{}) (Meta, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return Meta{}, err
	}

	if info.Mode()&os.ModeNamedPipe != os.ModeNamedPipe || info.Size() <= 0 {
		return Meta{}, fmt.Errorf("plugin executed without stdin")
	}

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return Meta{}, err
	}

	pi := Input{}
	err = yaml.Unmarshal(input, &pi)
	if err != nil {
		return Meta{}, err
	}

	return pi.Meta, yaml.Unmarshal(pi.Args, &cfg)
}
