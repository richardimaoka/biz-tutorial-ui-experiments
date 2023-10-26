package edits_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/edits"
)

func Test(t *testing.T) {
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
