package effect_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/preprocess/effect"
)

//write table-based go test for ToOperation
func Test_PageStateEffectToOperation(t *testing.T) {
	cmd := "TerminalCommand"
	cmdOut := "TerminalCommandWithOutput"
	cmdCd := "TerminalCommandWithCd"
	// cmdOutCd := "TerminalCommandWithOutputCd"
	scFile := "SourceCodeFileOperation"
	// scGit := "SourceCodeGitOperation"

	steps := []struct {
		TutorialName     string
		SeqNo            int
		SourceCodeOpType string
		TerminalOpType   string
		fileOpsSize      int
		CommitHash       string
	}{
		{"protoc-go-experiments", 0, scFile, cmd, 1, ""},
		{"protoc-go-experiments", 1, scFile, cmdCd, 0, ""},
		{"protoc-go-experiments", 3, scFile, cmdCd, 0, ""},
		{"protoc-go-experiments", 5, scFile, cmdOut, 0, ""},
		{"protoc-go-experiments", 8, scFile, cmd, 1, ""},
		{"gqlgensandbox", 0, scFile, cmd, 1, ""},
		{"gqlgensandbox", 1, scFile, cmdCd, 0, ""},
		{"gqlgensandbox", 2, scFile, "", 0, ""},
		{"gqlgensandbox", 5, scFile, cmd, 0, ""},
	}

	for i, step := range steps {
		t.Run(fmt.Sprintf("case[%d]", i), func(t *testing.T) {
			dirName := fmt.Sprintf("testdata/pagestate/%s", step.TutorialName)

			fileEffects, err := effect.ReadFileEffects(dirName + "/file-effects.json")
			if err != nil {
				t.Errorf("failed: %v", err)
			}

			terminalEffects, err := effect.ReadTerminalEffects(dirName + "/terminal-effects.json")
			if err != nil {
				t.Errorf("failed: %v", err)
			}

			markdownEffects, err := effect.ReadMarkdownEffects(dirName + "/markdown-effects.json")
			if err != nil {
				t.Errorf("failed: %v", err)
			}

			// TerminalEffect for seqNo
			tEff := terminalEffects.FindBySeqNo(step.SeqNo)

			// SourceCodeEffect for seqNo
			fEffs := fileEffects.FilterBySeqNo(step.SeqNo)
			scEff := effect.NewSourceCodeEffect(step.SeqNo, step.CommitHash, fEffs)

			// MarkdownEffect for seqNo
			mEff := markdownEffects.FindBySeqNo(step.SeqNo)

			// PageStateEffect for seqNo
			psEff := effect.NewPageStateEffect(step.SeqNo, "", "", scEff, tEff, mEff)
			op, err := psEff.ToOperation()
			if err != nil {
				t.Fatalf("ToOperation failed: %v", err)
			}

			if step.SourceCodeOpType == "" {
				if op.SourceCodeOperation != nil {
					t.Fatalf("unexpected terminal op type: %s, expected nil", reflect.TypeOf(op.SourceCodeOperation).Name())
				}
			} else {
				scOpType := reflect.TypeOf(op.SourceCodeOperation).Name()
				if scOpType != step.SourceCodeOpType {
					t.Fatalf("unexpected source code type: %s, expected %s", scOpType, step.SourceCodeOpType)
				}
			}

			if step.TerminalOpType == "" {
				if op.TerminalOperation != nil {
					t.Fatalf("unexpected terminal op type: %s, expected nil", reflect.TypeOf(op.TerminalOperation).Name())
				}
			} else {
				terminalOpType := reflect.TypeOf(op.TerminalOperation).Name()
				if terminalOpType != step.TerminalOpType {
					t.Fatalf("unexpected terminal op type: %s, expected %s", terminalOpType, step.TerminalOpType)
				}
			}
		})
	}
}
