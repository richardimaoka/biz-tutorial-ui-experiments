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
	Command          string        `json:"command"`
	TerminalName     string        `json:"terminalName"`
	Output           *string       `json:"output"`           //if zero value, no output after execution
	CurrentDirectory *string       `json:"currentDirectory"` //if zero value, current directory is not changed after execution
	FileDiff         GitDiff       `json:"fileDiff"`
	DirectoryDiff    DirectoryDiff `json:"directoryDiff"`
	Effect           DiffEffect    `json:"effect"`
}

type ManualUpdate struct {
	FileDiff      GitDiff       `json:"fileDiff"`
	DirectoryDiff DirectoryDiff `json:"directoryDiff"`
	Effect        DiffEffect    `json:"effect"`
}

func (c ActionCommand) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["actionType"] = "ActionCommand"
	m["command"] = &c.Command
	m["terminalName"] = &c.TerminalName
	m["output"] = c.Output
	m["currentDirectory"] = c.CurrentDirectory

	if c.FileDiff.size() > 0 && c.DirectoryDiff.size() > 0 {
		return nil, fmt.Errorf("ActionCommand's FileDiff and DirectoryDiff cannot co-exist")
	}
	m["fileDiff"] = c.FileDiff
	m["directoryDiff"] = c.DirectoryDiff

	m["effect"] = c.Effect

	// if c.effect == nil {
	// 	m["fileDiff"] = GitDiff{}
	// 	m["directoryDiff"] = DirectoryDiff{}
	// } else {
	// 	switch v := c.effect.(type) {
	// 	case GitDiff:
	// 		m["fileDiff"] = v
	// 		m["directoryDiff"] = DirectoryDiff{}
	// 	case DirectoryDiff:
	// 		m["fileDiff"] = GitDiff{}
	// 		m["directoryDiff"] = v
	// 	}
	// }

	return json.Marshal(m)
}

func (c ManualUpdate) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m["actionType"] = "ManualUpdate"

	// if c.Effect != nil {
	// 	switch v := c.Effect.(type) {
	// 	case GitDiff:
	// 		m["fileDiff"] = v
	// 	case DirectoryDiff:
	// 		m["directoryDiff"] = v
	// 	}
	// }

	return json.Marshal(m)
}

func (c *ActionCommand) UnmarshalJSON(data []byte) error {
	typeName, err := extractTypeName(data, "actionType")
	if err != nil {
		return fmt.Errorf("ActionCommand.UnmarshalJSON() failed to extract type name: %s", err)
	}
	if typeName != "ActionCommand" {
		return fmt.Errorf("ActionCommand.UnmarshalJSON() expected type name to be ActionCommand, got %s", typeName)
	}

	type skelton struct {
		Command          string        `json:"command"`
		TerminalName     string        `json:"terminalName"`
		Output           *string       `json:"output"`
		CurrentDirectory *string       `json:"currentDirectory"`
		FileDiff         GitDiff       `json:"fileDiff"`
		DirectoryDiff    DirectoryDiff `json:"directoryDiff"`
		Effect           interface{}   `json:"effect"`
	}
	var s skelton
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ActionCommand.UnmarshalJSON() failed to unmarshal: %s", err)
	}

	c.Command = s.Command
	c.TerminalName = s.TerminalName
	c.Output = s.Output
	c.CurrentDirectory = s.CurrentDirectory
	c.FileDiff = s.FileDiff
	c.DirectoryDiff = s.DirectoryDiff

	// if s.FileDiff.size() > 0 {
	// 	c.Effect = s.FileDiff
	// } else if s.DirectoryDiff.size() > 0 {
	// 	c.Effect = s.DirectoryDiff
	// }

	// if s.Effect != nil {
	// 	remarshaledEffect, err := json.Marshal(s.Effect)
	// 	if err != nil {
	// 		return fmt.Errorf("ActionCommand.UnmarshalJSON() failed to re-marshal effect: %s", err)
	// 	}

	// 	c.effect, err = unmarshalDiffEffect(remarshaledEffect)
	// 	if err != nil {
	// 		return fmt.Errorf("ActionCommand.UnmarshalJSON() failed to unmarshal effect: %s", err)
	// 	}
	// }
	return nil
}

func (m *ManualUpdate) UnmarshalJSON(data []byte) error {
	typeName, err := extractTypeName(data, "actionType")
	if err != nil {
		return fmt.Errorf("ManualUpdate.UnmarshalJSON() failed to extract type name: %s", err)
	}
	if typeName != "ManualUpdate" {
		return fmt.Errorf("ManualUpdate.UnmarshalJSON() expected type name to be ManualUpdate, got %s", typeName)
	}

	type subset struct {
		Effect interface{} `json:"effect"`
	}
	var s subset
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ManualUpdate.UnmarshalJSON() failed to unmarshal: %s", err)
	}

	// if s.Effect != nil {
	// 	remarshaledEffect, err := json.Marshal(s.Effect)
	// 	if err != nil {
	// 		return fmt.Errorf("ManualUpdate.UnmarshalJSON() failed to re-marshal effect: %s", err)
	// 	}

	// 	m.effect, err = unmarshalDiffEffect(remarshaledEffect)
	// 	if err != nil {
	// 		return fmt.Errorf("ManualUpdate.UnmarshalJSON() failed to unmarshal effect: %s", err)
	// 	}
	// }
	return nil
}

