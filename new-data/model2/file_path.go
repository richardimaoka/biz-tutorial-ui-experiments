package model2

import (
	"fmt"
	"strings"
)

func lessFilePath(a, b []*string) bool {
	if len(a) == 0 && len(b) == 0 {
		return false //even if len(b) == 0
	} else if /* len (a) != 0 && */ len(b) == 0 {
		return false //here, len(a) != 0
	} else if len(a) == 0 /* && len (b) != 0 */ {
		return false //even if len(b) == 0
	}

	// now len(a) != 0 AND len(b) != 0

	if a[0] == b[0] {
		return lessFilePath(a[1:], b[1:])
	} else {
		return *a[0] < *b[0]
	}
}

func filePathPtrSlice(filePath string) []*string {
	split := strings.Split(filePath, "/")

	var filePathSlice []*string
	for i := range split {
		filePathSlice = append(filePathSlice, &split[i]) // cannot use v of `for i, v := range ...` because v has the same address throughout the loop
	}

	return filePathSlice
}

func parentDirectoryPath(filePath string) string {
	split := strings.Split(filePath, "/")
	return strings.Join(split[:len(split)-1], "/")
}

func validateFilePath(filePath string) error {
	if filePath == "" {
		return fmt.Errorf("empty path")
	}
	if strings.HasSuffix(filePath, "/") {
		return fmt.Errorf("directory path = %s ends in slash", filePath)
	}
	return nil
}
