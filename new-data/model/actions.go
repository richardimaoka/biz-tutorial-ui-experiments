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

func (c *ActionCommand) WriteJsonToFile(filePath string) error {
	bytes, err := json.Marshal(c)
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
		if c.DirectoryDiff != nil {
			return fmt.Errorf("ActionCommand.Enrich() found FileAdd operation while DirectoryDiff is not nil")
		}
		if c.FileDiff == nil {
			c.FileDiff = &GitDiff{}
		}
		c.FileDiff.Added = append(c.FileDiff.Added, v)
	case FileDelete:
		if c.DirectoryDiff != nil {
			return fmt.Errorf("ActionCommand.Enrich() found FileDelete operation while DirectoryDiff is not nil")
		}
		if c.FileDiff == nil {
			c.FileDiff = &GitDiff{}
		}
		c.FileDiff.Deleted = append(c.FileDiff.Deleted, v)
	case FileUpdate:
		if c.DirectoryDiff != nil {
			return fmt.Errorf("ActionCommand.Enrich() found FileUpdate operation while DirectoryDiff is not nil")
		}
		if c.FileDiff == nil {
			c.FileDiff = &GitDiff{}
		}
		c.FileDiff.Updated = append(c.FileDiff.Updated, v)
	case DirectoryAdd:
		if c.FileDiff != nil {
			return fmt.Errorf("ActionCommand.Enrich() found DirectoryAdd operation while GitDiff is not nil")
		}
		if c.DirectoryDiff == nil {
			c.DirectoryDiff = &DirectoryDiff{}
		}
		c.DirectoryDiff.Added = append(c.DirectoryDiff.Added, v)
	case DirectoryDelete:
		if c.FileDiff != nil {
			return fmt.Errorf("ActionCommand.Enrich() found DirectoryDelete operation while GitDiff is not nil")
		}
		if c.DirectoryDiff == nil {
			c.DirectoryDiff = &DirectoryDiff{}
		}
		c.DirectoryDiff.Deleted = append(c.DirectoryDiff.Deleted, v)
	default:
		return fmt.Errorf("ActionCommand.Enrich() found invalid operation type = %T", op)
	}

	return nil
}

func (c ManualUpdate) MarshalJSON() ([]byte, error) {
	typeName := "ManualUpdate"

	m := make(map[string]*string)
	m["actionType"] = &typeName

	return json.Marshal(m)
}

func (c *ManualUpdate) WriteJsonToFile(filePath string) error {
	bytes, err := json.Marshal(c)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filePath, bytes, 0644); err != nil {
		return err
	}
	return nil
}

