package model2

import (
	"fmt"
	"strings"
)

func LessFilePath(a, b string) bool {
	splitA := strings.Split(a, "/")
	splitB := strings.Split(b, "/")
	return lessFilePathInner(splitA, splitB)
}

// for debugging
// func printStrPtrArray(a, b []*string) {
// 	for _, v := range a {
// 		fmt.Print(*v, ", ")
// 	}
// 	fmt.Print("vs. ")
// 	for _, v := range b {
// 		fmt.Print(*v, ", ")
// 	}
// 	fmt.Println()
// }

func lessFilePathInner(a, b []string) bool {
	if len(a) == 0 && len(b) == 0 {
		// a == b. even if len(b) == 0
		return false
	} else if /* len (a) != 0 && */ len(b) == 0 {
		// e.g. a = "aaa/b.txt", b = "aaa"
		return false
	} else if len(a) == 0 /* len (b) != 0 && */ {
		// e.g. a = "aaa", b = "aaa/abc"
		return true
	}

	// now len(a) != 0 AND len(b) != 0

	if a[0] == b[0] {
		return lessFilePathInner(a[1:], b[1:])
	} else {
		return a[0] < b[0]
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

func isValidFilePath(filePath string) error {
	if filePath == "" {
		return fmt.Errorf("empty path")
	}
	if strings.HasSuffix(filePath, "/") {
		return fmt.Errorf("directory path = %s ends in slash", filePath)
	}
	return nil
}
