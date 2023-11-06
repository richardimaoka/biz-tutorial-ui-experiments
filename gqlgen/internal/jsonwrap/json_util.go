package jsonwrap

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type JsonObj map[string]interface{}

func ExtractTypeName(jsonBytes []byte, fromField string) (string, error) {
	var unmarshaled map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &unmarshaled); err != nil {
		return "", err
	}

	typeNameRaw, ok := unmarshaled[fromField]
	if !ok {
		return "", fmt.Errorf("\"%s\" does not exist in JSON", fromField)
	}

	typeName, ok := typeNameRaw.(string)
	if !ok {
		return "", fmt.Errorf("\"%s\" is not a string, but found in type = %v", fromField, reflect.TypeOf(typeNameRaw))
	}

	return typeName, nil
}

func JsonRead(filePath string, v interface{}) error {
	jsonBytes, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file = '%s', %s", filePath, err)
	}

	err = json.Unmarshal(jsonBytes, v)
	if err != nil {
		return fmt.Errorf("failed to unmarshal file = '%s', %s", filePath, err)
	}

	return nil
}
