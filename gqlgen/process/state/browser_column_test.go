package state_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func TestBrowserColumn(t *testing.T) {
	// test case struct
	cases := []struct {
		imagePath      string
		expectedWidth  int
		expectedHeight int
	}{
		{"testdata/images/todos-GraphiQL-successful.png", 970, 1080},
	}

	col := state.NewBrowserColumn()
	for _, c := range cases {
		t.Run(c.imagePath, func(t *testing.T) {
			width, height, err := col.ImageDimension(c.imagePath)
			if err != nil {
				t.Fatalf("ImageDimension() failed, %s", err)
			}
			if width != c.expectedWidth {
				t.Errorf("ImageDimension() width = %d, want %d", width, c.expectedWidth)
			}
			if height != c.expectedHeight {
				t.Errorf("ImageDimension() height = %d, want %d", height, c.expectedHeight)
			}
		})
	}
}
