package pkg2

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func ExtractTypeName(bytes []byte, fromField string) (string, error) {
	var unmarshaled map[string]interface{}
	if err := json.Unmarshal(bytes, &unmarshaled); err != nil {
		return "", err
	}

	typeNameRaw, ok := unmarshaled[fromField]
	if !ok {
		return "", fmt.Errorf("\"%s\" does not exist in JSON", fromField)
	}

	typeName, ok := typeNameRaw.(string)
	if !ok {
		return "", fmt.Errorf("\"%s\" is not a string, but found %s", fromField, reflect.TypeOf(typeNameRaw))
	}

	return typeName, nil
}
