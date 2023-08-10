package state_test

import (
	"fmt"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
)

var storage *memory.Storage
var repoCache map[string]*git.Repository

func gitOpenOrClone(repoUrl string) (*git.Repository, error) {
	if repo, ok := repoCache[repoUrl]; ok {
		return repo, nil
	}

	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: repoUrl})
	if err != nil {
		return nil, fmt.Errorf("cannot clone repo %s, %s", repoUrl, err)
	}

	repoCache[repoUrl] = repo
	return repo, nil
}

func TestMain(m *testing.M) {
	repoCache = make(map[string]*git.Repository)

	fmt.Println("before all tests")
	m.Run()
}
