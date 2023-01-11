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

func getAction(filename string) (Action, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error in filename = %s, %s", filename, err)
	}

	var data interface{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, fmt.Errorf("Error in filename = %s, while unmarshaling JSON from file, %s", filename, err)
	}

	asserted, ok := data.(map[string]interface{}) //type assertion
	if !ok {
		return nil, fmt.Errorf("Error in filename = %s, while constructing Go data from JSON, perhaps the file is not in JSON 'object'", filename)
	}

	typename, ok := asserted["__typename"]
	if !ok {
		return nil, fmt.Errorf("Error in filename = %s, while validating action type: \"__typename\" does not exist", filename)
	}

	switch t := typename.(type) {
	case string:
		switch t {
		case "Command":
			fmt.Printf("successfully retrieved action = %s from %s\n", t, filename)
			var command Command
			if err := json.Unmarshal(bytes, &command); err != nil {
				return nil, fmt.Errorf("Error in filename = %s, while unmarshaling JSON from file, %s", filename, err)
			}

			return &command, nil
		default:
			return nil, fmt.Errorf("Error in filename = %s, while validating action type: %s is not a valid action type", filename, t)
		}
	default:
		return nil, fmt.Errorf("Error in filename = %s, while validating action type: \"__typename\" = %v is in wrong type %v", filename, t, reflect.TypeOf(t))
	}

}

func main() {
	filename := "step01/action.json"
	action, err := getAction(filename)
	if err != nil {
		panic(err)
	}
	fmt.Print(action)
}
