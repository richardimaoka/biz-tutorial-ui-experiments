package processing_test

/*
import (
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/preprocess/processing"
)

// test getting file diffs from git commit
func TestGetFileDiffsFromGitCommit(t *testing.T) {
	repoUrl := "https://github.com/richardimaoka/gqlgensandbox"
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: repoUrl})
	if err != nil {
		t.Fatalf("failed to initialize source code, cannot clone repo %s, %s", repoUrl, err)
	}

	currCommit, err := repo.CommitObject(plumbing.NewHash("490808086bded6b27f3651b095aefb7bb6708da2"))
	prevCommit, err := repo.CommitObject(plumbing.NewHash("86a03f4f18b081b07e058f0e9f96503772a50cf0"))

	ops, err := processing.FileOpsFromCommit(repo, currCommit, prevCommit)
	if err != nil {
		t.Fatalf("failed to get file ops, %s", err)
	}

	expectedOps := []processing.FileOperation{
		processing.FileDelete{"go.mod"},
		processing.FileDelete{"go.sum"},
		processing.FileDelete{"schema.graphql"},
		processing.FileDelete{"tools.go"},
	}

	for i, op := range ops {
		if op != expectedOps[i] {
			t.Errorf("expected = %+v, but got %+v", expectedOps[i], op)
		}
	}
}
*/