func (m ManualUpdate) WriteJsonToFile(filePath string) error {
	bytes, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filePath, bytes, 0644); err != nil {
		return err
	}
	return nil
}

func (c ActionCommand) WriteJsonToFile(filePath string) error {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filePath, bytes, 0644); err != nil {
		return err
	}
	return nil
}

func (c ActionCommand) Enrich(op FileSystemOperation) (Action, error) {
	// var err error
	// if c.effect, err = AppendDiffEffect(c.effect, op); err != nil {
	// 	return nil, fmt.Errorf("ActionCommand.Enrich() failed to append diff effect: %s", err)
	// }

	switch v := op.(type) {
	case FileAdd:
		if c.DirectoryDiff.size() > 0 {
			return nil, fmt.Errorf("ActionCommand.Enrich() received FileAdd operation while DirectoryDiff is populated")
		}
		c.FileDiff.Added = append(c.FileDiff.Added, v)
		return c, nil
	case FileDelete:
		if c.DirectoryDiff.size() > 0 {
			return nil, fmt.Errorf("ActionCommand.Enrich() received FileDelete operation while DirectoryDiff is populated")
		}
		c.FileDiff.Deleted = append(c.FileDiff.Deleted, v)
		return c, nil
	case FileUpdate:
		if c.DirectoryDiff.size() > 0 {
			return nil, fmt.Errorf("ActionCommand.Enrich() received FileDelete operation while DirectoryDiff is populated")
		}
		c.FileDiff.Updated = append(c.FileDiff.Updated, v)
		return c, nil
	case DirectoryAdd:
		if c.FileDiff.size() > 0 {
			return nil, fmt.Errorf("ActionCommand.Enrich() received DirectoryAdd operation while GitDiff is populated")
		}
		c.DirectoryDiff.Added = append(c.DirectoryDiff.Added, v)
		return c, nil
	case DirectoryDelete:
		if c.FileDiff.size() > 0 {
			return nil, fmt.Errorf("ActionCommand.Enrich() received DirectoryDelete operation while GitDiff is populated")
		}
		c.DirectoryDiff.Deleted = append(c.DirectoryDiff.Deleted, v)
		return c, nil
	default:
		return nil, fmt.Errorf("ActionCommand.Enrich() received invalid operation type = %T", op)
	}
}

func (m ManualUpdate) Enrich(op FileSystemOperation) (Action, error) {
	// var err error
	// if m.effect, err = AppendDiffEffect(m.effect, op); err != nil {
	// 	return nil, fmt.Errorf("ActionCommand.Enrich() failed to append diff effect: %s", err)
	// }

	switch v := op.(type) {
	case FileAdd:
		if m.DirectoryDiff.size() > 0 {
			return nil, fmt.Errorf("ManualUpdate.Enrich() received FileAdd operation while DirectoryDiff is populated")
		}
		m.FileDiff.Added = append(m.FileDiff.Added, v)
		return m, nil
	case FileDelete:
		if m.DirectoryDiff.size() > 0 {
			return nil, fmt.Errorf("ManualUpdate.Enrich() received FileDelete operation while DirectoryDiff is populated")
		}
		m.FileDiff.Deleted = append(m.FileDiff.Deleted, v)
		return m, nil
	case FileUpdate:
		if m.DirectoryDiff.size() > 0 {
			return nil, fmt.Errorf("ManualUpdate.Enrich() received FileDelete operation while DirectoryDiff is populated")
		}
		m.FileDiff.Updated = append(m.FileDiff.Updated, v)
		return m, nil
	case DirectoryAdd:
		if m.FileDiff.size() > 0 {
			return nil, fmt.Errorf("ManualUpdate.Enrich() received DirectoryAdd operation while GitDiff is populated")
		}
		m.DirectoryDiff.Added = append(m.DirectoryDiff.Added, v)
		return m, nil
	case DirectoryDelete:
		if m.FileDiff.size() > 0 {
			return nil, fmt.Errorf("ManualUpdate.Enrich() received DirectoryDelete operation while GitDiff is populated")
		}
		m.DirectoryDiff.Deleted = append(m.DirectoryDiff.Deleted, v)
		return m, nil
	default:
		return nil, fmt.Errorf("ManualUpdate.Enrich() received invalid operation type = %T", op)
	}
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
		if err != nil {
			return nil, err
		}
		if action.FileDiff.size() > 0 && action.DirectoryDiff.size() > 0 {
			return nil, fmt.Errorf("unmarshalAction() failed as FileDiff and DirectoryDiff cannot co-exist")
		}
		return action, nil
	case "ManualUpdate":
		var action ManualUpdate
		err := json.Unmarshal(bytes, &action)
		if err != nil {
			return nil, err
		}
		if action.FileDiff.size() > 0 && action.DirectoryDiff.size() > 0 {
			return nil, fmt.Errorf("unmarshalAction() failed as FileDiff and DirectoryDiff cannot co-exist")
		}
		return action, nil
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
