package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
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

	head, err := r.Head()
	if err != nil {
		fmt.Println(err)
		return
	}
	ref, err := r.CommitObject(head.Hash())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ref.Hash)
	fmt.Println(ref.Author)
	fmt.Println(ref.Committer)
	fmt.Println(ref.Message)

	fmt.Println("-------------------")

	tree, err := ref.Tree()
	if err != nil {
		fmt.Println(err)
		return
	}

	parent, err := ref.Parent(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	parentTree, err := parent.Tree()
	if err != nil {
		fmt.Println(err)
		return
	}

	diff, err := tree.Diff(parentTree)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("diff\n", diff.String())
	fmt.Println("-------------------------------------")

	fmt.Println("r", r)
}
