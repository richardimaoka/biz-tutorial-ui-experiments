package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func main() {
	var r *git.Repository
	r, err := git.PlainOpen("/tmp/foo")

	if err != nil {
		fmt.Println(err)

		r, err = git.PlainClone("/tmp/foo", false, &git.CloneOptions{
			URL:      "https://github.com/richardimaoka/gqlgensandbox",
			Progress: os.Stdout,
		})

		if err != nil {
			panic(err)
		}
	}

	commit, _ := r.CommitObject(plumbing.NewHash("91a99d0c0558d2fc03c930d19afa97fc141f0c2e"))

	fmt.Println("commit", commit.NumParents())
}
