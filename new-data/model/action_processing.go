package model

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// arbitrary JSON obj representation in Go map
type JsonObj map[string]interface{}

func readJsonArray(filename string) ([]JsonObj, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("reading %s failed, %s", filename, err)
	}

	var unmarshalled []JsonObj
	if err := json.Unmarshal(bytes, &unmarshalled); err != nil {
		return nil, fmt.Errorf("unmarshaling %s failed, %s", filename, err)
	}

	return unmarshalled, nil
}

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

func targetFileName(targetDir, targetPrefix string, index int) string {
	return fmt.Sprintf("%s/%s%03d.json", targetDir, targetPrefix, index)
}

func stateFileName(targetDir, targetPrefix, step string) string {
	return fmt.Sprintf("%s/%s-%s.json", targetDir, targetPrefix, step)
}

func WriteJsonToFile(v any, filePath string) error {
	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filePath, bytes, 0644); err != nil {
		return err
	}
	return nil
}

func SplitActionList(actionListFile, targetDir, targetPrefix string) error {
	errorPreceding := "Error in SplitInputListFile for filename = " + actionListFile

	// read and process the whole file
	jsonArray, err := readJsonArray(actionListFile)
	if err != nil {
		return fmt.Errorf("%s, %s", errorPreceding, err)
	}
	log.Printf("SplitActionList: read %d actions from %s", len(jsonArray), actionListFile)

	// write each array element into file
	for i, jsonObj := range jsonArray {
		jsonBytes, err := json.Marshal(jsonObj)
		if err != nil {
			return fmt.Errorf("%s, marshaling JSON failed, %s", errorPreceding, err)
		}

		// unmarshal to action once, to control the re-marshaling behavior
		action, err := unmarshalAction(jsonBytes)
		if err != nil {
			return fmt.Errorf("%s, reading actoin failed, %s", errorPreceding, err)
		}

		targetFile := targetFileName(targetDir, targetPrefix, i)
		if WriteJsonToFile(action, targetFile) != nil {
			return fmt.Errorf("%s, writing JSON to %s failed, %s", errorPreceding, targetFile, err)
		}
	}
	log.Printf("SplitActionList: wrote %d actions to %s", len(jsonArray), targetDir)

	return nil
}

func EnrichActionFiles(opsListFile, actionDir, targetDir, actionPrefix string) error {
	errorPreceding := "Error in EnrichActionFiles"

	// load actions into memory
	var actions []Action

	actionFiles, err := FilesInDir(actionDir, actionPrefix)
	if err != nil {
		return fmt.Errorf("%s, %s", errorPreceding, err)
	}

	for i, file := range actionFiles {
		expectedFileName := targetFileName(actionDir, actionPrefix, i)
		if expectedFileName != file {
			return fmt.Errorf("%s, expected file %s, got %s", errorPreceding, expectedFileName, file)
		}
		action, err := readAction(file)
		if err != nil {
			return fmt.Errorf("%s, reading action file failed, %s", errorPreceding, err)
		}
		actions = append(actions, action)
	}
	log.Printf("EnrichActionFiles: read %d actions from %s", len(actions), actionDir)

	// read operations and enrich actions
	jsonOpsArray, err := readJsonArray(opsListFile)
	if err != nil {
		return fmt.Errorf("%s, %s", errorPreceding, err)
	}
	log.Printf("EnrichActionFiles: read %d operations from %s", len(jsonOpsArray), opsListFile)

	for _, jsonObj := range jsonOpsArray {
		opBytes, err := json.Marshal(jsonObj)
		if err != nil {
			return fmt.Errorf("%s, marshaling operation JSON failed, %s", errorPreceding, err)
		}
		operation, err := unmarshalFileSystemOperation(opBytes)
		if err != nil {
			return fmt.Errorf("%s, reading operation failed, %s", errorPreceding, err)
		}

		seqNo, ok := jsonObj["seqNo"]
		if !ok {
			return fmt.Errorf("%s, seqNo not found in JSON = %s, %s", errorPreceding, opBytes, err)
		}
		seqNoFloat, ok := seqNo.(float64)
		if !ok {
			return fmt.Errorf("%s, seqNo not number in JSON = %s, %s", errorPreceding, opBytes, err)
		}
		seqNoInt := int(seqNoFloat)
		if len(actions) <= seqNoInt {
			return fmt.Errorf("%s, seqNo = %d is out of range, %s", errorPreceding, seqNoInt, err)
		}

		if enriched, err := actions[seqNoInt].Enrich(operation); err != nil {
			return fmt.Errorf("%s, enriching action %d failed, %s", errorPreceding, seqNoInt, err)
		} else {
			actions[seqNoInt] = enriched
		}
	}

	// write enriched actions to files
	for i, action := range actions {
		targetFile := targetFileName(targetDir, actionPrefix, i)
		if err := WriteJsonToFile(action, targetFile); err != nil {
			return fmt.Errorf("%s, writing JSON to %s failed, %s", errorPreceding, targetFile, err)
		}
	}
	log.Printf("EnrichActionFiles: wrote %d actions to %s", len(actions), targetDir)

	return nil
}

