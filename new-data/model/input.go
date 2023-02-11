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

func toUnflatJsonBytes(flatJsonObj map[string]interface{}) ([]byte, error) {
	errorPreceding := "Error in reflat"

	flatJsonBytes, err := json.Marshal(flatJsonObj)
	if err != nil {
		return nil, fmt.Errorf("%s, marshaling to flat JSON bytes failed, %s", errorPreceding, err)
	}

	var unflatJsonObj map[string]interface{}
	err = Unflatten(flatJsonBytes, &unflatJsonObj)
	if err != nil {
		return nil, fmt.Errorf("%s, unflattening failed, %s", errorPreceding, err)
	}

	unflatJsonBytes, err := json.Marshal(unflatJsonObj)
	if err != nil {
		return nil, fmt.Errorf("%s, marshaling to unflat JSON bytes failed, %s", errorPreceding, err)
	}

	return unflatJsonBytes, nil
}

func toActionJsonBytes(flatJsonBytes []byte) ([]byte, error) {
	unflattenedJsonBytes, err := UnflattenBytes(flatJsonBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to unflatten, %s", err)
	}

	action, err := UnmarshalToAction(unflattenedJsonBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal, %s", err)
	}

	actionJsonBytes, err := json.MarshalIndent(action, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("marshaling action failed, %s", err)
	}

	return actionJsonBytes, nil
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

func AAA(targetDir string) error {
	list, err := ListInputFiles(targetDir)
	if err != nil {
		return err
	}

	// process each element
	for i, inputFileName := range list {
		flatJsonBytes, err := os.ReadFile(inputFileName)
		if err != nil {
			return fmt.Errorf("failed to read %s, %s", inputFileName, err)
		}

		actionJsonBytes, err := toActionJsonBytes(flatJsonBytes)
		if err != nil {
			return fmt.Errorf("failed to convert %s to action json, %s", inputFileName, err)
		}

		actionFileName := fmt.Sprintf("%s/%s%03d.json", targetDir, actionFilePrefix, i)
		err = os.WriteFile(inputFileName, actionJsonBytes, 0644)
		if err != nil {
			return fmt.Errorf("failed to write %s, %s", actionFileName, err)
		}
	}

	return nil
}

func ListInputFiles(targetDir string) ([]string, error) {
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

func SplitActionListFile2(targetDir string) error {
	actionListFile := targetDir + "/" + actionListPrefix + ".json"
	errorPreceding := "Error in SplitActionListFile for filename = " + actionListFile

	// read and process the whole file
	bytes, err := os.ReadFile(actionListFile)
	if err != nil {
		return fmt.Errorf("%s, reading action list failed, %s", errorPreceding, err)
	}

	var unmarshalled []map[string]interface{}
	if err := json.Unmarshal(bytes, &unmarshalled); err != nil {
		return fmt.Errorf("%s, unmarshaling action failed, %s", errorPreceding, err)
	}

	// process each element
	for i, flatJsonObj := range unmarshalled {
		jsonBytes, err := toUnflatJsonBytes(flatJsonObj)
		if err != nil {
			return fmt.Errorf("%s, %s action, %s", errorPreceding, ordinal(i), err)
		}

		action, err := UnmarshalToAction(jsonBytes)
		if err != nil {
			return fmt.Errorf("%s, %s action, %s", errorPreceding, ordinal(i), err)
		}

		outBytes, err := json.MarshalIndent(action, "", "  ")
		if err != nil {
			return fmt.Errorf("%s, marshaling %s ActionCommand failed, %s", errorPreceding, ordinal(i), err)
		}

		inputFileName := fmt.Sprintf("%s/%s%03d.json", targetDir, inputFilePrefix, i)
		err = os.WriteFile(inputFileName, outBytes, 0644)
		if err != nil {
			return fmt.Errorf("%s, writing %s action failed, %s", errorPreceding, ordinal(i), err)
		}
	}

	return nil
}
