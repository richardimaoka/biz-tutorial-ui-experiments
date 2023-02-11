package model

import (
	"encoding/json"
	"fmt"
	"os"
)

func UnmarshalToAction(jsonBytes []byte) (Action, error) {
	errorPreceding := "Error in UnmarshalToAction"

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
	for i, flat := range unmarshalled {
		flatBytes, err := json.Marshal(flat)
		if err != nil {
			return fmt.Errorf("%s, marshaling %s action failed, %s", errorPreceding, ordinal(i), err)
		}

		var unflat map[string]interface{}
		err = Unflatten(flatBytes, &unflat)
		if err != nil {
			return fmt.Errorf("%s, unflattening %s action failed, %s", errorPreceding, ordinal(i), err)
		}

		marshaledUnflat, err := json.Marshal(unflat)
		if err != nil {
			return fmt.Errorf("%s, marshaling %s *unflattened* action failed, %s", errorPreceding, ordinal(i), err)
		}

		var action ActionCommand
		err = json.Unmarshal(marshaledUnflat, &action)
		if err != nil {
			return fmt.Errorf("%s, unmarshaling %s action to ActionCommand failed, %s", errorPreceding, ordinal(i), err)
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
