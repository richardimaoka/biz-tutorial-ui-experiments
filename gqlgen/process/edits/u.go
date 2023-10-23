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
	EACH_CHARACTER_ADDITIONS = "each-character"
	WORD_BY_WORD_ADDITIONS   = "word-by-word"
	WHOLE_LINE_ADDITIONS     = "whole-line"
)

func eachCharacterAdditions(toAdd string) []string {
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

func wordByWordAdditions(toAdd string) []string {
	return strings.SplitAfter(toAdd, " ")
}

func wholeLineAddition(toAdd string) []string {
	return []string{toAdd}
}

func condition(toAdd string) string {
	length := len(toAdd)
	if length < 10 {
		return EACH_CHARACTER_ADDITIONS
	} else if length < 100 {
		return WORD_BY_WORD_ADDITIONS
	} else {
		return WHOLE_LINE_ADDITIONS
	}
}

func conditionalAdditions(toAdd string) []string {
	cond := condition(toAdd)
	switch cond {
	case EACH_CHARACTER_ADDITIONS:
		return eachCharacterAdditions(toAdd)
	case WORD_BY_WORD_ADDITIONS:
		return wordByWordAdditions(toAdd)
	case WHOLE_LINE_ADDITIONS:
		return wholeLineAddition(toAdd)
	default:
		return wholeLineAddition(toAdd)
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
