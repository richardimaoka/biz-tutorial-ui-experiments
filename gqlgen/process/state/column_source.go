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
	var err error

	// process commit
	err = c.sourceCode.forwardCommit(fields.Commit)
	if err != nil {
		return fmt.Errorf("Commit() failed, %s", err)
	}

	// open file
	c.sourceCode.openFile(fields.DefaultOpenFilePath)

	// tooltip
	if fields.SourceTooltipContents != "" {
		if fields.SourceTooltipIsAppend {
			err = c.sourceCode.appendTooltipContents(fields.SourceTooltipContents)
		} else {
			err = c.sourceCode.newTooltip(
				fields.DefaultOpenFilePath,
				fields.SourceTooltipContents,
				SourceCodeTooltipTiming(fields.SourceTooltipTiming),
				fields.SourceTooltipLineNumber,
			)
		}
	}
	if err != nil {
		return fmt.Errorf("Commit() failed, %s", err)
	}

	return nil
}

func (c *SourceColumn) ShowFileTree() {
}

func (c *SourceColumn) SourceOpen(fields *SourceFields) error {
	// open file
	c.sourceCode.openFile(fields.DefaultOpenFilePath)

	// tooltip
	var err error
	if fields.SourceTooltipContents != "" {
		if fields.SourceTooltipIsAppend {
			err = c.sourceCode.appendTooltipContents(fields.SourceTooltipContents)
		} else {
			err = c.sourceCode.newTooltip(
				fields.DefaultOpenFilePath,
				fields.SourceTooltipContents,
				SourceCodeTooltipTiming(fields.SourceTooltipTiming),
				fields.SourceTooltipLineNumber,
			)
		}
	}
	if err != nil {
		return fmt.Errorf("SourceOpen() failed, %s", err)
	}

	return nil
}

func (c *SourceColumn) SourceError(fields *SourceFields) error {
	// open file
	c.sourceCode.openFile(fields.DefaultOpenFilePath)

	// tooltip
	err := c.sourceCode.newTooltip(
		fields.DefaultOpenFilePath,
		fields.SourceTooltipContents,
		SourceCodeTooltipTiming(fields.SourceTooltipTiming),
		fields.SourceTooltipLineNumber,
	)
	if err != nil {
		return fmt.Errorf("SourceError() failed, %s", err)
	}

	return nil

}

func (c *SourceColumn) CleanUp(fields *SourceFields) error {
	err := c.sourceCode.ClearTooltip()
	if err != nil {
		return fmt.Errorf("CleanUp() failed, %s", err)
	}

	return nil
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
		err = c.SourceError(fields)
	case SourceCommit:
		err = c.Commit(fields)
	case SourceCleanUp:
		err = c.CleanUp(fields)
	}

	// check if error happend
	if err != nil {
		return fmt.Errorf("SourceColumn Update() failed, %s", err)
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
