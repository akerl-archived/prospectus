package plugin

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

// ParseConfig loads the config from stdin
func ParseConfig(cfg interface{}) error {
	info, err := os.Stdin.Stat()
	if err != nil {
		return err
	}

	if info.Mode()&os.ModeNamedPipe != os.ModeNamedPipe || info.Size() <= 0 {
		return fmt.Errorf("plugin executed without stdin")
	}

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(input, &cfg)
}
