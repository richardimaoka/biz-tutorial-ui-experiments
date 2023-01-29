package pkg2

import (
	"encoding/json"
	"fmt"
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

func ExtractTypeName(bytes []byte, fromField string) (string, error) {
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
		return "", fmt.Errorf("\"%s\" is not a string, but found %s", fromField, reflect.TypeOf(typeNameRaw))
	}

	return typeName, nil
}
