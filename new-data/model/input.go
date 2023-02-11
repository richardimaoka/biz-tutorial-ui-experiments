package model

import (
	"encoding/json"
	"fmt"
	"os"
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

func SplitInputListFile(targetDir string) error {
	actionListFile := targetDir + "action_input_list.json"
	targetFilePrefix := "flat_input"

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

		err = os.WriteFile(fmt.Sprintf("%s/%s%03d.json", targetDir, targetFilePrefix, i), jsonBytes, 0644)
		if err != nil {
			return fmt.Errorf("%s, writing %s action failed, %s", errorPreceding, ordinal(i), err)
		}
	}

	return nil
}

func SplitActionListFile(actionListFile, targetDir, targetFilePrefix string) error {
	errorPreceding := "Error in SplitActionListFile for filename = " + actionListFile

	// read and process the whole file
	bytes, err := os.ReadFile(actionListFile)
	if err != nil {
		return fmt.Errorf("%s, reading action failed, %s", errorPreceding, err)
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

		err = os.WriteFile(fmt.Sprintf("%s/%s%03d.json", targetDir, targetFilePrefix, i), outBytes, 0644)
		if err != nil {
			return fmt.Errorf("%s, writing %s action failed, %s", errorPreceding, ordinal(i), err)
		}
	}

	return nil
}
