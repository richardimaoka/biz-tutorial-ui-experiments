package state_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func TestSourceCodePatterns(t *testing.T) {
	repoUrl := "https://github.com/richardimaoka/next-sandbox.git"
	currentCommitHash := "8adac375628219e020d4b5957ff24f45954cbd3f"
	_, err := state.NewSourceCode(repoUrl, currentCommitHash)
	if err != nil {
		t.Fatalf("failed in TestSourceCodePatterns to create state.SourceCode, %s", err)
	}

	t.Fatalf("intentionally failed")
}
