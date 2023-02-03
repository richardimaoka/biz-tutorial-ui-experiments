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
	for i, actionFlat := range unmarshalled {
		marshaledBack, err := json.Marshal(actionFlat)
		if err != nil {
			return fmt.Errorf("%s, marshaling %s action failed, %s", errorPreceding, ordinal(i), err)
		}

		var action map[string]interface{}
		err = Unflatten(marshaledBack, &action)
		if err != nil {
			return fmt.Errorf("%s, unflattening %s action failed, %s", errorPreceding, ordinal(i), err)
		}

		marshaledUnflat, err := json.MarshalIndent(action, "", "  ")
		if err != nil {
			return fmt.Errorf("%s, marshaling %s *unflattened* action failed, %s", errorPreceding, ordinal(i), err)
		}

		err = os.WriteFile(fmt.Sprintf("%s/%s%03d.json", targetDir, targetFilePrefix, i), marshaledUnflat, 0644)
		if err != nil {
			return fmt.Errorf("%s, writing %s action failed, %s", errorPreceding, ordinal(i), err)
		}
	}

	return nil
}
