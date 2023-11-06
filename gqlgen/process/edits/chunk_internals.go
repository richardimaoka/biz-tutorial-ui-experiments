package edits

import (
	"strings"
	"unicode/utf8"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/gitwrap"
)

type ChunkToAdd struct {
	TypingPosition
	Content string `json:"content"`
	Type    string `json:"type"`
}

type ChunkToDelete struct {
	RangeToDelete
	Content string //this is not necessary but convenient for debugging
}

type TypingPosition struct {
	LineNumber int `json:"lineNumber"`
	Column     int `json:"column"`
}

type RangeToDelete struct {
	StartLineNumber int
	EndLineNumber   int
	StartColumn     int
	EndColumn       int
}

type SingleLineChange struct {
	NewLineAtEnd          bool
	ContentWithoutNewLine string
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
func detectNewLine(lineChange string) SingleLineChange {
	if lineChange == "\n" {
		return SingleLineChange{
			NewLineAtEnd:          true,
			ContentWithoutNewLine: "", // the latter is empty "", since there is no remaining string after omitting '\n'
		}
	}

	if strings.HasSuffix(lineChange, "\n") {
		lastIndex := len(lineChange) - 1
		lastNewLineOmitted := lineChange[0:lastIndex]

		return SingleLineChange{
			NewLineAtEnd:          true,
			ContentWithoutNewLine: lastNewLineOmitted,
		}
	}

	return SingleLineChange{
		NewLineAtEnd:          false,
		ContentWithoutNewLine: lineChange,
	}
}

// Returns single-line changes with flags indicating new-line characters '\n' are included
func splitChunkToLines(chunk gitwrap.Chunk) []SingleLineChange {
	splitContent := splitAfterNewLine(chunk.Content)

	var ret []SingleLineChange
	for _, c := range splitContent {
		lineChange := detectNewLine(c)
		ret = append(ret, lineChange)
	}

	return ret
}

// Split a single-line change (addition) into a slice of small-piece `string`s
func lineToChunksToAdd(lineToAdd SingleLineChange, pos TypingPosition) (TypingPosition, []ChunkToAdd) {
	var chunks []ChunkToAdd

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
		chunks = append(chunks, firstChunk)
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
		chunks = append(chunks, c)
		currentColumn += utf8.RuneCountInString(b)
	}

	if lineToAdd.NewLineAtEnd {
		return TypingPosition{LineNumber: pos.LineNumber + 1, Column: 1}, chunks
	} else {
		return TypingPosition{LineNumber: pos.LineNumber, Column: currentColumn}, chunks
	}
}

// Split a single-line change (addition) into a slice of small-piece `string`s
func lineToChunksToDelete(lineToDelete SingleLineChange, pos TypingPosition) (TypingPosition, []ChunkToDelete) {
	var chunks []ChunkToDelete

	if lineToDelete.NewLineAtEnd {
		c := ChunkToDelete{
			Content: lineToDelete.ContentWithoutNewLine + "\n",
			RangeToDelete: RangeToDelete{
				StartLineNumber: pos.LineNumber,
				StartColumn:     pos.Column,
				EndLineNumber:   pos.LineNumber + 1,
				EndColumn:       0,
			},
		}
		chunks = append(chunks, c)
		newPos := TypingPosition{LineNumber: pos.LineNumber, Column: 1}
		return newPos, chunks
	} else {
		nChars := utf8.RuneCountInString(lineToDelete.ContentWithoutNewLine)
		c := ChunkToDelete{
			Content: lineToDelete.ContentWithoutNewLine,
			RangeToDelete: RangeToDelete{
				StartLineNumber: pos.LineNumber,
				EndLineNumber:   pos.LineNumber,
				StartColumn:     pos.Column,
				EndColumn:       pos.Column + nChars - 1,
			},
		}
		chunks = append(chunks, c)
		newPos := TypingPosition{LineNumber: pos.LineNumber, Column: 1}
		return newPos, chunks
	}
}

// If gitwrap.Chunk has Type == "Add", then return chunks to add.
//
// inernal.Chunk  : represents a chunk from git diff
// TypingPosition : is the position at which the function call is made
func toChunksToAdd(chunk gitwrap.Chunk, pos TypingPosition) (TypingPosition, []ChunkToAdd) {
	// Split into single-line changes first.
	// Special handling on '\n' is needed, so a slice of '\n'-aware structure is used
	linesToAdd := splitChunkToLines(chunk)

	var chunksToAdd []ChunkToAdd
	var currentPos TypingPosition = pos
	for _, v := range linesToAdd {
		var newChunksToAdd []ChunkToAdd
		currentPos, newChunksToAdd = lineToChunksToAdd(v, currentPos)

		chunksToAdd = append(chunksToAdd, newChunksToAdd...)
	}

	return currentPos, chunksToAdd
}

// If gitwrap.Chunk has Type == "Add", then return chunks to add.
// No need to move the typing position for deletion.
func toChunksToDelete(chunk gitwrap.Chunk, pos TypingPosition) []ChunkToDelete {
	// Split into single-line changes first.
	// Special handling on '\n' is needed, so a slice of '\n'-aware structure is used
	linesToDelete := splitChunkToLines(chunk)

	var chunksToDelete []ChunkToDelete
	var currentPos TypingPosition = pos
	for _, v := range linesToDelete {
		var newChunksToDelete []ChunkToDelete
		currentPos, newChunksToDelete = lineToChunksToDelete(v, currentPos)

		chunksToDelete = append(chunksToDelete, newChunksToDelete...)
	}

	return chunksToDelete
}

// If gitwrap.Chunk has Type == "Equal", then move the typing position but no edits to make
func moveTypingPosition(chunk gitwrap.Chunk, pos TypingPosition) TypingPosition {
	split := strings.Split(chunk.Content, "\n")
	lastLineChange := split[len(split)-1]

	return TypingPosition{
		LineNumber: pos.LineNumber + len(split) - 1,
		Column:     utf8.RuneCountInString(lastLineChange) + 1,
	}
}

func processChunk(chunk gitwrap.Chunk, pos TypingPosition) (TypingPosition, []SingleEditOperation) {
	currentPos := pos

	var ops []SingleEditOperation = []SingleEditOperation{}
	switch chunk.Type {
	case "Add":
		var chunks []ChunkToAdd
		currentPos, chunks = toChunksToAdd(chunk, currentPos)
		newOps := toOpsToAdd(chunks)
		ops = append(ops, newOps...)
	case "Equal":
		currentPos = moveTypingPosition(chunk, pos)
	case "Delete":
		chunks := toChunksToDelete(chunk, currentPos)
		newOps := toOpsToDelete(chunks)
		ops = append(ops, newOps...)
	}

	return currentPos, ops
}