func ApplyActions(actionDir, actionPrefix, targetDir, targetPrefix string) error {
	errorPreceding := "Error in ApplyActions"

	actionFiles, err := FilesInDir(actionDir, actionPrefix)
	if err != nil {
		return fmt.Errorf("%s, %s", errorPreceding, err)
	}
	log.Printf("ApplyActions: read %d actions from %s", len(actionFiles), actionDir)

	pageState := NewPageState()
	fileName := stateFileName(targetDir, targetPrefix, *pageState.Step)
	if err := WriteJsonToFile(pageState, fileName); err != nil {
		return fmt.Errorf("%s, writing JSON to %s failed, %s", errorPreceding, fileName, err)
	}

	for _, file := range actionFiles {
		action, err := readAction(file)
		if err != nil {
			return fmt.Errorf("%s, reading action file failed, %s", errorPreceding, err)
		}

		switch v := action.(type) {
		case ActionCommand:
			if err := pageState.TypeInCommand(v); err != nil {
				return fmt.Errorf("%s, %s", errorPreceding, err)
			}
			fileName = stateFileName(targetDir, targetPrefix, *pageState.Step)
			if err := WriteJsonToFile(pageState, fileName); err != nil {
				return fmt.Errorf("%s, writing JSON to %s failed, %s", errorPreceding, fileName, err)
			}
			// log.Printf("ApplyActions: typed in command  %d from %s, step = %s", i, file, *pageState.Step)

			if err := pageState.ExecuteLastCommand(v); err != nil {
				return fmt.Errorf("%s, %s", errorPreceding, err)
			}
			fileName = stateFileName(targetDir, targetPrefix, *pageState.Step)
			if err := WriteJsonToFile(pageState, fileName); err != nil {
				return fmt.Errorf("%s, writing JSON to %s failed, %s", errorPreceding, fileName, err)
			}
			// log.Printf("ApplyActions: executed command  %d from %s, step = %s", i, file, *pageState.Step)
		case ManualUpdate:
		}
	}
	log.Printf("ApplyActions: wrote %s states to %s", *pageState.Step, targetDir)

	return nil
}

func Processing() error {
	// 0. prereuisite: by-hand csv -> json conversion, and save action-list.json

	// 1. split action-list.json
	inputDir := "data/input"
	prefix := "action"
	if err := SplitActionList("data/action_list.json", inputDir, prefix); err != nil {
		return err
	}

	// 2. enrich action files
	enrichedDir := "data/enriched"
	if err := EnrichActionFiles("data/source_code_ops.json", inputDir, enrichedDir, prefix); err != nil {
		return err
	}

	// 3. apply action files
	if err := ApplyActions(enrichedDir, prefix, "data/state", "state"); err != nil {
		return err
	}

	return nil
}
