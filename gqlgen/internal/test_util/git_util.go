package test_util

import (
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
)

var repoCache map[string]*git.Repository

// thread unsafe
func GitOpenOrClone(t *testing.T, repoUrl string) *git.Repository {
	// if repoCache is nil, then initialize
	if repoCache == nil {
		repoCache = make(map[string]*git.Repository)
	}

	// return from cache, if exists
	if repo, ok := repoCache[repoUrl]; ok {
		return repo
	}

	// here, repo is not in cache, so clone it
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: repoUrl})
	if err != nil {
		t.Fatalf("GitOpenOrClone error, cannot clone repo %s, %s", repoUrl, err)
		return nil
	}

	repoCache[repoUrl] = repo
	return repo
}
