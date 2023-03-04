package model

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const (
	inputFilePrefix  = "input_flat"
	actionFilePrefix = "action"
	actionListPrefix = "action_input_list"
)

func toActionJsonBytes(flatJsonBytes []byte) ([]byte, error) {
	// pre-process JSON bytes
	unflattenedJsonBytes, err := ToUnflattenBytes(flatJsonBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to unflatten, %s", err)
	}

	// unmarshal to action
	action, err := unmarshalToAction(unflattenedJsonBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal, %s", err)
	}

	// marshal action
	actionJsonBytes, err := json.MarshalIndent(action, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("marshaling action failed, %s", err)
	}

	return actionJsonBytes, nil
}

func generateActionFile(flatJsonFileName, actionFileName string) error {
	// read
	flatJsonBytes, err := os.ReadFile(flatJsonFileName)
	if err != nil {
		return fmt.Errorf("failed to read %s, %s", flatJsonFileName, err)
	}

	// convert
	actionJsonBytes, err := toActionJsonBytes(flatJsonBytes)
	if err != nil {
		return fmt.Errorf("failed to convert %s to action json, %s", flatJsonFileName, err)
	}

	// write
	err = os.WriteFile(actionFileName, actionJsonBytes, 0644)
	if err != nil {
		return fmt.Errorf("failed to write %s, %s", actionFileName, err)
	}

	return nil
}

// map[string]interface{} represents JSON obj
// return a slice of map[string]interface{} (i.e.) []map[string]interface{}
func reacActionList(actionListFile string) ([]map[string]interface{}, error) {
	bytes, err := os.ReadFile(actionListFile)
	if err != nil {
		return nil, fmt.Errorf("reading %s failed, %s", actionListFile, err)
	}

	var unmarshalled []map[string]interface{}
	if err := json.Unmarshal(bytes, &unmarshalled); err != nil {
		return nil, fmt.Errorf("unmarshaling %s failed, %s", actionListFile, err)
	}

	return unmarshalled, nil
}

func writeFlatJson(flatJson map[string]interface{}, filename string) error {
	jsonBytes, err := json.MarshalIndent(flatJson, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling flat JSON failed, %s", err)
	}

	err = os.WriteFile(filename, jsonBytes, 0644)
	if err != nil {
		return fmt.Errorf("writing flat JSON to %s failed, %s", filename, err)
	}

	return nil
}

func SplitActionListFile(targetDir string) error {
	actionListFile := fmt.Sprintf("%s/%s.json", targetDir, actionListPrefix)
	errorPreceding := "Error in SplitInputListFile for filename = " + actionListFile

	// read and process the whole file
	jsonObjMaps, err := reacActionList(actionListFile)
	if err != nil {
		return fmt.Errorf("%s, %s", errorPreceding, err)
	}

	// write each array element into file
	var errorHappened bool = false
	for i, flatJsonObj := range jsonObjMaps {
		inputFileName := fmt.Sprintf("%s/%s%03d.json", targetDir, inputFilePrefix, i)
		err = writeFlatJson(flatJsonObj, inputFileName)
		if err != nil {
			fmt.Printf("GenerateInputActionFiles failed, %s", err)
			errorHappened = true
		}
	}

	if errorHappened {
		return fmt.Errorf("error happend while splitting %s", actionListFile)
	}

	return nil
}

// all input_flat00x files
func ListInputFlatFiles(targetDir string) ([]string, error) {
	entries, err := os.ReadDir(targetDir)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, e := range entries {
		if strings.HasPrefix(e.Name(), "input") && strings.HasSuffix(e.Name(), "json") {
			files = append(files, targetDir+"/"+e.Name())
		}
	}

	return files, nil
}

// generate action00x.json files from input_flat00x.json files
func GenerateInputActionFiles(targetDir string) error {
	inputFlatFiles, err := ListInputFlatFiles(targetDir)
	if err != nil {
		return err
	}

	var errorHappened bool = false
	for i, flatJsonFileName := range inputFlatFiles {
		actionFileName := fmt.Sprintf("%s/%s%03d.json", targetDir, actionFilePrefix, i)
		err = generateActionFile(flatJsonFileName, actionFileName)
		if err != nil {
			fmt.Printf("GenerateInputActionFiles failed, %s", err)
			errorHappened = true
		}
	}

	if errorHappened {
		return fmt.Errorf("error happend while processing action files")
	}

	return nil
}

func unmarshalToAction(jsonBytes []byte) (Action, error) {
	errorPreceding := "Error in unmarshalToAction"

	actionTypeField := "actionType"
	actionType, err := extractTypeName(jsonBytes, actionTypeField)
	if err != nil {
		return nil, fmt.Errorf("%s, extracting action type failed, %s", errorPreceding, err)
	}

	switch actionType {
	case "ActionCommand":
		var command ActionCommand
		err = json.Unmarshal(jsonBytes, &command)
		if err != nil {
			return nil, fmt.Errorf("%s, unmarshaling action to ActionCommand failed, %s", errorPreceding, err)
		}
		return &command, nil
	case "ManualUpdate":
		var manual ManualUpdate
		err = json.Unmarshal(jsonBytes, &manual)
		if err != nil {
			return nil, fmt.Errorf("%s, unmarshaling action to ManualUpdate failed, %s", errorPreceding, err)
		}
		return &manual, nil
	default:
		return nil, fmt.Errorf("%s, %s = %s is not a valid action type", errorPreceding, actionTypeField, actionType)
	}
}

// func unmarshalToAction2(jsonBytes []byte) (*ActionTerminal, error) {
// 	errorPreceding := "Error in unmarshalTo"

// 	actionTypeField := "actionType"
// 	actionType, err := extractTypeName(jsonBytes, actionTypeField)
// 	if err != nil {
// 		return nil, fmt.Errorf("%s, extracting action type failed, %s", errorPreceding, err)
// 	}

// 	switch actionType {
// 	case "ActionTerminal":
// 		var action ActionTerminal
// 		err = json.Unmarshal(jsonBytes, &action)
// 		if err != nil {
// 			return nil, fmt.Errorf("%s, unmarshaling action to ActionTerminal failed, %s", errorPreceding, err)
// 		}
// 		return &action, nil
// 	default:
// 		return nil, fmt.Errorf("%s, %s = %s is not a valid action type", errorPreceding, actionTypeField, actionType)
// 	}
// }
