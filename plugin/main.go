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
	err := preflightChecks()
	if err != nil {
		return err
	}

	configFile := os.Args[1]
	subcommand := os.Args[2]

	err = loadPluginConfig(configFile, p)
	if err != nil {
		return err
	}

	inputMsg, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	output, err := runSubcommand(subcommand, inputMsg, p)
	if err != nil {
		return err
	}

	outputMsg, err := writeMessage(output)
	if err != nil {
		return err
	}

	fmt.Print(string(outputMsg))
	return nil
}

func runSubcommand(subcommand string, inputMsg []byte, p Plugin) (interface{}, error) {
	var output interface{}

	switch subcommand {
	case "load":
		input := LoadInput{}
		if err := readMessage(inputMsg, &input); err != nil {
			return output, err
		}
		output = p.Load(input)
	case "check":
		input := Attribute{}
		if err := readMessage(inputMsg, &input); err != nil {
			return output, err
		}
		output = p.Check(input)
	case "fix":
		input := Result{}
		if err := readMessage(inputMsg, &input); err != nil {
			return output, err
		}
		output = p.Fix(input)
	default:
		return output, fmt.Errorf("unexpected command provided: %s", subcommand)
	}

	return output, nil
}

func preflightChecks() error {
	if len(os.Args) != 3 {
		return fmt.Errorf("unexpected number of args provided: %d", len(os.Args))
	}

	info, err := os.Stdin.Stat()
	if err != nil {
		return err
	}

	if info.Mode()&os.ModeNamedPipe != os.ModeNamedPipe || info.Size() <= 0 {
		return fmt.Errorf("plugin executed without stdin")
	}

	return nil
}

func loadPluginConfig(configFile string, p Plugin) error {
	output := p.GetConfigPointer()
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
