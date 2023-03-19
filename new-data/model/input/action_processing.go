package input

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// map[string]interface{} represents JSON obj
// return a slice of map[string]interface{} (i.e.) []map[string]interface{}
func readActionList(actionListFile string) ([]map[string]interface{}, error) {
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

func SplitActionListFile(actionListFile, targetDir, targetPrefix string) error {
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

func Processin() {
	// by-hand csv -> json conversion, and save action-list.json
	// SplitActionListFile("target", "filename")

	// files := listFilePaths()
	// for i, f := range files {
	// 	converted := convert(f)
	// }

	// GenerateInputActionFiles("")
}
