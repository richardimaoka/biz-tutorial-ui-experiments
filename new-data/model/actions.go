package model

import (
	"encoding/json"
	"fmt"
	"os"
)

type Action interface {
	WriteJsonToFile(filepath string) error
	Enrich(operation FileSystemOperation) (Action, error)
}

// ActionCommand represents each row of spreadsheet where type = "ActionCommand"
type ActionCommand struct {
	Command          string     `json:"command"`
	TerminalName     string     `json:"terminalName"`
	Output           *string    `json:"output"`           //if zero value, no output after execution
	CurrentDirectory *string    `json:"currentDirectory"` //if zero value, current directory is not changed after execution
	Effect           DiffEffect `json:"effect"`
}

type ManualUpdate struct {
	Effect DiffEffect `json:"effect"`
}

func (a ActionCommand) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["actionType"] = "ActionCommand"
	m["command"] = &a.Command
	m["terminalName"] = &a.TerminalName
	m["output"] = a.Output
	m["currentDirectory"] = a.CurrentDirectory
	m["effect"] = a.Effect

	return json.Marshal(m)
}

func (a ManualUpdate) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["actionType"] = "ManualUpdate"
	m["effect"] = a.Effect

	return json.Marshal(m)
}

func (a *ActionCommand) UnmarshalJSON(data []byte) error {
	typeName, err := extractTypeName(data, "actionType")
	if err != nil {
		return fmt.Errorf("ActionCommand.UnmarshalJSON() failed to extract type name: %s", err)
	}
	if typeName != "ActionCommand" {
		return fmt.Errorf("ActionCommand.UnmarshalJSON() expected type name to be ActionCommand, got %s", typeName)
	}

	type skelton struct {
		Command          string      `json:"command"`
		TerminalName     string      `json:"terminalName"`
		Output           *string     `json:"output"`
		CurrentDirectory *string     `json:"currentDirectory"`
		Effect           interface{} `json:"effect"`
	}
	var s skelton
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ActionCommand.UnmarshalJSON() failed to unmarshal: %s", err)
	}

	a.Command = s.Command
	a.TerminalName = s.TerminalName
	a.Output = s.Output
	a.CurrentDirectory = s.CurrentDirectory

	if s.Effect != nil {
		remarshaledEffect, err := json.Marshal(s.Effect)
		if err != nil {
			return fmt.Errorf("ActionCommand.UnmarshalJSON() failed to re-marshal effect: %s", err)
		}

		a.Effect, err = unmarshalDiffEffect(remarshaledEffect)
		if err != nil {
			return fmt.Errorf("ActionCommand.UnmarshalJSON() failed to unmarshal effect: %s", err)
		}
	}
	return nil
}

func (a *ManualUpdate) UnmarshalJSON(data []byte) error {
	typeName, err := extractTypeName(data, "actionType")
	if err != nil {
		return fmt.Errorf("ManualUpdate.UnmarshalJSON() failed to extract type name: %s", err)
	}
	if typeName != "ManualUpdate" {
		return fmt.Errorf("ManualUpdate.UnmarshalJSON() expected type name to be ManualUpdate, got %s", typeName)
	}

	type skelton struct {
		Effect interface{} `json:"effect"`
	}
	var s skelton
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ManualUpdate.UnmarshalJSON() failed to unmarshal: %s", err)
	}

	if s.Effect != nil {
		remarshaledEffect, err := json.Marshal(s.Effect)
		if err != nil {
			return fmt.Errorf("ActionCommand.UnmarshalJSON() failed to re-marshal effect: %s", err)
		}

		a.Effect, err = unmarshalDiffEffect(remarshaledEffect)
		if err != nil {
			return fmt.Errorf("ActionCommand.UnmarshalJSON() failed to unmarshal effect: %s", err)
		}
	}
	return nil
}

func (a ManualUpdate) WriteJsonToFile(filePath string) error {
	bytes, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filePath, bytes, 0644); err != nil {
		return err
	}
	return nil
}

func (a ActionCommand) WriteJsonToFile(filePath string) error {
	bytes, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filePath, bytes, 0644); err != nil {
		return err
	}
	return nil
}

func (a ActionCommand) Enrich(op FileSystemOperation) (Action, error) {
	var err error
	if a.Effect, err = AppendDiffEffect(a.Effect, op); err != nil {
		return nil, fmt.Errorf("ActionCommand.Enrich() failed to append diff effect: %s", err)
	}
	return a, nil
}

func (m ManualUpdate) Enrich(op FileSystemOperation) (Action, error) {
	var err error
	if m.Effect, err = AppendDiffEffect(m.Effect, op); err != nil {
		return nil, fmt.Errorf("ActionCommand.Enrich() failed to append diff effect: %s", err)
	}
	return m, nil
}

func unmarshalAction(bytes []byte) (Action, error) {
	typeName, err := extractTypeName(bytes, "actionType")
	if err != nil {
		return nil, fmt.Errorf("unmarshalAction() failed to extract actionType %s", err)
	}

	switch typeName {
	case "ActionCommand":
		var action ActionCommand
		err := json.Unmarshal(bytes, &action)
		return action, err
	case "ManualUpdate":
		var action ManualUpdate
		err := json.Unmarshal(bytes, &action)
		return action, err
	default:
		return nil, fmt.Errorf("unmarshalAction() found invalid actionType = %s", typeName)
	}
}

func readAction(filePath string) (Action, error) {
	jsonBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return unmarshalAction(jsonBytes)
}
