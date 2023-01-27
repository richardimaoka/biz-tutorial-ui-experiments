package model

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func (t *TerminalNode) UnmarshalJSON(b []byte) error {
	// unmarshal non-union fields first
	var partial struct {
		Index *int `json:"index"`
	}

	err := json.Unmarshal(b, &partial)
	if err != nil {
		return err
	}

	t.Index = partial.Index

	// then, unmarshal union fields
	var unmarshald map[string]interface{}
	err = json.Unmarshal(b, &unmarshald)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(unmarshald["content"])
	if err != nil {
		return err
	}

	content, err := terminalElementFromBytes(bytes)
	if err != nil {
		return err
	}
	t.Content = content

	return nil
}

func terminalElementFromBytes(bytes []byte) (TerminalElement, error) {
	var unmarshaled interface{}
	if err := json.Unmarshal(bytes, &unmarshaled); err != nil {
		return nil, err
	}
	if unmarshaled == nil {
		return nil, nil
	}

	asserted, ok := unmarshaled.(map[string]interface{}) //type assertion
	if !ok {
		return nil, fmt.Errorf("perhaps the given JSON is not a JSON 'object', as it is unmarshaled to type = %v", reflect.TypeOf(unmarshaled))
	}

	typename, ok := asserted["__typename"]
	if !ok {
		return nil, fmt.Errorf("\"__typename\" does not exist in JSON")
	}

	switch t := typename.(type) {
	case string:
		switch t {
		case "TerminalCommand":
			var cmd TerminalCommand
			if err := json.Unmarshal(bytes, &cmd); err != nil {
				return nil, err
			}
			return &cmd, nil

		case "TerminalOutput":
			var output TerminalOutput
			if err := json.Unmarshal(bytes, &output); err != nil {
				return nil, err
			}

			return &output, nil

		default:
			return nil, fmt.Errorf("\"__typename\" = %s is not a valid TerminalElement type", t)
		}
	default:
		return nil, fmt.Errorf("\"__typename\" = %v is in wrong type %v", t, reflect.TypeOf(t))
	}
}
