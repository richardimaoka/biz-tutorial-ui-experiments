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

func (c *SourceColumn) Commit(fields *SourceFields) error {
	if err := c.sourceCode.forwardCommit(fields.Commit); err != nil {
		return err
	}

	c.sourceCode.openFile(fields.DefaultOpenFilePath)

	if fields.SourceTooltipContents != "" {
		if fields.SourceTooltipIsAppend {
			c.sourceCode.appendTooltipContents(fields.SourceTooltipContents)
		} else {
			c.sourceCode.newTooltip(
				fields.SourceTooltipContents,
				SourceCodeTooltipTiming(fields.SourceTooltipTiming),
				fields.SourceTooltipLineNumber,
			)
		}
	}
	return nil
}

func (c *SourceColumn) ShowFileTree() {
}

func (c *SourceColumn) SourceOpen(fields *SourceFields) {
	c.sourceCode.openFile(fields.DefaultOpenFilePath)

	if fields.SourceTooltipContents != "" {
		if fields.SourceTooltipIsAppend {
			c.sourceCode.appendTooltipContents(fields.SourceTooltipContents)
		} else {
			c.sourceCode.newTooltip(
				fields.SourceTooltipContents,
				SourceCodeTooltipTiming(fields.SourceTooltipTiming),
				fields.SourceTooltipLineNumber,
			)
		}
	}
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
		c.SourceOpen(fields)
	case SourceError:
		c.SourceError(fields.DefaultOpenFilePath, fields.SourceTooltipFields)
	case SourceCommit:
		err = c.Commit(fields)
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
