package input

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

/**
 * SectionTitleRow type(s) and functions
 */
type SectionTitleRow struct {
	RowId     string `json:"rowId"`
	IsTrivial bool   `json:"isTrivial"`
	Comment   string `json:"comment"`
	Title     string `json:"title"`
}

/**
 * Function(s) to convert a row to a more specific row
 */

func toSectionTitleRow(fromRow *Row) (*SectionTitleRow, error) {
	errorPrefix := "failed in toSectionTitleRow"

	//
	// Check slide type
	//
	slide, err := toSlideType(fromRow.RowType)
	if err != nil || slide != SectionTitleSlide {
		return nil, fmt.Errorf("%s, called for wrong 'rowType' = %s", errorPrefix, fromRow.RowType)
	}

	//
	// Check instruction field
	//
	if fromRow.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}
	title := fromRow.Instruction

	//
	// Check trivial field
	//
	trivial, err := strToBool(fromRow.Trivial)
	if err != nil {
		return nil, fmt.Errorf("%s, 'trivial' is invalid, %s", errorPrefix, err)
	}

	return &SectionTitleRow{
		RowId:     fromRow.RowId,
		IsTrivial: trivial,
		Comment:   fromRow.Comment,
		Title:     title,
	}, nil
}

/**
 * Function(s) to break down a row to steps
 */
func breakdownSectionTitleRow(r *SectionTitleRow, finder *StepIdFinder, prevColumn state.ColumnType) []state.Step {
	// - step creation
	var steps []state.Step

	// cleanup step
	step := sectionTitleStep(r, finder)
	steps = append(steps, step)

	return steps
}

/**
 * Function(s) to convert a row to a step
 */
func sectionTitleStep(r *SectionTitleRow, StepIdFinder *StepIdFinder) state.Step {
	subId := "sectionTitleStep"
	stepId := StepIdFinder.StepIdFor(r.RowId, subId)

	step := state.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: state.FromRowFields{
			IsFromRow:  true,
			ParentStep: r.RowId,
			SubID:      subId,
		},
		IntrinsicFields: state.IntrinsicFields{
			StepId:    stepId,
			Comment:   r.Comment,
			Mode:      state.SlideshowMode,
			SlideType: state.SectionTitleSlideType,
		},
		SectionTitleFields: state.SectionTitleFields{
			SectionTitle: r.Title,
		},
	}

	// No tooltip - trivial step and no tooltip to show

	return step
}

func toSectionTitleSteps(
	r *Row,
	finder *StepIdFinder,
	prevColumn state.ColumnType,
) ([]state.Step, error) {
	errorPrefix := "failed in toSectionTitleSteps"

	specificRow, err := toSectionTitleRow(r)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}

	steps := breakdownSectionTitleRow(specificRow, finder, prevColumn)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}

	return steps, nil
}
