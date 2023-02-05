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

func extractTypeName(jsonBytes []byte, fromField string) (string, error) {
	var unmarshaled map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &unmarshaled); err != nil {
		return "", err
	}

	typeNameRaw, ok := unmarshaled[fromField]
	if !ok {
		return "", fmt.Errorf("\"%s\" does not exist in JSON", fromField)
	}

	typeName, ok := typeNameRaw.(string)
	if !ok {
		return "", fmt.Errorf("\"%s\" is not a string, but found in type = %v", fromField, reflect.TypeOf(typeNameRaw))
	}

	return typeName, nil
}

func terminalElementFromBytes(bytes []byte) (TerminalElement, error) {
	fromField := "contentType"
	typename, err := extractTypeName(bytes, fromField)
	if err != nil {
		return nil, err
	}

	switch typename {
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
		return nil, fmt.Errorf("\"%s\" = %s is not a valid TerminalElement type", fromField, typename)
	}
}
