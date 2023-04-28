package processing

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
	TerminalName        string     `json:"terminalName"`
	Command             string     `json:"command"`
	Output              *string    `json:"output"`           //if zero value, no output after execution
	CurrentDirectory    *string    `json:"currentDirectory"` //if zero value, current directory is not changed after execution
	Effect              DiffEffect `json:"effect"`
	DefaultOpenFilePath *string    `json:"defaultOpenFilePath"`
}

type ManualUpdate struct {
	Effect              DiffEffect `json:"effect"`
	DefaultOpenFilePath *string    `json:"defaultOpenFilePath"`
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

// ------------------------------
// JSON marshal/unmarshal
// ------------------------------

// to output "actionType" in JSON
func (a ActionCommand) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["actionType"] = "ActionCommand" // this has to be added
	m["effect"] = a.Effect
	m["defaultOpenFilePath"] = a.DefaultOpenFilePath

	m["command"] = &a.Command
	m["terminalName"] = &a.TerminalName
	m["output"] = a.Output
	m["currentDirectory"] = a.CurrentDirectory

	return json.Marshal(m)
}

// to output "actionType" in JSON
func (a ManualUpdate) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["actionType"] = "ManualUpdate" // this has to be added
	m["effect"] = a.Effect
	m["defaultOpenFilePath"] = a.DefaultOpenFilePath

	return json.Marshal(m)
}

// used in func unmarshalAction()
func (a *ActionCommand) UnmarshalJSON(data []byte) error {
	typeName, err := extractTypeName(data, "actionType")
	if err != nil {
		return fmt.Errorf("ActionCommand.UnmarshalJSON() failed to extract type name: %s", err)
	}
	if typeName != "ActionCommand" {
		return fmt.Errorf("ActionCommand.UnmarshalJSON() expected type name to be ActionCommand, got %s", typeName)
	}

	// a different struct is needed to for *typed* reading of JSON
	// also, if you use the same receiver struct in UnmarshalJSON, then infinite loop
	var interim struct {
		Command             string      `json:"command"`
		TerminalName        string      `json:"terminalName"`
		Output              *string     `json:"output"`
		CurrentDirectory    *string     `json:"currentDirectory"`
		Effect              interface{} `json:"effect"` // this needs furether processing based on "diffType"
		DefaultOpenFilePath *string     `json:"defaultOpenFilePath"`
	}
	if err := json.Unmarshal(data, &interim); err != nil {
		return fmt.Errorf("ActionCommand.UnmarshalJSON() failed to unmarshal: %s", err)
	}

	// *typed* assignment
	a.Command = interim.Command
	a.TerminalName = interim.TerminalName
	a.Output = interim.Output
	a.CurrentDirectory = interim.CurrentDirectory
	a.DefaultOpenFilePath = interim.DefaultOpenFilePath

	// Effect needs furether processing based on "diffType"
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

// used in func unmarshalAction()
func (a *ManualUpdate) UnmarshalJSON(data []byte) error {
	typeName, err := extractTypeName(data, "actionType")
	if err != nil {
		return fmt.Errorf("ManualUpdate.UnmarshalJSON() failed to extract type name: %s", err)
	}
	if typeName != "ManualUpdate" {
		return fmt.Errorf("ManualUpdate.UnmarshalJSON() expected type name to be ManualUpdate, got %s", typeName)
	}

	// a different struct is needed to for *typed* reading of JSON
	// also, if you use the same receiver struct in UnmarshalJSON, then infinite loop
	var interim struct {
		Effect interface{} `json:"effect"` // this needs furether processing based on "diffType"
	}
	if err := json.Unmarshal(data, &interim); err != nil {
		return fmt.Errorf("ManualUpdate.UnmarshalJSON() failed to unmarshal: %s", err)
	}

	// Effect needs furether processing based on "diffType"
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

func unmarshalAction(bytes []byte) (Action, error) {
	typeName, err := extractTypeName(bytes, "actionType")
	if err != nil {
		return nil, fmt.Errorf("unmarshalAction() failed to extract actionType %s", err)
	}

	switch typeName {
	case "ActionCommand":
		var action ActionCommand
		err := json.Unmarshal(bytes, &action) // calls UnmarshalJSON()
		return action, err
	case "ManualUpdate":
		var action ManualUpdate
		err := json.Unmarshal(bytes, &action) // calls UnmarshalJSON()
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
