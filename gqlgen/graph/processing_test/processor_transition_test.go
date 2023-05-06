package processing_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/processing"
)

var terminalEffectsFile = "testdata/terminal_effects.json"
var sourceCodeUnitEffectsFile = "testdata/source_code_unit_effects.json"

func Test_ReadTerminaEffects(t *testing.T) {
	_, err := processing.ReadTerminalEffects(terminalEffectsFile)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_ReadSourceCodeUnitEffects(t *testing.T) {
	_, err := processing.ReadSourceCodeUnitEffect(sourceCodeUnitEffectsFile)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_MergeEffects(t *testing.T) {
	terminalEffects, _ := processing.ReadTerminalEffects(terminalEffectsFile)
	sourceCodeUnitEffects, _ := processing.ReadSourceCodeUnitEffect(sourceCodeUnitEffectsFile)

	transitions, err := processing.MergeEffects(terminalEffects, sourceCodeUnitEffects)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(transitions); i++ {
		if transitions[i].SeqNo != i {
			t.Errorf("expected seqNo=%d, but got seqNo=%d", i, transitions[i].SeqNo)
		}
	}
}

//from git, do the same
