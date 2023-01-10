package main

import (
	"encoding/json"
	"fmt"
	"os"
)

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

type Action interface {
	IsAction()
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
	bytes, err := os.ReadFile("action01.json")
	if err != nil {
		panic(err)
	}

	var data interface{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		panic(err)
	}

	asserted, ok := data.(map[string]interface{})
	if !ok {
		panic("type assertion failed")
	}

	action, ok := asserted["action"]
	if !ok {
		panic("action does not exist")
	}

	switch v := action.(type) {
		case Command
	}

	fmt.Println(actions)
}
