package processing

import (
	"fmt"
	"strings"
)

func isValidFilePath(filePath string) error {
	if filePath == "" {
		return fmt.Errorf("empty path")
	}
	if strings.HasSuffix(filePath, "/") {
		return fmt.Errorf("directory path = %s ends in slash", filePath)
	}
	return nil
}
