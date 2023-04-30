package processing

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type Action interface {
	Enrich(operation FileSystemOperation)
	ToGraphQLNextAction() model.NextAction
}

// ActionCommand represents each row of spreadsheet where type = "ActionCommand"
type ActionCommand struct {
	// fields for terminal
	TerminalName     string  `json:"terminalName"`
	Command          string  `json:"command"`
	Output           *string `json:"output"`           //if zero value, no output after execution
	CurrentDirectory *string `json:"currentDirectory"` //if zero value, current directory is not changed after execution

	// fields for source code
	Diff                Diff    `json:"diff"`
	DefaultOpenFilePath *string `json:"defaultOpenFilePath"`
}

type ManualUpdate struct {
	// fields for source code
	Diff                Diff    `json:"diff"`
	DefaultOpenFilePath *string `json:"defaultOpenFilePath"`
}

func (a *ActionCommand) Enrich(op FileSystemOperation) {
	a.Diff.Append(op)
}

func (a *ManualUpdate) Enrich(op FileSystemOperation) {
	a.Diff.Append(op)
}

func (a *ActionCommand) ToGraphQLNextAction() model.NextAction {
	return model.NextAction{
		Content: model.NextActionTerminal{
			Command: &a.Command,
		},
	}
}

func (a *ManualUpdate) ToGraphQLNextAction() model.NextAction {
	fixedComment := "manual action"
	return model.NextAction{
		Content: model.NextActionManual{
			Comment: &fixedComment,
		},
	}
}

// ------------------------------
// JSON marshal/unmarshal
// ------------------------------

// to output "actionType" in JSON
func (a ActionCommand) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["actionType"] = "ActionCommand" // this has to be added

	// fields for terminal
	m["terminalName"] = &a.TerminalName
	m["command"] = &a.Command
	m["output"] = a.Output
	m["currentDirectory"] = a.CurrentDirectory

	// fields for source code
	m["diff"] = a.Diff
	m["defaultOpenFilePath"] = a.DefaultOpenFilePath

	return json.Marshal(m)
}

// to output "actionType" in JSON
func (a ManualUpdate) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["actionType"] = "ManualUpdate" // this has to be added

	// fields for source code
	m["diff"] = a.Diff
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

	// 2. read into a typed struct
	//    a different struct is needed to for *typed* reading of JSON
	//    also, if you use the same receiver struct in UnmarshalJSON, then infinite loop
	var interim struct {
		// actionType is removed

		// fields for terminal
		TerminalName     string  `json:"terminalName"`
		Command          string  `json:"command"`
		Output           *string `json:"output"`
		CurrentDirectory *string `json:"currentDirectory"`

		// fields for source code
		Diff                Diff    `json:"diff"`
		DefaultOpenFilePath *string `json:"defaultOpenFilePath"`
	}
	if err := json.Unmarshal(data, &interim); err != nil {
		return fmt.Errorf("ActionCommand.UnmarshalJSON() failed to unmarshal: %s", err)
	}

	// 3. *typed* assignment from here
	//    fields for source code
	a.TerminalName = interim.TerminalName
	a.Command = interim.Command
	a.Output = interim.Output
	a.CurrentDirectory = interim.CurrentDirectory

	// fields for source code
	a.DefaultOpenFilePath = interim.DefaultOpenFilePath
	a.Diff = interim.Diff

	return nil
}

// used in func unmarshalAction()
func (a *ManualUpdate) UnmarshalJSON(data []byte) error {
	// 1. actionType check
	typeName, err := extractTypeName(data, "actionType")
	if err != nil {
		return fmt.Errorf("ManualUpdate.UnmarshalJSON() failed to extract type name: %s", err)
	}
	if typeName != "ManualUpdate" {
		return fmt.Errorf("ManualUpdate.UnmarshalJSON() expected type name to be ManualUpdate, got %s", typeName)
	}

	// 2. read into a typed struct
	//    a different struct is needed to for *typed* reading of JSON
	//    also, if you use the same receiver struct in UnmarshalJSON, then infinite loop
	var interim struct {
		// actionType is removed

		// fields for source code
		Diff                Diff    `json:"diff"`
		DefaultOpenFilePath *string `json:"defaultOpenFilePath"`
	}
	if err := json.Unmarshal(data, &interim); err != nil {
		return fmt.Errorf("ManualUpdate.UnmarshalJSON() failed to unmarshal: %s", err)
	}

	// 3. *typed* assignment from here
	//    fields for source code
	a.DefaultOpenFilePath = interim.DefaultOpenFilePath
	a.Diff = interim.Diff

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
		return &action, err
	case "ManualUpdate":
		var action ManualUpdate
		err := json.Unmarshal(bytes, &action) // calls UnmarshalJSON()
		return &action, err
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
