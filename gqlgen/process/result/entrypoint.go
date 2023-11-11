package result

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/jsonwrap"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state2"
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
	var terminalColumn *state2.TerminalColumn

	for _, step := range steps {
		/**
		 * Update each column's state
		 */
		switch step.FocusColumn {
		case TerminalColumn:
			// if not initialized, initialize column
			if terminalColumn != nil {
				terminalColumn = state2.NewTerminalColumn()
			}

			toTerminalColumn(terminalColumn, step.StepId, &step.TerminalFields)
		}

		/**
		 * Organize columns into page
		 */
		// columns based on ColumnFields
		// page = page
	}

	return nil
}
