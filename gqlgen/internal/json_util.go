package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

type JsonObj map[string]interface{}

func JsonRead(filePath string, unmarshaller func(jsonBytes []byte) error) error {
	jsonBytes, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s, %s", filePath, err)
	}

	err = unmarshaller(jsonBytes)
	if err != nil {
		return fmt.Errorf("failed to unmarshal %s, %s", filePath, err)
	}

	return nil
}

func JsonReadArray(filePath string) ([]JsonObj, error) {
	var arrayOfObj []JsonObj
	unmarshaller := func(jsonBytes []byte) error { return json.Unmarshal(jsonBytes, &arrayOfObj) }
	err := JsonRead(filePath, unmarshaller)

	if err != nil {
		return nil, fmt.Errorf("JsonReadArray failed to read file, %s", err)
	}

	return arrayOfObj, nil
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
