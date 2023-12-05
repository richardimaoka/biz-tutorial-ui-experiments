package edits_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/gitwrap"
	"github.com/sergi/go-diff/diffmatchpatch"
)

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

	dmp := diffmatchpatch.New()

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			diffs := dmp.DiffMain(c.beforeText, c.afterText, true)

			var resultChunks []gitwrap.Chunk
			for _, d := range diffs {
				var chunkType string
				switch d.Type {
				case diffmatchpatch.DiffEqual:
					chunkType = "Equal"
				case diffmatchpatch.DiffInsert:
					chunkType = "Add"
				case diffmatchpatch.DiffDelete:
					chunkType = "Delete"
				default:
					t.Fatalf("chunkType = '%s' is invalid for test = %s", d.Type, c.name)
				}

				resultChunks = append(resultChunks, gitwrap.Chunk{Type: chunkType, Content: d.Text})
			}

			if diff := cmp.Diff(c.expectedChunks, resultChunks); diff != "" {
				t.Fatalf("mismatch (-expected +result):\n%s", diff)
			}
		})
	}
}
