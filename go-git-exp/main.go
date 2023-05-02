package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

func main() {
	var r *git.Repository
	r, err := git.PlainOpen("/tmp/foxo")

	if err != nil {
		fmt.Println(err)

		r, err = git.PlainClone("/tmp/foo2", false, &git.CloneOptions{
			URL:      "https://github.com/go-git/go-git",
			Progress: os.Stdout,
		})

		if err != nil {
			panic(err)
		}
	}

	fmt.Println("r", r)
}
