package model

import (
	"encoding/json"
	"fmt"
	"os"
)

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
