package pkg

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type Result interface {
	IsResult()
}

type SourceCodeUpdate struct {
	TypeName   string `json:"__typename"`
	FilesAdded []File
}

func (c *SourceCodeUpdate) IsResult() {}

type ChangeCurrentDirectoryInTerminal struct {
	TypeName   string `json:"__typename"`
	TerminalId string
	FilePath   []string
}

func (c *ChangeCurrentDirectoryInTerminal) IsResult() {}

func GetResultFromFile(filename string) (Result, error) {
	errorPreceding := "Error in GetResultFromFile for filename = " + filename

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPreceding, err)
	}

	result, err := GetResultFromBytes(bytes)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPreceding, err)
	}

	return result, nil
}

func GetResultFromBytes(bytes []byte) (Result, error) {
	var unmarshaled interface{}
	if err := json.Unmarshal(bytes, &unmarshaled); err != nil {
		return nil, err
	}

	asserted, ok := unmarshaled.(map[string]interface{}) //type assertion
	if !ok {
		return nil, fmt.Errorf("perhaps the given JSON is not a JSON 'object'")
	}

	typename, ok := asserted["__typename"]
	if !ok {
		return nil, fmt.Errorf("\"__typename\" does not exist in JSON")
	}

	switch t := typename.(type) {
	case string:
		switch t {
		case "SourceCodeUpdate":
			var srcUpdate SourceCodeUpdate
			if err := json.Unmarshal(bytes, &srcUpdate); err != nil {
				return nil, err
			}
			return &srcUpdate, nil

		case "ChangeCurrentDirectoryInTerminal":
			var cd ChangeCurrentDirectoryInTerminal
			if err := json.Unmarshal(bytes, &cd); err != nil {
				return nil, err
			}

			return &cd, nil

		default:
			return nil, fmt.Errorf("\"__typename\" = %s is not a valid action type", t)
		}
	default:
		return nil, fmt.Errorf("\"__typename\" = %v is in wrong type %v", t, reflect.TypeOf(t))
	}
}

func GetResultSliceFromFile(filename string) ([]Result, error) {
	errorPreceding := "Error in GetResultSliceFromFile for filename = " + filename

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPreceding, err)
	}

	resultList, err := GetResultSliceFromBytes(bytes)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPreceding, err)
	}

	return resultList, nil
}

func GetResultSliceFromBytes(bytes []byte) ([]Result, error) {
	var unmarshaled []map[string]interface{}
	if err := json.Unmarshal(bytes, &unmarshaled); err != nil {
		return nil, err
	}

	var results []Result
	for i, r := range unmarshaled {
		resultBytes, err := json.Marshal(r)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal back %s result %s", ordinal(i), err)
		}
		result, err := GetResultFromBytes(resultBytes)
		if err != nil {
			return nil, fmt.Errorf("failed construct %s result, %s", ordinal(i), err)
		}
		results = append(results, result)
	}

	return results, nil
}
