package edits

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/gitwrap"

func ToOperations(chunks []gitwrap.Chunk) []SingleEditOperation {
	currentPos := TypingPosition{LineNumber: 1, Column: 1}

	var ops []SingleEditOperation
	for _, c := range chunks {
		var newOps []SingleEditOperation
		currentPos, newOps = processChunk(c, currentPos)
		ops = append(ops, newOps...)
	}

	return ops
}
