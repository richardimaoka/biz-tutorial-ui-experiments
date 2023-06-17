package processing_test

import (
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
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

	patch, err := prevCommit.Patch(currCommit)
	if err != nil {
		t.Fatalf("failed to get patch, %s", err)
	}

	flag := false
	if flag {
		t.Errorf("patch message %s", patch.Message())
		for _, v := range patch.FilePatches() {
			from, to := v.Files()
			if from == nil {
				t.Errorf("%s added ", to.Path())
			} else if to == nil {
				t.Errorf("%s removed ", from.Path())
			} else {
				t.Errorf("%s updated ", to.Path())
			}
		}
	}
}
