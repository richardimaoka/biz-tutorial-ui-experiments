package pkg

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Input struct {
	Action  Action
	Results []Result
}

func ReadActionFile(filename string) (*Input, error) {
	errorPreceding := "Error in readActionFile for filename = " + filename

	// read and process the whole file
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPreceding, err)
	}

	var unmarshalled struct {
		Action  map[string]interface{}
		Results []map[string]interface{}
	}
	if err := json.Unmarshal(bytes, &unmarshalled); err != nil {
		return nil, fmt.Errorf("%s, %s", errorPreceding, err)
	}

	// process the action part
	actionBytes, err := json.Marshal(unmarshalled.Action)
	if err != nil {
		return nil, fmt.Errorf("%s, failed to marshal action, %s", errorPreceding, err)
	}

	action, err := GetActionFromBytes(actionBytes)
	if err != nil {
		return nil, fmt.Errorf("%s, failed construct action, %s", errorPreceding, err)
	}

	// process the results part
	var results []Result
	for i, r := range unmarshalled.Results {
		resultBytes, err := json.Marshal(r)
		if err != nil {
			return nil, fmt.Errorf("%s, failed to marshal %s result %s", errorPreceding, ordinal(i), err)
		}
		result, err := GetResultFromBytes(resultBytes)
		if err != nil {
			return nil, fmt.Errorf("%s, failed construct %s result, %s", errorPreceding, ordinal(i), err)
		}
		results = append(results, result)
	}

	return &Input{action, results}, nil
}

type inputUnmarshalled struct {
	Action  map[string]interface{}
	Results []map[string]interface{}
}

func SplitInputFile(inputFilePath string, outputDirPath string) error {
	errorPreceding := "Error in SplitInputFile for filename = " + inputFilePath

	// read and process the input file
	bytes, err := os.ReadFile(inputFilePath)
	if err != nil {
		return fmt.Errorf("%s, %s", errorPreceding, err)
	}

	var unmarshalled struct {
		Action  map[string]interface{}
		Results []map[string]interface{}
	}
	if err := json.Unmarshal(bytes, &unmarshalled); err != nil {
		return fmt.Errorf("%s, %s", errorPreceding, err)
	}

	// process the action part
	actionBytes, err := json.MarshalIndent(unmarshalled.Action, "", "  ")
	if err != nil {
		return fmt.Errorf("%s, failed to marshal action, %s", errorPreceding, err)
	}

	actionFilePath := outputDirPath + "/action.json"
	if err := os.WriteFile(actionFilePath, actionBytes, 0666); err != nil {
		return fmt.Errorf("%s, failed write action to %s, %s", errorPreceding, actionFilePath, err)

	}

	// process the results part
	resultsBytes, err := json.MarshalIndent(unmarshalled.Results, "", "  ")
	if err != nil {
		return fmt.Errorf("%s, failed to marshal results, %s", errorPreceding, err)
	}

	resultsFilePath := outputDirPath + "/results.json"
	if err := os.WriteFile(resultsFilePath, resultsBytes, 0666); err != nil {
		return fmt.Errorf("%s, failed write results to %s, %s", errorPreceding, resultsFilePath, err)
	}

	return nil
}

func ordinal(x int) string {
	suffix := "th"
	switch x % 10 {
	case 1:
		if x%100 != 11 {
			suffix = "st"
		}
	case 2:
		if x%100 != 12 {
			suffix = "nd"
		}
	case 3:
		if x%100 != 13 {
			suffix = "rd"
		}
	}
	return strconv.Itoa(x) + suffix
}
