package edits

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"

func ProcessChunk(chunk internal.Chunk, pos TypingPosition) (TypingPosition, []SingleEditOperation) {
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

func ProcessChunks(chunks []internal.Chunk) []SingleEditOperation {
	currentPos := TypingPosition{LineNumber: 1, Column: 1}

	var ops []SingleEditOperation
	for _, c := range chunks {
		var newOps []SingleEditOperation
		currentPos, newOps = ProcessChunk(c, currentPos)
		ops = append(ops, newOps...)
	}

	return ops
}
