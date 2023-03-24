package model

import (
	"encoding/json"
	"fmt"
	"os"
)

// arbitrary JSON obj representation in Go map
type JsonObj map[string]interface{}

type Action interface {
	WriteJsonToFile(filepath string) error
	Enrich(operation FileSystemOperation) error
}

// ActionCommand represents each row of spreadsheet where type = "ActionCommand"
type ActionCommand struct {
	Command          string        `json:"command"`
	TerminalName     string        `json:"terminalName"`
	Output           *string       `json:"output"`           //if zero value, no output after execution
	CurrentDirectory *string       `json:"currentDirectory"` //if zero value, current directory is not changed after execution
	FileDiff         GitDiff       `json:"fileDiff"`
	DirectoryDiff    DirectoryDiff `json:"directoryDiff"`
}

type ManualUpdate struct {
	FileDiff      GitDiff       `json:"fileDiff"`
	DirectoryDiff DirectoryDiff `json:"directoryDiff"`
}

func (c ActionCommand) MarshalJSON() ([]byte, error) {
	typeName := "ActionCommand"

	m := make(map[string]interface{})
	m["actionType"] = &typeName
	m["command"] = &c.Command
	m["terminalName"] = &c.TerminalName
	m["output"] = c.Output
	m["currentDirectory"] = c.CurrentDirectory

	if c.FileDiff.size() > 0 && c.DirectoryDiff.size() > 0 {
		return nil, fmt.Errorf("ActionCommand's FileDiff and DirectoryDiff cannot co-exist")
	}
	m["fileDiff"] = c.FileDiff
	m["directoryDiff"] = c.DirectoryDiff

	return json.Marshal(m)
}

func (c ManualUpdate) MarshalJSON() ([]byte, error) {
	typeName := "ManualUpdate"

	m := make(map[string]*string)
	m["actionType"] = &typeName

	return json.Marshal(m)
}

func (c ManualUpdate) WriteJsonToFile(filePath string) error {
	bytes, err := json.MarshalIndent(c, "", "  ")
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

// func enrichAction(action Action, operation FileSystemOperation) Action {
// 	action.effect = action.effect.append(operation)
// }

func (c *ActionCommand) Enrich(op FileSystemOperation) error {
	switch v := op.(type) {
	case FileAdd:
		if c.DirectoryDiff.size() > 0 {
			return fmt.Errorf("ActionCommand.Enrich() received FileAdd operation while DirectoryDiff is populated")
		}
		c.FileDiff.Added = append(c.FileDiff.Added, v)
	case FileDelete:
		if c.DirectoryDiff.size() > 0 {
			return fmt.Errorf("ActionCommand.Enrich() received FileDelete operation while DirectoryDiff is populated")
		}
		c.FileDiff.Deleted = append(c.FileDiff.Deleted, v)
	case FileUpdate:
		if c.DirectoryDiff.size() > 0 {
			return fmt.Errorf("ActionCommand.Enrich() received FileDelete operation while DirectoryDiff is populated")
		}
		c.FileDiff.Updated = append(c.FileDiff.Updated, v)
	case DirectoryAdd:
		if c.FileDiff.size() > 0 {
			return fmt.Errorf("ActionCommand.Enrich() received DirectoryAdd operation while GitDiff is populated")
		}
		c.DirectoryDiff.Added = append(c.DirectoryDiff.Added, v)
	case DirectoryDelete:
		if c.FileDiff.size() > 0 {
			return fmt.Errorf("ActionCommand.Enrich() received DirectoryDelete operation while GitDiff is populated")
		}
		c.DirectoryDiff.Deleted = append(c.DirectoryDiff.Deleted, v)
	default:
		return fmt.Errorf("ActionCommand.Enrich() received invalid operation type = %T", op)
	}

	return nil
}

func (c *ManualUpdate) Enrich(op FileSystemOperation) error {
	switch v := op.(type) {
	case FileAdd:
		if c.DirectoryDiff.size() > 0 {
			return fmt.Errorf("ManualUpdate.Enrich() received FileAdd operation while DirectoryDiff is populated")
		}
		c.FileDiff.Added = append(c.FileDiff.Added, v)
	case FileDelete:
		if c.DirectoryDiff.size() > 0 {
			return fmt.Errorf("ManualUpdate.Enrich() received FileDelete operation while DirectoryDiff is populated")
		}
		c.FileDiff.Deleted = append(c.FileDiff.Deleted, v)
	case FileUpdate:
		if c.DirectoryDiff.size() > 0 {
			return fmt.Errorf("ManualUpdate.Enrich() received FileDelete operation while DirectoryDiff is populated")
		}
		c.FileDiff.Updated = append(c.FileDiff.Updated, v)
	case DirectoryAdd:
		if c.FileDiff.size() > 0 {
			return fmt.Errorf("ManualUpdate.Enrich() received DirectoryAdd operation while GitDiff is populated")
		}
		c.DirectoryDiff.Added = append(c.DirectoryDiff.Added, v)
	case DirectoryDelete:
		if c.FileDiff.size() > 0 {
			return fmt.Errorf("ManualUpdate.Enrich() received DirectoryDelete operation while GitDiff is populated")
		}
		c.DirectoryDiff.Deleted = append(c.DirectoryDiff.Deleted, v)
	default:
		return fmt.Errorf("ManualUpdate.Enrich() received invalid operation type = %T", op)
	}

	return nil
}

func readActionFromBytes(bytes []byte) (Action, error) {
	typeName, err := extractTypeName(bytes, "actionType")
	if err != nil {
		return nil, fmt.Errorf("readActionFromBytes() failed to extract actionType %s", err)
	}

	switch typeName {
	case "ActionCommand":
		var action ActionCommand
		err := json.Unmarshal(bytes, &action)
		if err != nil {
			return nil, err
		}
		if action.FileDiff.size() > 0 && action.DirectoryDiff.size() > 0 {
			return nil, fmt.Errorf("readActionFromBytes() failed as FileDiff and DirectoryDiff cannot co-exist")
		}
		return &action, nil
	case "ManualUpdate":
		var action ManualUpdate
		err := json.Unmarshal(bytes, &action)
		if err != nil {
			return nil, err
		}
		return &action, nil
	default:
		return nil, fmt.Errorf("readActionFromBytes() found invalid actionType = %s", typeName)
	}
}

func readActionFromFile(filePath string) (Action, error) {
	jsonBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return readActionFromBytes(jsonBytes)
}
