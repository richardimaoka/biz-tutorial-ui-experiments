package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type Action interface {
	IsAction()
}

type Command struct {
	TypeName string `json:"__typename"`
	Command  string
}

func (c *Command) IsAction() {}

type File struct {
	TypeName    string `json:"__typename"`
	FilePath    []string
	FileContent string
	Offset      int
}

type SourceCodeUpdate struct {
	TypeName   string `json:"__typename"`
	FilesAdded []File
}

type ChangeCurrentDirectory struct {
	TypeName   string `json:"__typename"`
	TerminalId string
	FilePath   []string
}

type Terminal struct {
	elements []interface{}
}

type ActionInfo struct {
	Action  Action
	Results interface{}
}

type State struct {
	SourceCode interface{}
	Terminal   interface{}
}

func main() {
	filename := "step01/action.json"
	bytes, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Error in reading action from file: %s", err))
	}

	var data interface{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		panic(fmt.Sprintf("Error in unmarshaling JSON from file: %s is not a valid JSON, %s", filename, err))
	}

	asserted, ok := data.(map[string]interface{}) //type assertion
	if !ok {
		panic(fmt.Sprintf("Error in constructing Go data from JSON: Perhaps %s is not in JSON 'object'", filename))
	}

	action, ok := asserted["__typename"]
	if !ok {
		panic(fmt.Sprintf("Error in validating action type: \"__typename\" does not exist in %s", filename))
	}

	switch s := action.(type) {
	case string:
		switch s {
		case "Command":
			fmt.Printf("successfully retrieved action = %s from %s", s, filename)
		default:
			panic(fmt.Sprintf("Error in validating action type: %s is not a valid action type", s))
		}
	default:
		panic(fmt.Sprintf("Error in validating action type: \"__typename\" = %v is in wrong type %v", s, reflect.TypeOf(s)))
	}

}
