package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Action interface {
	IsAction()
}

type Command struct {
	TypeName string `json:"__typename"`
	Command  string
}

func (c *Command) IsAction() {}

type File struct {
	TypeName    string `json:"__typename"`
	FilePath    []string
	FileContent string
	Offset      int
}

type Terminal struct {
	elements []interface{}
}

type State struct {
	SourceCode interface{}
	Terminal   interface{}
}

func getActionFromFile(filename string) (Action, error) {
	errorPreceding := "Error in getActionFromFile for filename = " + filename

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPreceding, err)
	}

	action, err := getActionFromBytes(bytes)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPreceding, err)
	}

	return action, nil
}

func getActionFromBytes(bytes []byte) (Action, error) {
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
		case "Command":
			var command Command
			if err := json.Unmarshal(bytes, &command); err != nil {
				return nil, err
			}

			return &command, nil
		default:
			return nil, fmt.Errorf("\"__typename\" = %s is not a valid action type", t)
		}
	default:
		return nil, fmt.Errorf("\"__typename\" = %v is in wrong type %v", t, reflect.TypeOf(t))
	}
}

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

func getResultFromFile(filename string) (Result, error) {
	errorPreceding := "Error in getResultFromFile for filename = " + filename

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPreceding, err)
	}

	result, err := getResultFromBytes(bytes)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPreceding, err)
	}

	return result, nil
}

func getResultFromBytes(bytes []byte) (Result, error) {
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

type ActionInfo struct {
	Action  Action
	Results []Result
}

func readActionFile(filename string) (*ActionInfo, error) {
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

	action, err := getActionFromBytes(actionBytes)
	if err != nil {
		return nil, fmt.Errorf("%s, failed construct action, %s", errorPreceding, err)
	}

	// process the results part
	var results []Result
	for i, r := range unmarshalled.Results {
		resultBytes, err := json.Marshal(r)
		if err != nil {
			return nil, fmt.Errorf("%s, failed to marshal %s result %s", errorPreceding, Ordinal(i), err)
		}
		result, err := getResultFromBytes(resultBytes)
		if err != nil {
			return nil, fmt.Errorf("%s, failed construct %s result, %s", errorPreceding, Ordinal(i), err)
		}
		results = append(results, result)
	}

	return &ActionInfo{action, results}, nil
}

func Ordinal(x int) string {
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

func main() {
	filename := "step01/action.json"
	action, err := getActionFromFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(action)

	a, err := readActionFile("action01.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
	// filename = "step01/result.json"
	// result, err := getResult(filename)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(result)
}
