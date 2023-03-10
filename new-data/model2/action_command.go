package model2

import "encoding/json"

type ActionCommand struct {
	Command          string  `json:"command"`
	TerminalName     string  `json:"terminalName"`
	Output           *string `json:"output"`           //if zero value, no output after execution
	CurrentDirectory *string `json:"currentDirectory"` //if zero value, current directory is not changed after execution
}

func (c ActionCommand) MarshalJSON() ([]byte, error) {
	typeName := "ActionCommand"

	m := make(map[string]*string)
	m["actionType"] = &typeName
	m["command"] = &c.Command
	m["terminalName"] = &c.TerminalName
	m["output"] = c.Output
	m["currentDirectory"] = c.CurrentDirectory

	return json.Marshal(m)
}
