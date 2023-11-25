package edits

type Range struct {
	StartLineNumber int `json:"startLineNumber"`
	StartColumn     int `json:"startColumn"`
	EndLineNumber   int `json:"endLineNumber"`
	EndColumn       int `json:"endColumn"`
}

type SingleEditOperation struct {
	Text  string `json:"text"`
	Range Range  `json:"range"`
}

func toOpToAdd(chunk ChunkToAdd) SingleEditOperation {
	return SingleEditOperation{
		Text: chunk.Content,
		Range: Range{
			StartLineNumber: chunk.LineNumber,
			EndLineNumber:   chunk.LineNumber,
			StartColumn:     chunk.Column,
			EndColumn:       chunk.Column,
		},
	}
}

func toOpToDelete(chunk ChunkToDelete) SingleEditOperation {
	return SingleEditOperation{
		Text: "", // replace by "" means deletion of the range
		Range: Range{
			StartLineNumber: chunk.StartLineNumber,
			EndLineNumber:   chunk.EndLineNumber,
			StartColumn:     chunk.StartColumn,
			EndColumn:       chunk.EndColumn,
		},
	}
}

func toOpsToAdd(chunks []ChunkToAdd) []SingleEditOperation {
	var ops []SingleEditOperation
	for _, v := range chunks {
		op := toOpToAdd(v)
		ops = append(ops, op)
	}

	return ops
}

func toOpsToDelete(chunks []ChunkToDelete) []SingleEditOperation {
	var ops []SingleEditOperation
	for _, v := range chunks {
		op := toOpToDelete(v)
		ops = append(ops, op)
	}

	return ops
}
