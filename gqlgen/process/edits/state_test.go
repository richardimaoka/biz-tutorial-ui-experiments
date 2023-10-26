package edits_test

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/edits"
)

func TestProcessChunk(t *testing.T) {
	cases := []struct {
		inputPos    edits.TypingPosition
		inputChunk  internal.Chunk
		expected    []edits.SingleEditOperation
		expectedPos edits.TypingPosition
	}{
		{
			edits.TypingPosition{LineNumber: 1, Column: 1},
			internal.Chunk{
				Content: "\"use client\";\n\n",
				Type:    "Equal",
			},
			[]edits.SingleEditOperation{},
			edits.TypingPosition{LineNumber: 3, Column: 1},
		},
	}

	for index, c := range cases {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			resultPos, resultOps := edits.ProcessChunk(c.inputChunk, c.inputPos)
			if !cmp.Equal(c.expected, resultOps) {
				t.Errorf(cmp.Diff(c.expected, resultOps))
			}
			if resultPos != c.expectedPos {
				t.Errorf(cmp.Diff(c.expectedPos, resultPos))
			}
		})
	}
}

func TestProcessChunks(t *testing.T) {
	cases := []struct {
		inputFile  string
		goldenFile string
	}{
		{"testdata/chunks1.json", "testdata/ops-golden1.json"},
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			var chunks []internal.Chunk
			err := internal.JsonRead2(c.inputFile, &chunks)
			if err != nil {
				t.Fatal(err)
			}
			result := edits.ProcessChunks(chunks)
			internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, result)
		})
	}
}
