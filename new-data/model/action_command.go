package model

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func (c ActionCommand) MarshalJSON() ([]byte, error) {
	fmt.Println("ActionCommand MarshalJSON")
	extendedCommand := struct {
		ActionType       string
		Command          string
		TerminalName     string
		UpdateTerminal   UpdateTerminal
		UpdateSourceCode UpdateSourceCode
	}{
		"ActionCommand",
		c.Command,
		c.TerminalName,
		c.UpdateTerminal,
		c.UpdateSourceCode,
	}

	return json.Marshal(extendedCommand)
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

func readActionFromBytes(bytes []byte) (*ActionCommand, error) {
	typeName, err := extractTypeName(bytes, "actionType")
	if err != nil {
		return nil, fmt.Errorf("readActionFromBytes() failed to extract actionType %s", err)
	}

	switch typeName {
	case "ActionCommand":
		var action ActionCommand
		err := json.Unmarshal(bytes, &action)
		if err != nil {
			return nil, err
		}
		return &action, nil
	default:
		return nil, fmt.Errorf("readActionFromBytes() found invalid typeName = %s", typeName)
	}
}
