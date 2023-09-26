package rough

import (
	"fmt"
	"reflect"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

func ToBool(s string) bool {
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

func alterStringToBool(jsonObj internal.JsonObj, field reflect.StructField) error {
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

		jsonObj[field.Name] = ToBool(s)
	}

	return nil
}

func ConvertBoolean(inputFile, targetFile string) error {
	jsonObjArray, err := internal.JsonReadArray(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read json array from %s, %s", inputFile, err)
	}

	step := DetailedStep{}
	structFields := reflect.VisibleFields(reflect.TypeOf(step))

	for i, jsonMap := range jsonObjArray {
		for _, field := range structFields {
			alterStringToBool(jsonMap, field)
		}
		jsonObjArray[i] = jsonMap
	}

	err = internal.WriteJsonToFile(jsonObjArray, targetFile)

	return nil
}
