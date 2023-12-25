package input

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

/**
 * TocRow type(s) and functions
 */
type TocRow struct {
	StepId    string `json:"stepId"`
	IsTrivial bool   `json:"isTrivial"`
	Comment   string `json:"comment"`
}

/**
 * Function(s) to convert a row to a more specific row
 */

func toTocRow(fromRow *Row) (*TocRow, error) {
	errorPrefix := "failed in toTocRow"

	//
	// Check slide type
	//
	slide, err := toSlideType(fromRow.RowType)
	if err != nil || slide != TocSlide {
		return nil, fmt.Errorf("%s, called for wrong 'rowType' = %s", errorPrefix, fromRow.RowType)
	}

	//
	// Check trivial field
	//
	trivial, err := strToBool(fromRow.Trivial)
	if err != nil {
		return nil, fmt.Errorf("%s, 'trivial' is invalid, %s", errorPrefix, err)
	}

	return &TocRow{
		StepId:    fromRow.StepId,
		IsTrivial: trivial,
		Comment:   fromRow.Comment,
	}, nil
}

/**
 * Function(s) to break down a row to steps
 */
func breakdownTocRow(r *TocRow, finder *StepIdFinder, prevColumn state.ColumnType) []state.Step {
	// - step creation
	var steps []state.Step

	// cleanup step
	step := TocStep(r, finder)
	steps = append(steps, step)

	return steps
}

/**
 * Function(s) to convert a row to a step
 */
func TocStep(r *TocRow, StepIdFinder *StepIdFinder) state.Step {
	subId := "TocStep"
	stepId := StepIdFinder.StepIdFor(r.StepId, subId)

	step := state.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: state.FromRowFields{
			IsFromRow:  true,
			ParentStep: r.StepId,
			SubID:      subId,
		},
		IntrinsicFields: state.IntrinsicFields{
			StepId:    stepId,
			Comment:   r.Comment,
			Mode:      state.SlideshowMode,
			SlideType: state.TocSlideType,
		},
	}

	// No tooltip - trivial step and no tooltip to show

	return step
}

func toTocSteps(
	r *Row,
	finder *StepIdFinder,
	prevColumn state.ColumnType,
) ([]state.Step, error) {
	errorPrefix := "failed in toTocSteps"

	specificRow, err := toTocRow(r)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}

	steps := breakdownTocRow(specificRow, finder, prevColumn)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}

	return steps, nil
}
