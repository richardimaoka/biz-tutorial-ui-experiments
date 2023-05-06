package processing_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/processing"
)

func Test_ReadTerminaEffects(t *testing.T) {
	_, err := processing.ReadTerminalEffects("testdata/test.json")
	if err != nil {
		t.Fatal(err)
	}

	// cases := []struct {
	// 	ExpectedFile string
	// 	Step 1
	// }

	// combine source code effects and terminal effects
	// write them as gold files
	// then compare them
}

//from git, do the same
