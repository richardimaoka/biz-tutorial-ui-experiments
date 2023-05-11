package processing_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/processing"
)

var terminalEffectsFile = "testdata/terminal_effects.json"
var fileEffectsFile = "testdata/file_effects.json"

func Test_ReadTerminaEffects(t *testing.T) {
	_, err := processing.ReadTerminalEffects(terminalEffectsFile)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_ReadSourceCodeUnitEffects(t *testing.T) {
	_, err := processing.ReadFileEffects(fileEffectsFile)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_MergeEffects(t *testing.T) {
	terminalEffects, _ := processing.ReadTerminalEffects(terminalEffectsFile)
	sourceCodeUnitEffects, _ := processing.ReadFileEffects(fileEffectsFile)

	effects, err := processing.MergeEffects(terminalEffects, sourceCodeUnitEffects)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(effects); i++ {
		if effects[i].SeqNo != i {
			t.Errorf("expected seqNo=%d, but got seqNo=%d", i, effects[i].SeqNo)
		}
	}

	if *updateFlag {
		internal.WriteGoldenFile(t, "testdata/transition_effects.json", effects)
	}
}

//from git, do the same
