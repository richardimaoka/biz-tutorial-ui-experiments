package processing_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/processing"
)

func Test_ReadTerminaEffects(t *testing.T) {
	_, err := processing.ReadTerminalEffects("testdata/terminal_effects.json")
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

func Test_ReadSourceCodeUnitEffects(t *testing.T) {
	_, err := processing.ReadSourceCodeUnitEffect("testdata/source_code_unit_effects.json")
	if err != nil {
		t.Fatal(err)
	}
}

//from git, do the same
