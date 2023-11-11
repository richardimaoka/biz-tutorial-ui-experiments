package state2

import "github.com/go-git/go-git/v5"

/**
 * Source Code Column type and methods
 */

type SourceColumn struct {
	sourceCode *SourceCode
}

func NewSourceColumn(repo *git.Repository, projectDir, tutorial string) *SourceColumn {
	return &SourceColumn{
		sourceCode: NewSourceCode(repo, projectDir, tutorial),
	}
}

func (c *SourceColumn) InitialCommit(commit string) error {
	return nil
}

func (c *SourceColumn) ForwardCommit(nextCommit string) {
}

func (c *SourceColumn) ShowFileTree() {
}

func (c *SourceColumn) OpenFile(filePath string) {
}

func (c *SourceColumn) Update(fields *SourceFields) error {
	return nil
}
