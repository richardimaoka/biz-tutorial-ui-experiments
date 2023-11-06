package internal

import (
	"encoding/json"
	"os"
)

type JsonObj map[string]interface{}

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
