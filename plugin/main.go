package plugin

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/akerl/prospectus/checks"

	"github.com/ghodss/yaml"
)

// Plugin defines a Golang plugin object for prospectus request handling
type Plugin interface {
	GetConfigPointer() interface{}
	Load(checks.LoadInput) checks.CheckSet
	Execute(checks.Check) checks.Result
	Fix(checks.Result) checks.Result
}

// Execute runs a plugin
func Execute(p Plugin) error {
	if len(os.Args) != 3 {
		return fmt.Errorf("Unexpected number of args provided: %d", len(os.Args))
	}

	configFile := os.Args[1]
	subcommand := os.Args[2]

	c := p.GetConfigPointer()
	err := loadPluginConfig(c)
	if err != nil {
		return err
	}

	switch subcommand {
	case "load":
	case "execute":
	case "fix":
	default:
		return fmt.Errorf("Unexpected command provided: %s", subcommand)
	}
}

func loadPluginConfig(output interface{}) error {
	if len(os.Args) < 2 {
		return fmt.Errorf("no config file path provided")
	}
	file := os.Args[1]

	fileInfo, err := os.Stat(file)
	if os.IsNotExist(err) {
		return fmt.Errorf("config file does not exist")
	}
	if fileInfo.IsDir() {
		return fmt.Errorf("config file is a directory")
	}

	data, err := ioutil.ReadFile(file)
	return yaml.Unmarshal(data, output)
}
