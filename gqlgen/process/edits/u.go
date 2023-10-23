package edits

import "strings"

func SplitSingleLineAdd(singleLine string) []string {
	split := strings.SplitAfter(singleLine, "\n")

	// if singleLine ends in "\n", the new-line character
	lastIndex := len(split) - 1
	if split[lastIndex] == "" {
		// then omit the last "", empty string
		return split[0:lastIndex]
	}

	return split
}

// type AtomicAddition struct {
// 	// LineNumber int
// 	ToAdd string
// }

// func OneLetterAdditions(toAdd string) []AtomicAddition {
// }

// func WordByWordAdditions(toAdd string) []AtomicAddition {
// }

// func WholeLineAddition(toAdd string) AtomicAddition {
// }

// func SplitChange(toAdd string) []string {
// 	return strings.SplitAfter(toAdd, "\n")
// }

// func InsertNewLineFirst(toAdd string) (AtomicAddition, []AtomicAddition) {
// }
