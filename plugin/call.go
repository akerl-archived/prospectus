package plugin

import (
	"encoding/json"
	"os/exec"
)

func call(file, command string, input interface{}, output interface{}) error {
	cmd := exec.Command(file, command)

	inputBytes, err := WriteMessage(input)
	if err != nil {
		return err
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	stdin.Write(inputBytes)
	stdin.Close()

	stdout, err := cmd.Output()
	if err != nil {
		return err
	}

	return ReadMessage(stdout, output)
}
