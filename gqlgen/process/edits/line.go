package edits

import (
	"strings"
	"unicode/utf8"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

type ChunkToAdd struct {
	TypingPosition
	Content string `json:"content"`
	Type    string `json:"type"`
}

type TypingPosition struct {
	LineNumber int `json:"lineNumber"`
	Column     int `json:"column"`
}

type RangeToDelete struct {
	LineNumber  int
	StartColumn int
	EndColumn   int
}

type ChunkToDelete struct {
	Range RangeToDelete
}

type SingleLineToAdd struct {
	NewLineAtEnd          bool
	ContentWithoutNewLine string
}

type state struct {
	lineNumber int
	column     int
}

func Convert(fileChunks []internal.Chunk) []ChunkToAdd {
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
func detectNewLine(lineToAdd string) SingleLineToAdd {
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

// Returns single-line changes with flags indicating new-line characters '\n' are included
func splitChunkToLines(chunk internal.Chunk) []SingleLineToAdd {
	splitContent := splitAfterNewLine(chunk.Content)

	var ret []SingleLineToAdd
	for _, c := range splitContent {
		lineToAdd := detectNewLine(c)
		ret = append(ret, lineToAdd)
	}

	return ret
}

// Split a single-line change (addition) into a slice of small-piece `string`s
//
// parameters:
//   singleLineToAdd: the input string, which potentially has '\n' at the end
//                    but cannot have '\n' in the middle
func lineToPosChunks(lineToAdd SingleLineToAdd, pos TypingPosition) (TypingPosition, []ChunkToAdd) {
	var pChunks []ChunkToAdd

	if lineToAdd.NewLineAtEnd {
		// if new line '\n' at the end, then moves it to the beginning
		firstChunk := ChunkToAdd{
			TypingPosition: TypingPosition{
				LineNumber: pos.LineNumber,
				Column:     pos.Column,
			},
			Type:    "Add",
			Content: "\n",
		}
		pChunks = append(pChunks, firstChunk)
	}

	currentColumn := pos.Column
	breakDowns := breakdownLineToAdd(lineToAdd.ContentWithoutNewLine)
	for _, b := range breakDowns {
		c := ChunkToAdd{
			TypingPosition: TypingPosition{
				LineNumber: pos.LineNumber,
				Column:     currentColumn,
			},
			Type:    "Add",
			Content: b,
		}
		pChunks = append(pChunks, c)
		currentColumn += utf8.RuneCountInString(b)
	}

	if lineToAdd.NewLineAtEnd {
		return TypingPosition{LineNumber: pos.LineNumber + 1, Column: 1}, pChunks
	} else {
		return TypingPosition{LineNumber: pos.LineNumber, Column: currentColumn}, pChunks
	}
}

// Split the given chunk to single-line changes,
// then convert each line into a slice of positioned chunks
//
// inernal.Chunk  : represents a chunk from git diff
// TypingPosition : is the position at which the function call is made
func toPositionedChunks(chunk internal.Chunk, pos TypingPosition) (TypingPosition, []ChunkToAdd) {
	linesToAdd := splitChunkToLines(chunk)

	var pChunks []ChunkToAdd
	var currentPos TypingPosition = pos
	for _, v := range linesToAdd {
		var newPosChunks []ChunkToAdd
		currentPos, newPosChunks = lineToPosChunks(v, currentPos)

		pChunks = append(pChunks, newPosChunks...)
	}

	return currentPos, pChunks
}

func processEqual(chunk internal.Chunk, pos TypingPosition) TypingPosition {
	split := strings.Split(chunk.Content, "\n")
	lastLineChange := split[len(split)-1]

	return TypingPosition{
		LineNumber: pos.LineNumber + len(split) - 1,
		Column:     utf8.RuneCountInString(lastLineChange),
	}
}

func processDelete(chunk internal.Chunk, pos TypingPosition) []ChunkToDelete {
	linesToDelete := splitAfterNewLine(chunk.Content)

	var dChunks []ChunkToDelete
	for _, lineString := range linesToDelete {
		nChars := utf8.RuneCountInString(lineString)
		c := ChunkToDelete{
			Range: RangeToDelete{
				LineNumber:  pos.LineNumber,
				StartColumn: pos.Column,
				EndColumn:   pos.Column + nChars,
			},
		}
		dChunks = append(dChunks, c)
	}

	return dChunks
}
