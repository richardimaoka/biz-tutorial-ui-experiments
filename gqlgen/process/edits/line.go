package edits

import (
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

type PositionedChunk struct {
	LineNumber int    `json:"lineNumber"`
	Column     int    `json:"column"`
	Content    string `json:"contentn"`
	Type       string `json:"type"`
}

type SingleLineToAdd struct {
	NewLineAtEnd          bool
	ContentWithoutNewLine string
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
func splitAfterNewLine(chunkContent string) []string {
	split := strings.SplitAfter(chunkContent, "\n")

	lastIndex := len(split) - 1
	// if chunkContent ends in "\n", strings.SplitAfter() returns "" as the last element, which is unnecessary...
	if split[lastIndex] == "" {
		// ...then omit the last ""
		return split[0:lastIndex]
	}

	return split
}

// Returns a pair indicating 1) whether the input has a new character '\n'
// and 2) the string without '\n'
// parameters:
//   lineToAdd: the input string, which potentially has '\n' at the end
//              but cannot have '\n' in the middle
func toSingleLineToAdd(lineToAdd string) SingleLineToAdd {
	if lineToAdd == "\n" {
		return SingleLineToAdd{
			NewLineAtEnd:          true,
			ContentWithoutNewLine: "", // the latter is empty "", since there is no remaining string after omitting '\n'
		}
	}

	if strings.HasSuffix(lineToAdd, "\n") {
		lastIndex := len(lineToAdd) - 1
		lastNewLineOmitted := lineToAdd[0:lastIndex]

		return SingleLineToAdd{
			NewLineAtEnd:          true,
			ContentWithoutNewLine: lastNewLineOmitted,
		}
	}

	return SingleLineToAdd{
		NewLineAtEnd:          false,
		ContentWithoutNewLine: lineToAdd,
	}
}

func splitIntoSingleLines(chunk internal.Chunk) []SingleLineToAdd {
	splitContent := splitAfterNewLine(chunk.Content)

	var ret []SingleLineToAdd
	for _, c := range splitContent {
		sa := toSingleLineToAdd(c)
		ret = append(ret, sa)
	}

	return ret
}

// Split a single-line change (addition) into a slice of small-piece `string`s
//
// parameters:
//   singleLineToAdd: the input string, which potentially has '\n' at the end
//                    but cannot have '\n' in the middle
// func splitIntoChunks(singleLineToAdd string, lineNumber, column int) []PositionedChunk {
// 	hasNewLine, contentWithoutNewLine := toSingleLineToAdd(singleLineToAdd)

// 	var pChunks []PositionedChunk

// 	if hasNewLine {
// 		firstChunk := PositionedChunk{
// 			LineNumber: lineNumber,
// 			Column:     column,
// 			Type:       "Add",
// 			Content:    "\n",
// 		}
// 		pChunks = append(pChunks, firstChunk)
// 	}

// 	currentColumn := column
// 	breakDowns := breakdownAddition(contentWithoutNewLine)
// 	for _, b := range breakDowns {
// 		c := PositionedChunk{
// 			LineNumber: lineNumber,
// 			Column:     currentColumn,
// 			Type:       "Add",
// 			Content:    b,
// 		}
// 		pChunks = append(pChunks, c)
// 		currentColumn += utf8.RuneCountInString(b)
// 	}

// 	return pChunks
// }
