package state

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/jsonwrap"
)

func Process(repo *git.Repository, stepFile string) error {
	var steps []Step
	err := jsonwrap.Read(stepFile, &steps)
	if err != nil {
		return fmt.Errorf("result.Process failed, %v", err)
	}

	return nil
}

func p(steps []Step) error {
	var page *Page

	for _, step := range steps {
		page.Update(&step)
		// gqlModel := page.ToGraphQL()
		// jsonwrap.WriteJsonToFile()
	}

	return nil
}
