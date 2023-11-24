package state

import (
	"fmt"

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

func (c *SourceColumn) Commit(commitHash, defaultOpenFilePath string) error {
	if err := c.sourceCode.forwardCommit(commitHash); err != nil {
		return err
	}
	c.sourceCode.openFile(defaultOpenFilePath)
	return nil
}

func (c *SourceColumn) ShowFileTree() {
}

func (c *SourceColumn) SourceOpen(filePath string) {
	c.sourceCode.openFile(filePath)
}

func (c *SourceColumn) SourceError(filePath string, tooltip SourceTooltipFields) {
	c.sourceCode.openFile(filePath)
	c.sourceCode.newTooltip(tooltip.SourceTooltipContents, SourceCodeTooltipTiming(tooltip.SourceTooltipTiming), tooltip.SourceTooltipLineNumber)
}

func (c *SourceColumn) Update(fields *SourceFields) error {
	var err error
	switch fields.SourceStepType {
	case FileTree:
		// no update is needed, just changing FocusColumn is fine
	case SourceMove:
		// no update is needed, just changing FocusColumn is fine
	case SourceOpen:
		c.SourceOpen(fields.DefaultOpenFilePath)
	case SourceError:
		c.SourceError(fields.DefaultOpenFilePath, fields.SourceTooltipFields)
	case SourceCommit:
		err = c.Commit(fields.Commit, fields.DefaultOpenFilePath)
	}

	// checi if error happend
	if err != nil {
		return fmt.Errorf("SourceCode Update() failed, %s", err)
	}

	return nil
}

func (c *SourceColumn) ToGraphQL() *model.SourceCodeColumn {
	return &model.SourceCodeColumn{
		SourceCode: c.sourceCode.ToGraphQL(),
	}
}

func (c *SourceColumn) ToGraphQLColumnWrapper() *model.ColumnWrapper {
	return &model.ColumnWrapper{
		Column:            c.ToGraphQL(),
		ColumnName:        "SourceCode",
		ColumnDisplayName: stringRef("source"),
	}
}
