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
	errorPrefix := fmt.Sprintf("Commit() for %s failed", fields.Commit)

	// guess file path
	var filePath string
	if fields.DefaultOpenFilePath == "" {
		var err error
		filePath, err = c.sourceCode.openFileBestGuess(fields.Commit)
		if err != nil {
			return fmt.Errorf("%s, %s", errorPrefix, err)
		}
	} else {
		filePath = fields.DefaultOpenFilePath
	}

	// process commit
	err := c.sourceCode.forwardCommit(fields.Commit)
	if err != nil {
		return fmt.Errorf("%s, %s", errorPrefix, err)
	}

	// open file
	c.sourceCode.openFile(filePath)

	// tooltip
	if fields.SourceTooltipContents != "" {
		if fields.SourceTooltipIsAppend {
			err := c.sourceCode.appendTooltipContents(fields.SourceTooltipContents)
			if err != nil {
				return fmt.Errorf("%s, %s", errorPrefix, err)
			}
		} else {
			err := c.sourceCode.newTooltip(
				filePath,
				fields.SourceTooltipContents,
				SourceCodeTooltipTiming(fields.SourceTooltipTiming),
				fields.SourceTooltipLineNumber,
			)
			if err != nil {
				return fmt.Errorf("%s, %s", errorPrefix, err)
			}
		}
	}

	return nil
}

func (c *SourceColumn) FileTree(fields *SourceFields) error {
	c.sourceCode.showFileTree = true
	return nil
}

func (c *SourceColumn) SourceOpen(fields *SourceFields) error {
	funcName := "SourceOpen()"

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
		return fmt.Errorf("%s failed, %s", funcName, err)
	}

	return nil
}

func (c *SourceColumn) SourceError(fields *SourceFields) error {
	funcName := "SourceError()"

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
		return fmt.Errorf("%s failed, %s", funcName, err)
	}

	return nil

}

func (c *SourceColumn) CleanUpPrevStep() error {
	funcName := "CleanUpPrevStep()"

	if err := c.sourceCode.clearTooltip(); err != nil {
		return fmt.Errorf("%s failed, %s", funcName, err)
	}

	return nil
}

func (c *SourceColumn) Update(fields *SourceFields) error {
	errorPrefix := "Update() failed"

	err := c.sourceCode.clearEdits()
	if err != nil {
		return fmt.Errorf("%s, %s", errorPrefix, err)
	}

	err = c.sourceCode.closeFileTree()
	if err != nil {
		return fmt.Errorf("%s, %s", errorPrefix, err)
	}

	switch fields.SourceStepType {
	case FileTree:
		err = c.FileTree(fields)
	case SourceMove:
		// no update is needed, just changing FocusColumn is fine
	case SourceOpen:
		err = c.SourceOpen(fields)
	case SourceError:
		err = c.SourceError(fields)
	case SourceCommit:
		err = c.Commit(fields)
	default:
		err = fmt.Errorf("soruce step type = '%s' is not implemented yet", fields.SourceStepType)
	}
	// check if error happend
	if err != nil {
		return fmt.Errorf("%s failed, %s", errorPrefix, err)
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
