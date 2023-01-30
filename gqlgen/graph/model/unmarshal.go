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

	contentType := "contentType"
	typename, ok := asserted[contentType]
	if !ok {
		return nil, fmt.Errorf("\"%s\" does not exist in JSON", contentType)
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
			return nil, fmt.Errorf("\"%s\" = %s is not a valid TerminalElement type", contentType, t)
		}
	default:
		return nil, fmt.Errorf("\"%s\" = %v is in wrong type %v", contentType, t, reflect.TypeOf(t))
	}
}
