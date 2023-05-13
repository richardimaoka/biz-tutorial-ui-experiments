package effect

import (
	"encoding/json"
	"fmt"
	"os"
)

func jsonRead(filePath string, unmarshaller func(jsonBytes []byte) error) error {
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
