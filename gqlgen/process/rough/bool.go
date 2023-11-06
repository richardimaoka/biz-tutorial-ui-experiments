package rough

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/jsonwrap"
)

func toBool(s string) bool {
	switch s {
	case "TRUE":
		return true
	case "true":
		return true
	case "FALSE":
		return false
	case "false":
		return false
	case "": //mapping string zero value = "" to bool zero value = false
		return false
	default:
		return false
	}
}

func toInt(s string) int {
	switch s {
	case "": //mapping string zero value = "" to bool zero value = false
		return 0
	default:
		return 0
	}
}

func alterStringToBool(jsonObj jsonwrap.JsonObj, field reflect.StructField) error {
	// for those struct fields where type = bool, convert string to bool
	if field.Type.String() == "bool" {
		// validate st
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			return fmt.Errorf("json tag not found for field %s", field.Name)
		}

		jsonValue, ok := jsonObj[jsonTag]
		if !ok {
			// if not existent in jsonObj then skip
			return nil
		}

		s, ok := jsonValue.(string)
		if !ok {
			return fmt.Errorf("failed to convert \"%s\" in %v to bool", field.Name, jsonObj)
		}

		jsonObj[jsonTag] = toBool(s)
	}

	return nil
}

func alterStringToInt(jsonObj jsonwrap.JsonObj, field reflect.StructField) error {
	// for those struct fields where type = bool, convert string to bool
	if field.Type.String() == "int" {
		// validate st
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			return fmt.Errorf("json tag not found for field %s", field.Name)
		}

		jsonValue, ok := jsonObj[jsonTag]
		if !ok {
			// if not existent in jsonObj then skip
			return nil
		}

		s, ok := jsonValue.(string)
		if !ok {
			return fmt.Errorf("failed to convert \"%s\" in %v to int", field.Name, jsonObj)
		}

		jsonObj[jsonTag] = toInt(s)
	}

	return nil
}

func ConvertBoolean(inputFile, targetFile string) error {
	var jsonObjArray []jsonwrap.JsonObj
	err := jsonwrap.Read(inputFile, &jsonObjArray)
	if err != nil {
		return fmt.Errorf("failed to read json array from %s, %s", inputFile, err)
	}

	step := DetailedStep{}
	structFields := reflect.VisibleFields(reflect.TypeOf(step))

	for i, jsonMap := range jsonObjArray {
		for _, field := range structFields {
			alterStringToBool(jsonMap, field)
			alterStringToInt(jsonMap, field)
		}
		jsonObjArray[i] = jsonMap
	}

	bytes, err := json.Marshal(jsonObjArray)
	if err != nil {
		return fmt.Errorf("failed to marshal json array, %s", err)
	}

	var dst []DetailedStep
	err = json.Unmarshal(bytes, &dst)
	if err != nil {
		return fmt.Errorf("failed to unmarshal to DetailedStep, %s", err)
	}

	err = jsonwrap.WriteJsonToFile(dst, targetFile)

	return nil
}
