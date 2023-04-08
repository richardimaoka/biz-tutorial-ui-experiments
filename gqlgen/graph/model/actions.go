package model

import (
	"encoding/json"
	"fmt"
	"os"
)

type Action interface {
	Enrich(operation FileSystemOperation) (Action, error)
}

// ActionCommand represents each row of spreadsheet where type = "ActionCommand"
type ActionCommand struct {
	Command          string     `json:"command"`
	TerminalName     string     `json:"terminalName"`
	Output           *string    `json:"output"`           //if zero value, no output after execution
	CurrentDirectory *string    `json:"currentDirectory"` //if zero value, current directory is not changed after execution
	Effect           DiffEffect `json:"effect"`
	OpenFilePath     *string    `json:"openFilePath"`
}

type ManualUpdate struct {
	Effect       DiffEffect `json:"effect"`
	OpenFilePath *string    `json:"openFilePath"`
}

func (a ActionCommand) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["actionType"] = "ActionCommand"
	m["command"] = &a.Command
	m["terminalName"] = &a.TerminalName
	m["output"] = a.Output
	m["currentDirectory"] = a.CurrentDirectory
	m["effect"] = a.Effect
	m["openFilePath"] = a.OpenFilePath

	return json.Marshal(m)
}

func (a ManualUpdate) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["actionType"] = "ManualUpdate"
	m["effect"] = a.Effect
	m["openFilePath"] = a.OpenFilePath

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

	var interim struct {
		Command          string      `json:"command"`
		TerminalName     string      `json:"terminalName"`
		Output           *string     `json:"output"`
		CurrentDirectory *string     `json:"currentDirectory"`
		Effect           interface{} `json:"effect"`
	}
	if err := json.Unmarshal(data, &interim); err != nil {
		return fmt.Errorf("ActionCommand.UnmarshalJSON() failed to unmarshal: %s", err)
	}

	a.Command = interim.Command
	a.TerminalName = interim.TerminalName
	a.Output = interim.Output
	a.CurrentDirectory = interim.CurrentDirectory

	if interim.Effect != nil {
		remarshaledEffect, err := json.Marshal(interim.Effect)
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

	var interim struct {
		Effect interface{} `json:"effect"`
	}
	if err := json.Unmarshal(data, &interim); err != nil {
		return fmt.Errorf("ManualUpdate.UnmarshalJSON() failed to unmarshal: %s", err)
	}

	if interim.Effect != nil {
		remarshaledEffect, err := json.Marshal(interim.Effect)
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