func (c *ManualUpdate) Enrich(op FileSystemOperation) error {
	// switch v := op.(type) {
	// case FileAdd:
	// 	if c.DirectoryDiff != nil {
	// 		return fmt.Errorf("ActionCommand.enrich() found FileAdd operation while DirectoryDiff is not nil")
	// 	}
	// 	if c.FileDiff == nil {
	// 		c.FileDiff = &GitDiff{}
	// 	}
	// 	c.FileDiff.Added = append(c.FileDiff.Added, v)
	// case FileDelete:
	// 	if c.DirectoryDiff != nil {
	// 		return fmt.Errorf("ActionCommand.enrich() found FileDelete operation while DirectoryDiff is not nil")
	// 	}
	// 	if c.FileDiff == nil {
	// 		c.FileDiff = &GitDiff{}
	// 	}
	// 	c.FileDiff.Deleted = append(c.FileDiff.Deleted, v)
	// case FileUpdate:
	// 	if c.DirectoryDiff != nil {
	// 		return fmt.Errorf("ActionCommand.enrich() found FileUpdate operation while DirectoryDiff is not nil")
	// 	}
	// 	if c.FileDiff == nil {
	// 		c.FileDiff = &GitDiff{}
	// 	}
	// 	c.FileDiff.Updated = append(c.FileDiff.Updated, v)
	// case DirectoryAdd:
	// 	if c.FileDiff != nil {
	// 		return fmt.Errorf("ActionCommand.enrich() found DirectoryAdd operation while GitDiff is not nil")
	// 	}
	// 	if c.DirectoryDiff == nil {
	// 		c.DirectoryDiff = &DirectoryDiff{}
	// 	}
	// 	c.DirectoryDiff.Added = append(c.DirectoryDiff.Added, v)
	// case DirectoryDelete:
	// 	if c.FileDiff != nil {
	// 		return fmt.Errorf("ActionCommand.enrich() found DirectoryDelete operation while GitDiff is not nil")
	// 	}
	// 	if c.DirectoryDiff == nil {
	// 		c.DirectoryDiff = &DirectoryDiff{}
	// 	}
	// 	c.DirectoryDiff.Deleted = append(c.DirectoryDiff.Deleted, v)
	// default:
	// 	return fmt.Errorf("ActionCommand.enrich() found invalid operation type = %T", op)
	// }

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
		if action.FileDiff != nil && action.DirectoryDiff != nil {
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

func readOperationFromBytes(bytes []byte) (FileSystemOperation, error) {
	typeName, err := extractTypeName(bytes, "operationType")
	if err != nil {
		return nil, fmt.Errorf("readActionFromBytes() failed to extract operationType %s", err)
	}

	switch typeName {
	case "FileAdd":
		var op FileAdd
		if err := json.Unmarshal(bytes, &op); err != nil {
			return nil, err
		}
		return &op, nil
	case "FileUpdate":
		var op FileUpdate
		if err := json.Unmarshal(bytes, &op); err != nil {
			return nil, err
		}
		return &op, nil
	case "FileDelete":
		var op FileDelete
		if err := json.Unmarshal(bytes, &op); err != nil {
			return nil, err
		}
		return &op, nil
	case "DirectoryAdd":
		var op DirectoryAdd
		if err := json.Unmarshal(bytes, &op); err != nil {
			return nil, err
		}
		return &op, nil
	case "DirectoryDelete":
		var op DirectoryDelete
		if err := json.Unmarshal(bytes, &op); err != nil {
			return nil, err
		}
		return &op, nil
	default:
		return nil, fmt.Errorf("readOperationFromBytes() found invalid operationType = %s", typeName)
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
		jsonBytes, err := json.Marshal(jsonObj)
		if err != nil {
			return fmt.Errorf("%s, marshaling JSON failed, %s", errorPreceding, err)
		}

		// unmarshal to action once, to control the re-marshaling behavior
		action, err := readActionFromBytes(jsonBytes)
		if err != nil {
			return fmt.Errorf("%s, reading actoin failed, %s", errorPreceding, err)
		}

		targetFile := fmt.Sprintf("%s/%s%03d.json", targetDir, targetPrefix, i)
		if action.WriteJsonToFile(targetFile) != nil {
			return fmt.Errorf("%s, writing JSON to %s failed, %s", errorPreceding, targetFile, err)
		}
	}

	return nil
}

func EnrichActionFiles(opsListFile, targetDir, targetPrefix string) error {
	errorPreceding := "Error in EnrichActionFiles for filename = " + opsListFile

	// read and process the whole file
	jsonArray, err := readActionList(opsListFile)
	if err != nil {
		return fmt.Errorf("%s, %s", errorPreceding, err)
	}

	// write each array element into file
	for _, jsonObj := range jsonArray {
		opBytes, err := json.Marshal(jsonObj)
		if err != nil {
			return fmt.Errorf("%s, marshaling JSON failed, %s", errorPreceding, err)
		}

		operation, err := readOperationFromBytes(opBytes)
		if err != nil {
			return fmt.Errorf("%s, reading operation failed, %s", errorPreceding, err)
		}

		seqNo, ok := jsonObj["seqNo"]
		if !ok {
			return fmt.Errorf("%s, seqNo not found in JSON, %s", errorPreceding, err)
		}
		actionFile := fmt.Sprintf("%s/%s%03d.json", targetDir, targetPrefix, seqNo)

		action, err := readActionFromFile(actionFile)
		if err != nil {
			return fmt.Errorf("%s, reading action file failed, %s", errorPreceding, err)
		}

		if err := action.Enrich(operation); err != nil {
			return fmt.Errorf("%s, enriching action failed, %s", errorPreceding, err)
		}

		action.WriteJsonToFile(actionFile)
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
	// enrichDir := "data/enriched"
	// if err := EnrichActionFiles("data/source_code_ops.json", inputDir, enrichDir, prefix); err != nil {
	// 	return err
	// }

	// 3.
	files, err := FilesInDir(inputDir, prefix)
	if err != nil {
		return err
	}
	for _, f := range files {
		fmt.Println(f)
	}

	// GenerateInputActionFiles("")

	return nil
}
