package model

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// arbitrary JSON obj representation in Go map
type JsonObj map[string]interface{}

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
	default:
		return nil, fmt.Errorf("readActionFromBytes() found invalid typeName = %s", typeName)
	}
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
			return fmt.Errorf("marshaling flat JSON failed, %s", err)
		}

		targetFile := fmt.Sprintf("%s/%s%03d.json", targetDir, targetPrefix, i)
		if err = os.WriteFile(targetFile, jsonBytes, 0644); err != nil {
			return fmt.Errorf("%s, writing flat JSON to %s failed, %s", errorPreceding, targetFile, err)
		}
	}

	return nil
}

func Processing() error {
	// by-hand csv -> json conversion, and save action-list.json
	if err := SplitActionList("data/action_list.json", "data/input", "action"); err != nil {
		return err
	}

	// files, err := listFilePaths()
	// 	if err != nil {
	// 	return fmt.Errorf("%s, %s", errorPreceding, err)
	// }
	// for i, f := range files {
	// 	converted := convert(f)
	// }

	// GenerateInputActionFiles("")

	return nil
}
