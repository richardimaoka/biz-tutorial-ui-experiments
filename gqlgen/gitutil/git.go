package gitutil

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
)

var repoCache map[string]*git.Repository

// thread unsafe
func GitOpenOrClone(repoUrl string) (*git.Repository, error) {
	// if repoCache is nil, then initialize
	if repoCache == nil {
		repoCache = make(map[string]*git.Repository)
	}

	// return from cache, if exists
	if repo, ok := repoCache[repoUrl]; ok {
		return repo, nil
	}

	// here, repo is not in cache, so clone it
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: repoUrl})
	if err != nil {
		return nil, fmt.Errorf("GitOpenOrClone error, cannot clone repo %s, %s", repoUrl, err)
	}

	repoCache[repoUrl] = repo
	return repo, nil
}
