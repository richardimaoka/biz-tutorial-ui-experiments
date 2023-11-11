package state

import (
	"github.com/go-git/go-git/v5"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

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

func (c *SourceColumn) ToGraphQL() *model.SourceCodeColumn2 {
	return &model.SourceCodeColumn2{
		SourceCode: c.sourceCode.ToGraphQL(),
	}
}

func (c *SourceColumn) ToGraphQLColumnWrapper() *model.ColumnWrapper2 {
	return &model.ColumnWrapper2{
		Column:            c.ToGraphQL(),
		ColumnName:        "SourceCode",
		ColumnDisplayName: stringRef("source"),
	}
}
