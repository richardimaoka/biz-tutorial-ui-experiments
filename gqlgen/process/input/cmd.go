package input

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
)

func Run(subArgs []string) error {
	inputCmd := flag.NewFlagSet("input", flag.ExitOnError)
	dirName := inputCmd.String("dir", "", "directory name where input.json is located and steps.json will be written")
	repoUrl := inputCmd.String("repo", "", "GitHub Repository URL of the tutorial")

	if len(subArgs) < 1 {
		writer := inputCmd.Output()
		fmt.Fprintln(writer, "Error - insufficient options. Pass the following options:")
		inputCmd.PrintDefaults()
		os.Exit(1)
	}

	inputCmd.Parse(subArgs)

	inputFile := fmt.Sprintf("%s/input.json", *dirName)
	targetFile := fmt.Sprintf("%s/steps.json", *dirName)

	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: *repoUrl})
	if err != nil {
		return fmt.Errorf("input.Process() failed, %s", err)
	}

	return Process(repo, inputFile, targetFile)
}
