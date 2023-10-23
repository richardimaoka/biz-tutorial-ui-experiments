package edits_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

func TestChunkDesign(t *testing.T) {
	cases := []struct {
		input      internal.Chunk
		goldenFile string
	}{
		{
			internal.Chunk{
				Content: "import Editor from \"@monaco-editor/react\";\n",
				Type:    "Add",
			},
			"testdata/1.json",
		},
	}

	for _, c := range cases {
		t.Run(c.goldenFile, func(t *testing.T) {
			// internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, filePatch)
		})
	}
}
