package model2

import (
	"encoding/json"
	"fmt"
	"os"
)

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
