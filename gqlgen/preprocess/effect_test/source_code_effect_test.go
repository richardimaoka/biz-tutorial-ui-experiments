package effect_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/preprocess/effect"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/preprocess/processing"
)

//write table-based go test for ToOperation
func Test_SourceCodeEffectToOperation1(t *testing.T) {
	cases := []struct {
		fileEffectsFile string
		seqNo           int
		fileOpsSize     int
	}{
		{"testdata/sourcecode/file-effects1.json", 0, 1},
		{"testdata/sourcecode/file-effects1.json", 1, 0},
		{"testdata/sourcecode/file-effects1.json", 21, 2},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case[%d]", i), func(t *testing.T) {
			fileEffects, err := effect.ReadFileEffects(c.fileEffectsFile)
			if err != nil {
				t.Fatalf("ReadFileEffects failed: %v", err)
			}

			fEffs := fileEffects.FilterBySeqNo(c.seqNo)
			scEff := effect.NewSourceCodeEffect(c.seqNo, "", fEffs)
			scOp, err := scEff.ToOperation()
			if err != nil {
				t.Fatalf("ToOperation failed: %v", err)
			}

			op, ok := scOp.(processing.SourceCodeFileOperation)
			if !ok {
				t.Fatalf("unexpected type: %s", reflect.TypeOf(scOp))
			}

			fileOpsSize := len(op.FileOps)
			if fileOpsSize != c.fileOpsSize {
				t.Fatalf("file ops size mismatch: expected %d, but got %d", c.fileOpsSize, fileOpsSize)
			}
		})
	}
}

func Test_SourceCodeEffectToOperation2(t *testing.T) {
	cases := []struct {
		fileEffectsFile string
		seqNo           int
		commitHash      string
	}{
		{"testdata/sourcecode/file-effects2.json", 6, "7e799a483e2ba57c1d4897fc364398af22ea1627"},
		{"testdata/sourcecode/file-effects2.json", 7, "c59843373c53e3a0185576d0c1eb192b59582134"},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("case[%d]", i), func(t *testing.T) {
			fileEffects, err := effect.ReadFileEffects(c.fileEffectsFile)
			if err != nil {
				t.Fatalf("ReadFileEffects failed: %v", err)
			}

			fEffs := fileEffects.FilterBySeqNo(c.seqNo)
			scEff := effect.NewSourceCodeEffect(c.seqNo, c.commitHash, fEffs)
			scOp, err := scEff.ToOperation()
			if err != nil {
				t.Fatalf("ToOperation failed: %v", err)
			}

			op, ok := scOp.(processing.SourceCodeGitOperation)
			if !ok {
				t.Fatalf("unexpected type: %s", reflect.TypeOf(scOp))
			}

			if op.CommitHash != c.commitHash {
				t.Fatalf("commit hash mismatch: expected %s, but got %s", c.commitHash, op.CommitHash)
			}
		})
	}
}
