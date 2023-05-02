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

	itr, err := r.CommitObjects()
	for {
		ref, err := itr.Next()
		if err != nil {
			break
		}
		fmt.Println(ref.Hash)
		fmt.Println(ref.Author)
		fmt.Println(ref.Committer)
		fmt.Println(ref.Message)

		fileItr, err := ref.Files()
		if err != nil {
			break
		}
		for {
			file, err := fileItr.Next()
			if err != nil {
				break
			}
			fmt.Println(file.Name)
		}
	}

	fmt.Println("r", r)
}
