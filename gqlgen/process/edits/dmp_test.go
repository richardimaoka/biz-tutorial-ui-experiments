package edits_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/gitwrap"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/edits"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func A(beforeText, afterText string) []gitwrap.Chunk {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(beforeText, afterText, true)

	var chunks []gitwrap.Chunk
	for _, d := range diffs {
		var chunkType string
		switch d.Type {
		case diffmatchpatch.DiffEqual:
			chunkType = "Equal"
		case diffmatchpatch.DiffInsert:
			chunkType = "Add"
		case diffmatchpatch.DiffDelete:
			chunkType = "Delete"
		}

		chunks = append(chunks, gitwrap.Chunk{Type: chunkType, Content: d.Text})
	}

	return chunks
}

func TestDmp(t *testing.T) {
	cases := []struct {
		name           string
		beforeText     string
		afterText      string
		expectedChunks []gitwrap.Chunk
	}{
		{
			"test1",
			"FROM ubuntu\nCMD [\"executable\", \"param1\", \"param2\"]",
			"FROM ubuntu\nCMD [\"echo\", \"param1\", \"param2\"]",
			[]gitwrap.Chunk{
				{Type: "Equal", Content: "FROM ubuntu\nCMD [\"e"},
				{Type: "Delete", Content: "xe"},
				{Type: "Equal", Content: "c"},
				{Type: "Delete", Content: "utable"},
				{Type: "Add", Content: "ho"},
				{Type: "Equal", Content: "\", \"param1\", \"param2\"]"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			resultChunks := A(c.beforeText, c.afterText)
			if diff := cmp.Diff(c.expectedChunks, resultChunks); diff != "" {
				t.Fatalf("mismatch (-expected +result):\n%s", diff)
			}
		})
	}
}

func TestDmp2(t *testing.T) {
	cases := []struct {
		name       string
		beforeText string
		afterText  string
		expected   []edits.SingleEditOperation
	}{
		{
			"test1",
			"FROM ubuntu\nCMD [\"executable\", \"param1\", \"param2\"]",
			"FROM ubuntu\nCMD [\"echo\", \"param1\", \"param2\"]",
			[]edits.SingleEditOperation{
				{
					Text:  "",
					Range: edits.Range{StartLineNumber: 2, EndLineNumber: 2, StartColumn: 8, EndColumn: 10},
				},
				{
					Text:  "",
					Range: edits.Range{StartLineNumber: 2, EndLineNumber: 2, StartColumn: 9, EndColumn: 15},
				},
				{
					Text:  "h",
					Range: edits.Range{StartLineNumber: 2, EndLineNumber: 2, StartColumn: 9, EndColumn: 9},
				},
				{
					Text:  "o",
					Range: edits.Range{StartLineNumber: 2, EndLineNumber: 2, StartColumn: 10, EndColumn: 10},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			chunks := A(c.beforeText, c.afterText)
			result := edits.ToOperations(chunks)
			if diff := cmp.Diff(c.expected, result); diff != "" {
				t.Fatalf("mismatch (-expected +result):\n%s", diff)
			}
		})
	}
}
