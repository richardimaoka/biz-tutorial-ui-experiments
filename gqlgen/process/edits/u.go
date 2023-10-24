package edits

import (
	"strings"
	"unicode/utf8"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

type PositionedChunk struct {
	LineNumber int              `json:"lineNumber"`
	Column     int              `json:"column"`
	Chunks     []internal.Chunk `json:"chunks"`
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

// Split chunkContent, which potentially has many '\n' characters inside,
// into slice of single-line strings, where each of them can have '\n' at the end.
//
// (e.g.)
//    chunkContent = "func myFunc() int {\n  var myVar string\n  var anotherVar int"
//
//  will be split into:
//
//    []string {
//      "func myFunc() int {\n",   // '\n' at the end.
//      "var myVar string\n",      // '\n' at the end
//      "var anotherVar int",      // doesn't have '\n'
//    }
func splitIntoSingleLines(chunkContent string) []string {
	split := strings.SplitAfter(chunkContent, "\n")

	// if singleLine ends in "\n", the new-line character
	lastIndex := len(split) - 1
	if split[lastIndex] == "" {
		// then omit the last "", empty string
		return split[0:lastIndex]
	}

	return split
}

// Returns a pair indicating 1) whether the input has a new character '\n'
// and 2) the string without '\n'
// parameters:
//   singleLineToAdd: the input string, which potentially has '\n' at the end
//                    but cannot have '\n' in the middle
//
// returns:
//   bool   : whether there is `\n` at the end
//   string : the remaining string without '\n'
func stripNewLineAtEnd(singleLineToAdd string) (bool, string) {
	if singleLineToAdd == "\n" {
		return true, "" // the latter is empty "", since there is no remaining string after omitting '\n'
	}

	if strings.HasSuffix(singleLineToAdd, "\n") {
		lastIndex := len(singleLineToAdd) - 1
		lastNewLineOmitted := singleLineToAdd[0:lastIndex]

		return true, lastNewLineOmitted
	}

	return false, singleLineToAdd
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

// Split a single-line change (addition) into a slice of small-piece `string`s
//
//   singleLineToAdd: a string add, which potentially has '\n' at the end
//                    (but cannot have '\n' in the middle)
// func breakDownSingleLine(singleLineToAdd string) []string {
// 	sliced := moveNewLineToHead(singleLineToAdd)

// 	var split []string
// 	for _, s := range sliced {
// 		breakDowns := breakdownAddition(s)
// 		split = append(split, breakDowns...)
// 	}

// 	return split
// }

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
