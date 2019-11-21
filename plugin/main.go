package plugin

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

// TODO: add timber logging
// TODO: add parallelization

// Plugin defines a Golang plugin object for prospectus request handling
type Plugin interface {
	GetConfigPointer() interface{}
	Load(LoadInput) AttributeSet
	Check(Attribute) Result
	Fix(Result) Result
}

// Start runs a plugin
func Start(p Plugin) error {
	if len(os.Args) != 3 {
		return fmt.Errorf("Unexpected number of args provided: %d", len(os.Args))
	}

	configFile := os.Args[1]
	subcommand := os.Args[2]

	c := p.GetConfigPointer()
	err := loadPluginConfig(configFile, c)
	if err != nil {
		return err
	}

	info, err := os.Stdin.Stat()
	if err != nil {
		return err
	}

	if info.Mode()&os.ModeNamedPipe != os.ModeNamedPipe || info.Size() <= 0 {
		return fmt.Errorf("Plugin executed without stdin")
	}

	inputMsg, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	var input interface{}
	var output interface{}

	switch subcommand {
	case "load":
		input = LoadInput{}
	case "check":
		input = Attribute{}
	case "fix":
		input = Result{}
	default:
		return fmt.Errorf("Unexpected command provided: %s", subcommand)
	}

	if err := ReadMessage(inputMsg, &input); err != nil {
		return err
	}

	switch subcommand {
	case "load":
		output = p.Load(input.(LoadInput))
	case "check":
		output = p.Check(input.(Attribute))
	case "fix":
		output = p.Fix(input.(Result))
	}

	outputMsg, err := WriteMessage(output)
	fmt.Print(outputMsg)
	return nil
}

func loadPluginConfig(configFile string, output interface{}) error {
	fileInfo, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		return fmt.Errorf("config file does not exist")
	}
	if fileInfo.IsDir() {
		return fmt.Errorf("config file is a directory")
	}

	data, err := ioutil.ReadFile(configFile)
	return yaml.Unmarshal(data, output)
}
