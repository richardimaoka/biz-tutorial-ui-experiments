package main

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
)

func main() {
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/richardimaoka/gqlgensandbox",
	})
	if err != nil {
		log.Fatalf("error cloning repo: %v", err)
	}

	t1, err := repo.TreeObject(plumbing.NewHash("30419cb45f05d32ce9c8b273fe1a61b7a2c05c26"))
	if err != nil {
		panic(fmt.Sprintf("t1:%s", err))
	}
	iter := t1.Files()
	for {
		file, err := iter.Next()
		if err != nil {
			break
		}

		fmt.Println("file", file.Name)
	}

	t2, err := repo.TreeObject(plumbing.NewHash("8ac76d64b959641b11dbad47517f4873e71954e0"))
	if err != nil {
		panic(fmt.Sprintf("t2:%s", err))
	}
	diff, err := t1.Diff(t2)
	if err != nil {
		panic(fmt.Sprintf("diff:%s", err))
	}

	fmt.Println("diff", diff)

	commit, _ := repo.CommitObject(plumbing.NewHash("91a99d0c0558d2fc03c930d19afa97fc141f0c2e"))
	fmt.Println("commit", commit.NumParents())
}
