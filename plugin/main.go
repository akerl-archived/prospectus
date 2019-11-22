package plugin

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/akerl/timber/v2/log"
	"github.com/ghodss/yaml"
)

// TODO: add parallelization
// TODO: add logging

var mainLogger = log.NewLogger("prospectus")
var pluginLogger = log.NewLogger("prospectus:plugin")

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

	var output interface{}

	switch subcommand {
	case "load":
		input := LoadInput{}
		if err := ReadMessage(inputMsg, &input); err != nil {
			return err
		}
		output = p.Load(input)
	case "check":
		input := Attribute{}
		if err := ReadMessage(inputMsg, &input); err != nil {
			return err
		}
		output = p.Check(input)
	case "fix":
		input := Result{}
		if err := ReadMessage(inputMsg, &input); err != nil {
			return err
		}
		output = p.Fix(input)
	default:
		return fmt.Errorf("Unexpected command provided: %s", subcommand)
	}

	outputMsg, err := WriteMessage(output)
	if err != nil {
		return err
	}
	fmt.Print(string(outputMsg))
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
