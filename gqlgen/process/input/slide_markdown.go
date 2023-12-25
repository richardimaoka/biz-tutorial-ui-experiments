package input

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

/**
 * MarkdownRow type(s) and functions
 */
type MarkdownRow struct {
	RowId            string `json:"rowId"`
	IsTrivial        bool   `json:"isTrivial"`
	Comment          string `json:"comment"`
	MarkdownContents string `json:"markdownContents"`
}

/**
 * Function(s) to convert a row to a more specific row
 */

func toMarkdownRow(fromRow *Row) (*MarkdownRow, error) {
	errorPrefix := "failed in toMarkdownRow"

	//
	// Check slide type
	//
	slide, err := toSlideType(fromRow.RowType)
	if err != nil || slide != MarkdownSlide {
		return nil, fmt.Errorf("%s, called for wrong 'rowType' = %s", errorPrefix, fromRow.RowType)
	}

	//
	// Check instruction field
	//
	if fromRow.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}
	markdownContents := fromRow.Instruction

	//
	// Check trivial field
	//
	trivial, err := strToBool(fromRow.Trivial)
	if err != nil {
		return nil, fmt.Errorf("%s, 'trivial' is invalid, %s", errorPrefix, err)
	}

	return &MarkdownRow{
		RowId:            fromRow.RowId,
		IsTrivial:        trivial,
		Comment:          fromRow.Comment,
		MarkdownContents: markdownContents,
	}, nil
}

/**
 * Function(s) to break down a row to steps
 */
func breakdownMarkdownRow(r *MarkdownRow, finder *StepIdFinder, prevColumn state.ColumnType) []state.Step {
	// - step creation
	var steps []state.Step

	// cleanup step
	step := markdownStep(r, finder)
	steps = append(steps, step)

	return steps
}

/**
 * Function(s) to convert a row to a step
 */
func markdownStep(r *MarkdownRow, StepIdFinder *StepIdFinder) state.Step {
	subId := "markdownStep"
	stepId := StepIdFinder.StepIdFor(r.RowId, subId)

	step := state.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: state.FromRowFields{
			IsFromRow:   true,
			ParentRowId: r.RowId,
			SubID:       subId,
		},
		IntrinsicFields: state.IntrinsicFields{
			StepId:    stepId,
			Comment:   r.Comment,
			Mode:      state.SlideshowMode,
			SlideType: state.MarkdownSlideType,
		},

		MarkdownFields: state.MarkdownFields{
			MarkdownContents: r.MarkdownContents,
		},
	}

	// No tooltip - trivial step and no tooltip to show

	return step
}

func toMarkdownSteps(
	r *Row,
	finder *StepIdFinder,
	prevColumn state.ColumnType,
) ([]state.Step, error) {
	errorPrefix := "failed in func toMarkdownSteps"

	specificRow, err := toMarkdownRow(r)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}

	steps := breakdownMarkdownRow(specificRow, finder, prevColumn)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}

	return steps, nil
}
