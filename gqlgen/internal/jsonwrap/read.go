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
		return "", fmt.Errorf("\"%s\" does not exist in JSON = %s", fromField, jsonBytes)
	}

	typeName, ok := typeNameRaw.(string)
	if !ok {
		return "", fmt.Errorf("\"%s\" is not a string, but found in type = %v", fromField, reflect.TypeOf(typeNameRaw))
	}

	return typeName, nil
}

func Read(filePath string, v interface{}) error {
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

func MarshalThenUnmarshal(obj JsonObj, unmarshaller func(jsonBytes []byte) error) error {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	err = unmarshaller(jsonBytes)
	if err != nil {
		return err
	}

	return nil
}

func WriteJsonValueToFile(v any, filePath string) error {
	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filePath, bytes, 0644); err != nil {
		return err
	}
	return nil
}
