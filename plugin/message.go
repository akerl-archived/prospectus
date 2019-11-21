package plugin

import (
	"encoding/json"
	"fmt"
)

const (
	apiVersion = 1
)

type message struct {
	Version  int             `json:"version"`
	Contents json.RawMessage `json:"contents"`
}

func WriteMessage(input interface{}) ([]byte, error) {
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

func ReadMessage(input []byte, output interface{}) error {
	var m message
	err := json.Unmarshal(input, &m)
	if err != nil {
		return err
	}
	if m.Version != apiVersion {
		return fmt.Errorf("plugin version mismatch: %s (expected) vs %s (actual)", apiVersion, m.Version)
	}
	return json.Unmarshal(m.Contents, output)
}
