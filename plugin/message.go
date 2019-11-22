package plugin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

const (
	apiVersion = 1
)

type message struct {
	Version  int             `json:"version"`
	Contents json.RawMessage `json:"contents"`
}

func writeMessage(input interface{}) ([]byte, error) {
	contents, err := json.Marshal(input)
	if err != nil {
		return []byte{}, err
	}
	m := message{
		Version:  apiVersion,
		Contents: contents,
	}
	return json.Marshal(m)
}

func readMessage(input []byte, output interface{}) error {
	var m message
	err := json.Unmarshal(input, &m)
	if err != nil {
		return err
	}
	if m.Version != apiVersion {
		return fmt.Errorf(
			"plugin version mismatch: %d (expected) vs %d (actual)",
			apiVersion,
			m.Version,
		)
	}
	return json.Unmarshal(m.Contents, output)
}

func call(file, command string, input interface{}, output interface{}) error {
	cmd := exec.Command(file, command)

	inputBytes, err := writeMessage(input)
	if err != nil {
		return err
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	stdin.Write(inputBytes)
	stdin.Close()

	var stdoutBytes bytes.Buffer
	var stderrBytes bytes.Buffer

	cmd.Stdout = &stdoutBytes
	cmd.Stderr = &stderrBytes
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("%s: %s", err, stderrBytes.String())
	}

	return readMessage(stdoutBytes.Bytes(), output)
}
