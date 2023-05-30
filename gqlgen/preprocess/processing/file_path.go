package processing

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
		// e.g. a = "aaa/a.txt", b = "aaa/b.txt"
		return lessFilePathInner(a[1:], b[1:])
	} else {
		// e.g. a = "aaa/a.txt", b = "bbb/b.txt"
		// e.g. a = "a.txt",     b = "b.txt"
		return a[0] < b[0]
	}
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
