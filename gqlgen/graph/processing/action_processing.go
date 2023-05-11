package processing

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

func ReadActionFromFiles(actionDir, actionPrefix string) ([]Action, error) {
	var actions []Action

	actionFiles, err := FilesInDir(actionDir, actionPrefix)
	if err != nil {
		return nil, err
	}

	for i, file := range actionFiles {
		expectedFileName := targetFileName(actionDir, actionPrefix, i)
		if expectedFileName != file {
			return nil, fmt.Errorf("expected file %s, got %s", expectedFileName, file)
		}
		action, err := readAction(file)
		if err != nil {
			return nil, fmt.Errorf("reading action file failed, %s", err)
		}
		actions = append(actions, action)
	}

	return actions, nil
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

	// 1. read and process the whole file
	jsonArray, err := readJsonArray(actionListFile)
	if err != nil {
		return fmt.Errorf("%s, %s", errorPreceding, err)
	}
	log.Printf("SplitActionList: read %d actions from %s", len(jsonArray), actionListFile)

	// 2. write each array element into file
	for i, jsonObj := range jsonArray {
		jsonBytes, err := json.Marshal(jsonObj)
		if err != nil {
			return fmt.Errorf("%s, marshaling JSON failed, %s", errorPreceding, err)
		}

		// unmarshal to action once, to control the re-marshaling behavior
		action, err := unmarshalAction(jsonBytes)
		if err != nil {
			return fmt.Errorf("%s, reading action failed, %s", errorPreceding, err)
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

	// 1. load actions into memory
	actions, err := ReadActionFromFiles(actionDir, actionPrefix)
	if err != nil {
		return fmt.Errorf("%s, %s", errorPreceding, err)
	}
	log.Printf("EnrichActionFiles: read %d actions from %s", len(actions), actionDir)

	// 2. read operations
	jsonOpsArray, err := readJsonArray(opsListFile)
	if err != nil {
		return fmt.Errorf("%s, %s", errorPreceding, err)
	}
	log.Printf("EnrichActionFiles: read %d operations from %s", len(jsonOpsArray), opsListFile)

	// 3. enrich actions
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

		actions[seqNoInt].Enrich(operation)
	}

	// 4. write enriched actions to files
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

	// 1. load actions into memory
	actions, err := ReadActionFromFiles(actionDir, actionPrefix)
	if err != nil {
		return fmt.Errorf("%s, %s", errorPreceding, err)
	}
	log.Printf("%s: read %d actions from %s", errorPreceding, len(actions), actionDir)

	// 2.   apply actions
	// 2.1. initial action
	pageState, err := InitPageStateProcessor(actions[0])
	if err != nil {
		return fmt.Errorf("%s, %s", errorPreceding, err)
	}
	fileName := stateFileName(targetDir, targetPrefix, pageState.step.currentStep)
	if err := WriteJsonToFile(pageState.ToGraphQLPageState(), fileName); err != nil {
		return fmt.Errorf("%s, writing JSON to %s failed, %s", errorPreceding, fileName, err)
	}

	// 2.2. actions up to last - 1
	for i := 1; i < len(actions)-1; i++ {
		if err := pageState.StateTransition(actions[i]); err != nil {
			return fmt.Errorf("%s, %s", errorPreceding, err)
		}
		fileName := stateFileName(targetDir, targetPrefix, pageState.step.currentStep)
		if err := WriteJsonToFile(pageState.ToGraphQLPageState(), fileName); err != nil {
			return fmt.Errorf("%s, writing JSON to %s failed, %s", errorPreceding, fileName, err)
		}
	}

	// 2.3. last action
	pageState.LastTransition()
	fileName = stateFileName(targetDir, targetPrefix, pageState.step.currentStep)
	if err := WriteJsonToFile(pageState.ToGraphQLPageState(), fileName); err != nil {
		return fmt.Errorf("%s, writing JSON to %s failed, %s", errorPreceding, fileName, err)
	}

	return nil
}

func Processing() error {
	actionPrefix := "action"

	// 1. split action-list.json
	inputDir := "data/input"
	if err := os.MkdirAll(inputDir, 0755); err != nil {
		return err
	}
	if err := SplitActionList("data/action_list.json", inputDir, actionPrefix); err != nil {
		return err
	}

	// 2. enrich action files
	enrichedDir := "data/enriched"
	if err := os.MkdirAll(enrichedDir, 0755); err != nil {
		return err
	}
	if err := EnrichActionFiles("data/source_code_ops.json", inputDir, enrichedDir, actionPrefix); err != nil {
		return err
	}

	// 3. apply action files
	stateDir := "data/state"
	if err := os.MkdirAll(stateDir, 0755); err != nil {
		return err
	}
	if err := ApplyActions(enrichedDir, actionPrefix, stateDir, "state"); err != nil {
		return err
	}

	return nil
}
