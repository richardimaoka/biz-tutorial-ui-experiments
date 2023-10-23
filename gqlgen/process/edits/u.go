package edits

import (
	"strings"
	"unicode/utf8"
)

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

func MoveNewLineToHead(singleLine string) []string {
	if singleLine == "\n" {
		return []string{"\n"}
	}

	if strings.HasSuffix(singleLine, "\n") {
		lastIndex := len(singleLine) - 1
		lastNewLineOmitted := singleLine[0:lastIndex]

		return []string{"\n", lastNewLineOmitted}
	}

	return []string{singleLine}
}

func OneLetterAdditions(toAdd string) []string {
	if toAdd == "" {
		return []string{""}
	}

	var additions []string

	// From プログラミング言語Go chap. 3, p. 78
	// since a UTF-8 character varies in its size, we need to use rune (via utf8.DecodeRuneInString)
	// to extract each character. (e.g.) Source code comments can have multi-byte characters.
	for i := 0; i < len(toAdd); {
		r, size := utf8.DecodeRuneInString(toAdd[i:])
		additions = append(additions, string(r))
		i += size
	}

	return additions
}

func WordByWordAdditions(toAdd string) []string {
	return []string{""}
}

func WholeLineAddition(toAdd string) string {
	return ""
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
