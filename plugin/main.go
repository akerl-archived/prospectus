package plugin

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/akerl/prospectus/checks"

	"github.com/ghodss/yaml"
)

type Check checks.Check
type Result checks.Result

// LoadPluginConfig reads a plugin config from the provided source file
func LoadPluginConfig(output interface{}) error {
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
