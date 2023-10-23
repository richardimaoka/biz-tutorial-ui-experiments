package edits

import (
	"strings"
	"unicode/utf8"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

type PositionedChunk struct {
	LineNumber int            `json:"lineNumber"`
	Column     int            `json:"column"`
	Chunk      internal.Chunk `json:"chunk"`
}

type state struct {
	lineNumber int
	column     int
}

func Convert(fileChunks []internal.Chunk) []PositionedChunk {
	s := state{0, 0}

	// placeholder operation to avoid compilation error `declared but not used`
	s.lineNumber = 0

	// 1. for each chunks
	return nil
}

func splitSingleLineAdd(singleLine string) []string {
	split := strings.SplitAfter(singleLine, "\n")

	// if singleLine ends in "\n", the new-line character
	lastIndex := len(split) - 1
	if split[lastIndex] == "" {
		// then omit the last "", empty string
		return split[0:lastIndex]
	}

	return split
}

func moveNewLineToHead(singleLine string) []string {
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

const (
	BREAKDOWN_TO_CHARACTERS = "each-character"
	BREAKDOWN_TO_WORDS      = "word-by-word"
	BREAKDOWN_TO_WHOLE_LINE = "whole-line"
)

func breakdownToCharacters(toAdd string) []string {
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

func breakdownToWords(toAdd string) []string {
	return strings.SplitAfter(toAdd, " ")
}

func breakdownToWholeLine(toAdd string) []string {
	return []string{toAdd}
}

func condition(toAdd string) string {
	length := len(toAdd)
	if length < 10 {
		return BREAKDOWN_TO_CHARACTERS
	} else if length < 100 {
		return BREAKDOWN_TO_WORDS
	} else {
		return BREAKDOWN_TO_WHOLE_LINE
	}
}

func breakdownAddition(toAdd string) []string {
	cond := condition(toAdd)
	switch cond {
	case BREAKDOWN_TO_CHARACTERS:
		return breakdownToCharacters(toAdd)
	case BREAKDOWN_TO_WORDS:
		return breakdownToWords(toAdd)
	case BREAKDOWN_TO_WHOLE_LINE:
		return breakdownToWholeLine(toAdd)
	default:
		return breakdownToWholeLine(toAdd)
	}
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
