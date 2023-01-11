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

func getAction(filename string) (Action, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error in filename = %s, %s", filename, err)
	}

	var unmarshaled interface{}
	if err := json.Unmarshal(bytes, &unmarshaled); err != nil {
		return nil, fmt.Errorf("Error in filename = %s, while unmarshaling JSON from file, %s", filename, err)
	}

	asserted, ok := unmarshaled.(map[string]interface{}) //type assertion
	if !ok {
		return nil, fmt.Errorf("Error in filename = %s, while constructing Go data from JSON, perhaps the file is not in JSON 'object'", filename)
	}

	typename, ok := asserted["__typename"]
	if !ok {
		return nil, fmt.Errorf("Error in filename = %s, while validating action type: \"__typename\" does not exist", filename)
	}

	switch t := typename.(type) {
	case string:
		switch t {
		case "Command":
			var command Command
			if err := json.Unmarshal(bytes, &command); err != nil {
				return nil, fmt.Errorf("Error in filename = %s, while unmarshaling JSON from file, %s", filename, err)
			}

			return &command, nil
		default:
			return nil, fmt.Errorf("Error in filename = %s, while validating action type: %s is not a valid action type", filename, t)
		}
	default:
		return nil, fmt.Errorf("Error in filename = %s, while validating action type: \"__typename\" = %v is in wrong type %v", filename, t, reflect.TypeOf(t))
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

func getResult(filename string) (Result, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error in filename = %s, %s", filename, err)
	}

	var unmarshaled interface{}
	if err := json.Unmarshal(bytes, &unmarshaled); err != nil {
		return nil, fmt.Errorf("Error in filename = %s, while unmarshaling JSON from file, %s", filename, err)
	}

	asserted, ok := unmarshaled.(map[string]interface{}) //type assertion
	if !ok {
		return nil, fmt.Errorf("Error in filename = %s, while constructing Go data from JSON, perhaps the file is not in JSON 'object'", filename)
	}

	typename, ok := asserted["__typename"]
	if !ok {
		return nil, fmt.Errorf("Error in filename = %s, while validating action type: \"__typename\" does not exist", filename)
	}

	switch t := typename.(type) {
	case string:
		switch t {
		case "SourceCodeUpdate":
			var srcUpdate SourceCodeUpdate
			if err := json.Unmarshal(bytes, &srcUpdate); err != nil {
				return nil, fmt.Errorf("Error in filename = %s, while unmarshaling JSON from file, %s", filename, err)
			}

			return &srcUpdate, nil

		case "ChangeCurrentDirectoryInTerminal":
			var cd ChangeCurrentDirectoryInTerminal
			if err := json.Unmarshal(bytes, &cd); err != nil {
				return nil, fmt.Errorf("Error in filename = %s, while unmarshaling JSON from file, %s", filename, err)
			}

			return &cd, nil

		default:
			return nil, fmt.Errorf("Error in filename = %s, while validating action type: %s is not a valid action type", filename, t)
		}

	default:
		return nil, fmt.Errorf("Error in filename = %s, while validating action type: \"__typename\" = %v is in wrong type %v", filename, t, reflect.TypeOf(t))
	}
}

// func getEachResult(bytes []byte, filename string) (Result, error) {
// 	var unmarshaled interface{}
// 	if err := json.Unmarshal(bytes, &unmarshaled); err != nil {
// 		return nil, fmt.Errorf("Error in filename = %s, while unmarshaling JSON from file, %s", filename, err)
// 	}

// 	asserted, ok := unmarshaled.(map[string]interface{}) //type assertion
// 	if !ok {
// 		return nil, fmt.Errorf("Error in filename = %s, while constructing Go data from JSON, perhaps the file is not in JSON 'object'", filename)
// 	}

// 	typename, ok := asserted["__typename"]
// 	if !ok {
// 		return nil, fmt.Errorf("Error in filename = %s, while validating action type: \"__typename\" does not exist", filename)
// 	}

// 	switch t := typename.(type) {
// 	case string:
// 		switch t {
// 		case "SourceCodeUpdate":
// 			var sourceCodeUpdate SourceCodeUpdate
// 			if err := json.Unmarshal(bytes, &sourceCodeUpdate); err != nil {
// 				return nil, fmt.Errorf("Error in filename = %s, while unmarshaling JSON from file, %s", filename, err)
// 			}

// 			return &sourceCodeUpdate, nil
// 		default:
// 			return nil, fmt.Errorf("Error in filename = %s, while validating result type: %s is not a valid result type", filename, t)
// 		}
// 	default:
// 		return nil, fmt.Errorf("Error in filename = %s, while validating result type: \"__typename\" = %v is in wrong type %v", filename, t, reflect.TypeOf(t))
// 	}
// }

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

// func getResult(filename string) ([]Result, error) {
// 	bytes, err := os.ReadFile(filename)
// 	if err != nil {
// 		return nil, fmt.Errorf("Error in filename = %s, %s", filename, err)
// 	}

// 	var data interface{}
// 	if err := json.Unmarshal(bytes, &data); err != nil {
// 		return nil, fmt.Errorf("Error in filename = %s, while unmarshaling JSON from file, %s", filename, err)
// 	}

// 	asserted, ok := data.([]map[string]string) //type assertion
// 	if !ok {
// 		fmt.Println(reflect.TypeOf(data))
// 		return nil, fmt.Errorf("Error in getResult, filename = %s, while constructing Go data from JSON, perhaps the file is not in JSON 'array of object'", filename)
// 	}

// 	var results []Result
// 	for i, v := range asserted {
// 		elementBytes, err := json.Marshal(v)
// 		if err != nil {
// 			ordinal := Ordinal(i)
// 			return nil, fmt.Errorf("Error in filename = %s, while marshaling JSON array %s element, %s", filename, ordinal, err)
// 		}

// 		eachResult, err := getEachResult(elementBytes, filename)
// 		if err != nil {
// 			ordinal := Ordinal(i)
// 			return nil, fmt.Errorf("Error in filename = %s, while constructing Go data from JSON array %s element, %s", filename, ordinal, err)
// 		}
// 		results = append(results, eachResult)
// 	}

// 	return results, nil
// }

func main() {
	filename := "step01/action.json"
	action, err := getAction(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(action)

	// filename = "step01/result.json"
	// result, err := getResult(filename)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(result)
}
