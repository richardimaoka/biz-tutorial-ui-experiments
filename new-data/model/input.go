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

func toUnflattenBytes(flatJsonBytes []byte) ([]byte, error) {
	var unflatJsonObj map[string]interface{}
	err := Unflatten(flatJsonBytes, &unflatJsonObj)
	if err != nil {
		return nil, fmt.Errorf("unflattening failed, %s", err)
	}

	unflatJsonBytes, err := json.Marshal(unflatJsonObj)
	if err != nil {
		return nil, fmt.Errorf("marshaling to unflattened JSON bytes failed, %s", err)
	}

	return unflatJsonBytes, nil
}

func toActionJsonBytes(flatJsonBytes []byte) ([]byte, error) {
	// pre-process JSON bytes
	unflattenedJsonBytes, err := toUnflattenBytes(flatJsonBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to unflatten, %s", err)
	}

	// unmarshal to action
	action, err := UnmarshalToAction(unflattenedJsonBytes)
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
}

func SplitActionListFile(targetDir string) error {
	actionListFile := fmt.Sprintf("%s/%s.json", targetDir, actionListPrefix)
	errorPreceding := "Error in SplitInputListFile for filename = " + actionListFile

	// read and process the whole file
	bytes, err := os.ReadFile(actionListFile)
	if err != nil {
		return fmt.Errorf("%s, reading action failed, %s", errorPreceding, err)
	}

	var unmarshalled []map[string]interface{}
	if err := json.Unmarshal(bytes, &unmarshalled); err != nil {
		return fmt.Errorf("%s, unmarshaling action failed, %s", errorPreceding, err)
	}

	// write each array element into file
	for i, flatJsonObj := range unmarshalled {
		jsonBytes, err := json.MarshalIndent(flatJsonObj, "", "  ")
		if err != nil {
			return fmt.Errorf("%s, marshaling %s ActionCommand failed, %s", errorPreceding, ordinal(i), err)
		}

		inputFileName := fmt.Sprintf("%s/%s%03d.json", targetDir, inputFilePrefix, i)
		err = os.WriteFile(inputFileName, jsonBytes, 0644)
		if err != nil {
			return fmt.Errorf("%s, writing %s action failed, %s", errorPreceding, ordinal(i), err)
		}
	}

	return nil
}

func GenerateInputActionFiles(targetDir string) error {
	inputFlatFiles, err := ListInputFlatFiles(targetDir)
	if err != nil {
		return err
	}

	var errorHappened bool = false
	for i, flatJsonFileName := range inputFlatFiles {
		actionFileName := fmt.Sprintf("%s/%s%03d.json", targetDir, actionFilePrefix, i)
		errorHappened = errorHappened && generateActionFile(flatJsonFileName, actionFileName) != nil
	}

	if errorHappened {
		return fmt.Errorf("error happend while processing action files")
	}
	return nil
}

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
