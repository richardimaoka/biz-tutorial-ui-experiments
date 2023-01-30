package pkg2

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type AddDirectory struct {
	FilePath []string
}

type DeleteDirectory struct {
	FilePath []string
}

type AddFile struct {
	FilePath []string
	Content  string
}

type UpdateFile struct {
	FilePath []string
	Content  string
}

type DeleteFile struct {
	FilePath []string
}

type UpdateSourceCode struct {
	AddDirectories    []AddDirectory
	DeleteDirectories []DeleteDirectory
	AddFiles          []AddFile
	UpdateFiles       []UpdateFile
	DeleteFiles       []DeleteFile
}

type UpdateTerminal struct {
	Output           string
	CurrentDirectory []string
}

type ActionCommand struct {
	ActionType       string
	Command          string
	TerminalName     string
	UpdateTerminal   UpdateTerminal
	UpdateSourceCode UpdateSourceCode
}

func extractTypeName(bytes []byte, fromField string) (string, error) {
	var unmarshaled map[string]interface{}
	if err := json.Unmarshal(bytes, &unmarshaled); err != nil {
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

func readAction(filePath string) (*ActionCommand, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("readAction() failed, %s", err)
	}

	typeName, err := extractTypeName(bytes, "actionType")
	if err != nil {
		return nil, fmt.Errorf("readAction() failed to actionType from %s, %s", filePath, err)
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
		return nil, fmt.Errorf("readAction() found invalid typeName = %s in file = %s", typeName, filePath)
	}
}

func Process(actionFile string, stepFile string) error {
	_, err := readAction(actionFile)
	if err != nil {
		return err
	}

	return nil
}
