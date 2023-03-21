package model

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// arbitrary JSON obj representation in Go map
type JsonObj map[string]interface{}

type Action interface {
	IsAction()
}

// ActionCommand represents each row of spreadsheet where type = "ActionCommand"
type ActionCommand struct {
	Command          string         `json:"command"`
	TerminalName     string         `json:"terminalName"`
	Output           *string        `json:"output"`           //if zero value, no output after execution
	CurrentDirectory *string        `json:"currentDirectory"` //if zero value, current directory is not changed after execution
	FileDiff         *GitDiff       `json:"fileDiff"`
	DirectoryDiff    *DirectoryDiff `json:"directoryDiff"`
}

func (c *ActionCommand) IsAction() {}

type ManualUpdate struct {
}

func (c *ManualUpdate) IsAction() {}

// methods

func (c ActionCommand) MarshalJSON() ([]byte, error) {
	typeName := "ActionCommand"

	m := make(map[string]interface{})
	m["actionType"] = &typeName
	m["command"] = &c.Command
	m["terminalName"] = &c.TerminalName
	m["output"] = c.Output
	m["currentDirectory"] = c.CurrentDirectory

	if c.FileDiff != nil && c.DirectoryDiff != nil {
		return nil, fmt.Errorf("ActionCommand's FileDiff and DirectoryDiff cannot co-exist")
	} else if c.FileDiff != nil {
		m["fileDiff"] = c.FileDiff
	} else if c.DirectoryDiff != nil {
		m["directoryDiff"] = c.DirectoryDiff
	}

	return json.Marshal(m)
}

func (c ManualUpdate) MarshalJSON() ([]byte, error) {
	typeName := "ManualUpdate"

	m := make(map[string]*string)
	m["actionType"] = &typeName

	return json.Marshal(m)
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
		return &action, nil
	case "ManualUpdate":
		var action ManualUpdate
		err := json.Unmarshal(bytes, &action)
		if err != nil {
			return nil, err
		}
		return &action, nil
	default:
		return nil, fmt.Errorf("readActionFromBytes() found invalid typeName = %s", typeName)
	}
}

func reMarshalAction(bytes []byte) ([]byte, error) {
	action, err := readActionFromBytes(bytes)
	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(action, "", "  ")
}

// map[string]interface{} represents JSON obj
// return a slice of map[string]interface{} (i.e.) []map[string]interface{}
func readActionList(actionListFile string) ([]JsonObj, error) {
	bytes, err := os.ReadFile(actionListFile)
	if err != nil {
		return nil, fmt.Errorf("reading %s failed, %s", actionListFile, err)
	}

	var unmarshalled []JsonObj
	if err := json.Unmarshal(bytes, &unmarshalled); err != nil {
		return nil, fmt.Errorf("unmarshaling %s failed, %s", actionListFile, err)
	}

	return unmarshalled, nil
}

func applyAction(bytes []byte) error {
	// action, err := readActionFromBytes(bytes)
	// if err != nil {
	// 	return err
	// }

	return nil
}

// all input_flat00x files
func FilesInDir(targetDir, prefix string) ([]string, error) {
	entries, err := os.ReadDir(targetDir)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, e := range entries {
		if strings.HasPrefix(e.Name(), prefix) && strings.HasSuffix(e.Name(), "json") {
			files = append(files, targetDir+"/"+e.Name())
		}
	}

	return files, nil
}

func SplitActionList(actionListFile, targetDir, targetPrefix string) error {
	errorPreceding := "Error in SplitInputListFile for filename = " + actionListFile

	// read and process the whole file
	jsonArray, err := readActionList(actionListFile)
	if err != nil {
		return fmt.Errorf("%s, %s", errorPreceding, err)
	}

	// write each array element into file
	for i, jsonObj := range jsonArray {
		jsonBytes, err := json.MarshalIndent(jsonObj, "", "  ")
		if err != nil {
			return fmt.Errorf("marshaling JSON failed, %s", err)
		}

		actionBytes, err := reMarshalAction(jsonBytes)
		if err != nil {
			return fmt.Errorf("re-marshaling JSON failed, %s", err)
		}
		targetFile := fmt.Sprintf("%s/%s%03d.json", targetDir, targetPrefix, i)
		if err = os.WriteFile(targetFile, actionBytes, 0644); err != nil {
			return fmt.Errorf("%s, writing JSON to %s failed, %s", errorPreceding, targetFile, err)
		}
	}

	return nil
}

func Processing() error {
	// 1. prereuisite: by-hand csv -> json conversion, and save action-list.json

	inputDir := "data/input"
	prefix := "action"

	// 2. split action-list.json
	if err := SplitActionList("data/action_list.json", inputDir, prefix); err != nil {
		return err
	}

	files, err := FilesInDir(inputDir, prefix)
	if err != nil {
		return err
	}

	// 3.
	for _, f := range files {
		fmt.Println(f)
	}

	// GenerateInputActionFiles("")

	return nil
}
