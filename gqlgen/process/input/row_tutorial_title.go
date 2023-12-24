package input

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

/**
 * TutorialTitleRow type(s) and functions
 */
type TutorialTitleRow struct {
	StepId        string `json:"stepId"`
	IsTrivial     bool   `json:"isTrivial"`
	Comment       string `json:"comment"`
	Title         string `json:"title"`
	ModalContents string `json:"modalContents"`
}

/**
 * Function(s) to convert a row to a more specific row
 */

func toTutorialTitleRow(fromRow *Row) (*TutorialTitleRow, error) {
	errorPrefix := "failed in toTerminalCommandRow"

	//
	// Check slide type
	//
	slide, err := toSlideType(fromRow.RowType)
	if err != nil || slide != TutorialTitleSlide {
		return nil, fmt.Errorf("%s, called for wrong 'rowType' = %s", errorPrefix, fromRow.RowType)
	}

	//
	// Check instruction fields
	//
	if fromRow.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}
	tutorialTitle := fromRow.Instruction

	//
	// Check trivial field
	//
	trivial, err := strToBool(fromRow.Trivial)
	if err != nil {
		return nil, fmt.Errorf("%s, 'trivial' is invalid, %s", errorPrefix, err)
	}

	return &TutorialTitleRow{
		StepId:        fromRow.StepId,
		IsTrivial:     trivial,
		Comment:       fromRow.Comment,
		ModalContents: fromRow.ModalContents,
		Title:         tutorialTitle,
	}, nil
}

/**
 * Function(s) to break down a row to steps
 */
func breakdownTutotirlaTitleRow(r *TutorialTitleRow, finder *StepIdFinder, prevColumn state.ColumnType) []state.Step {
	// - step creation
	var steps []state.Step

	// // insert move-to-terminal step if current column != "Terminal", and this is not the very first step
	// if prevColumn != state.TerminalColumnType && prevColumn != state.NoColumnType {
	// 	moveToTerminalStep := moveToTerminalStep(r, finder)
	// 	steps = append(steps, moveToTerminalStep)
	// }

	// if r.Type == TerminalCommand {
	// 	// command step
	// 	cmdStep := terminalCommandStep(r, finder)
	// 	steps = append(steps, cmdStep)

	// 	// cd step
	// 	if strings.HasPrefix(r.Text, "cd ") {
	// 		cmdStep := terminalCdStep(r, finder)
	// 		steps = append(steps, cmdStep)
	// 	}
	// } else if r.Type == TerminalOutput {
	// 	outputStep := terminalOutputStep(r, finder)
	// 	steps = append(steps, outputStep)
	// }

	// // cleanup step
	// step := terminalCleanUpStep(r, finder)
	// steps = append(steps, step)

	return steps
}

func toTutorialTitleSteps(
	r *Row,
	finder *StepIdFinder,
	prevColumn state.ColumnType,
) ([]state.Step, error) {
	// row -> specific row
	tutorialTitleRow, err := toTutorialTitleRow(r)
	if err != nil {
		return nil, fmt.Errorf("toTutorialTitleSteps failed, %s", err)
	}

	// specific row -> step
	steps := breakdownTutotirlaTitleRow(tutorialTitleRow, finder, prevColumn)
	if err != nil {
		return nil, fmt.Errorf("toTutorialTitleSteps failed, %s", err)
	}
	return steps, nil

}
