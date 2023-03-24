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

func (c ActionCommand) IsAction() {}

type ManualUpdate struct {
	FileDiff      GitDiff       `json:"fileDiff"`
	DirectoryDiff DirectoryDiff `json:"directoryDiff"`
}

func (c ManualUpdate) IsAction() {}

// methods

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

func jsonArrayFromFile(filename string) ([]JsonObj, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("reading %s failed, %s", filename, err)
	}

	var unmarshalled []JsonObj
	if err := json.Unmarshal(bytes, &unmarshalled); err != nil {
		return nil, fmt.Errorf("unmarshaling %s failed, %s", filename, err)
	}

	return unmarshalled, nil
}

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

func targetFileName(targetDir, targetPrefix string, index int) string {
	return fmt.Sprintf("%s/%s%03d.json", targetDir, targetPrefix, index)
}

func SplitActionList(actionListFile, targetDir, targetPrefix string) error {
	errorPreceding := "Error in SplitInputListFile for filename = " + actionListFile

	// read and process the whole file
	jsonArray, err := jsonArrayFromFile(actionListFile)
	if err != nil {
		return fmt.Errorf("%s, %s", errorPreceding, err)
	}

	// write each array element into file
	for i, jsonObj := range jsonArray {
		jsonBytes, err := json.Marshal(jsonObj)
		if err != nil {
			return fmt.Errorf("%s, marshaling JSON failed, %s", errorPreceding, err)
		}

		// unmarshal to action once, to control the re-marshaling behavior
		action, err := readActionFromBytes(jsonBytes)
		if err != nil {
			return fmt.Errorf("%s, reading actoin failed, %s", errorPreceding, err)
		}

		targetFile := targetFileName(targetDir, targetPrefix, i)
		if action.WriteJsonToFile(targetFile) != nil {
			return fmt.Errorf("%s, writing JSON to %s failed, %s", errorPreceding, targetFile, err)
		}
	}

	return nil
}

func EnrichActionFiles(opsListFile, actionDir, targetDir, actionPrefix string) error {
	errorPreceding := "Error in EnrichActionFiles"

	// load actions into memory
	var actions []Action

	actionFiles, err := FilesInDir(actionDir, actionPrefix)
	if err != nil {
		return fmt.Errorf("%s, %s", errorPreceding, err)
	}

	for i, file := range actionFiles {
		expectedFileName := targetFileName(actionDir, actionPrefix, i)
		if expectedFileName != file {
			return fmt.Errorf("%s, expected file %s, got %s", errorPreceding, expectedFileName, file)
		}
		action, err := readActionFromFile(file)
		if err != nil {
			return fmt.Errorf("%s, reading action file failed, %s", errorPreceding, err)
		}
		actions = append(actions, action)
	}

	// read operations and enrich actions
	jsonOpsArray, err := jsonArrayFromFile(opsListFile)
	if err != nil {
		return fmt.Errorf("%s, %s", errorPreceding, err)
	}

	for _, jsonObj := range jsonOpsArray {
		opBytes, err := json.Marshal(jsonObj)
		if err != nil {
			return fmt.Errorf("%s, marshaling operation JSON failed, %s", errorPreceding, err)
		}
		operation, err := readOperationFromBytes(opBytes)
		if err != nil {
			return fmt.Errorf("%s, reading operation failed, %s", errorPreceding, err)
		}

		seqNo, ok := jsonObj["seqNo"]
		if !ok {
			return fmt.Errorf("%s, seqNo not found in JSON = %s, %s", errorPreceding, opBytes, err)
		}
		seqNoFloat, ok := seqNo.(float64)
		if !ok {
			return fmt.Errorf("%s, seqNo not number in JSON = %s, %s", errorPreceding, opBytes, err)
		}
		seqNoInt := int(seqNoFloat)
		if len(actions) <= seqNoInt {
			return fmt.Errorf("%s, seqNo = %d is out of range, %s", errorPreceding, seqNoInt, err)
		}

		if err := actions[seqNoInt].Enrich(operation); err != nil {
			return fmt.Errorf("%s, enriching action %d failed, %s", errorPreceding, seqNoInt, err)
		}
	}

	// write enriched actions to files
	for i, action := range actions {
		targetFile := targetFileName(targetDir, actionPrefix, i)
		if err := action.WriteJsonToFile(targetFile); err != nil {
			return fmt.Errorf("%s, writing JSON to %s failed, %s", errorPreceding, targetFile, err)
		}
	}

	return nil
}

func ApplyActions(actionDir, actionPrefix string) error {
	errorPreceding := "Error in ApplyActions"

	actionFiles, err := FilesInDir(actionDir, actionPrefix)
	if err != nil {
		return fmt.Errorf("%s, %s", errorPreceding, err)
	}

	// pageState := NewPageState()
	for _, file := range actionFiles {
		action, err := readActionFromFile(file)
		if err != nil {
			return fmt.Errorf("%s, reading action file failed, %s", errorPreceding, err)
		}

		action.IsAction()
		// if err := pageState.processAction(action); err != nil {
		// 	return fmt.Errorf("%s, applying action failed, %s", errorPreceding, err)
		// }
	}

	return nil
}

func Processing() error {
	// 0. prereuisite: by-hand csv -> json conversion, and save action-list.json

	// 1. split action-list.json
	inputDir := "data/input"
	prefix := "action"
	if err := SplitActionList("data/action_list.json", inputDir, prefix); err != nil {
		return err
	}

	// 2. enrich action files
	enrichedDir := "data/enriched"
	if err := EnrichActionFiles("data/source_code_ops.json", inputDir, enrichedDir, prefix); err != nil {
		return err
	}

	// 3. apply action files
	if err := ApplyActions(enrichedDir, prefix); err != nil {
		return err
	}

	return nil
}
