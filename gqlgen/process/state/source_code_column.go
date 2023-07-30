package state

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type SourceCodeColumn struct {
	step string
	sc   *SourceCode
}

func NewSourceCodeColumn(repoUrl, commitStr, initialStep, defaultOpenFilePath, tutorial string) (*SourceCodeColumn, error) {
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: repoUrl})
	if err != nil {
		return nil, fmt.Errorf("failed in NewSourceCodeColumn, could not clone git repo, %s", err)
	}

	sc, err := InitialSourceCode(repo, commitStr, defaultOpenFilePath, tutorial)
	if err != nil {
		return nil, fmt.Errorf("failed in NewSourceCodeColumn, %s", err)
	}

	return &SourceCodeColumn{
		step: initialStep,
		sc:   sc,
	}, nil
}

func (c *SourceCodeColumn) Transition(step, commitStr, defaultOpenFilePath string) error {
	sc, err := NewSourceCode(c.sc.repo, commitStr, c.sc.commit.String(), c.sc.tutorial, step, defaultOpenFilePath)
	if err != nil {
		return fmt.Errorf("failed in Transition, %s", err)
	}

	c.sc = sc
	c.step = step

	return nil
}

func (c *SourceCodeColumn) ToGraphQLSourceCodeColumn() *model.SourceCodeColumn {
	return &model.SourceCodeColumn{
		SourceCode: c.sc.ToGraphQLSourceCode(),
	}
}
