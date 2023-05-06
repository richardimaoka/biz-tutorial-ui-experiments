package processing

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func extractTypeName(jsonBytes []byte, fromField string) (string, error) {
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

func jsonRead(callerName, filePath string, unmarshaller func(jsonBytes []byte) error) error {
	jsonBytes, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("%s failed to read file, %s", callerName, err)
	}

	err = unmarshaller(jsonBytes)
	if err != nil {
		return fmt.Errorf("%s failed to unmarshal, %s", callerName, err)
	}

	return nil
}
