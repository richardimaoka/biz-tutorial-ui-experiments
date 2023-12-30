package input

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

/**
 * ImageRow type(s) and functions
 */
type ImageRow struct {
	RowId         string `json:"rowId"`
	IsTrivial     bool   `json:"isTrivial"`
	Comment       string `json:"comment"`
	ModalContents string `json:"modalContents"`
	ImagePath     string `json:"imagePath"`
	ImageWidth    int    `json:"imageWidth"`
	ImageHeight   int    `json:"imageHeight"`
	ImageCaption  string `json:"imageCaption"`
}

/**
 * Function(s) to convert a row to a more specific row
 */

func toImageRow(fromRow *Row) (*ImageRow, error) {
	errorPrefix := "failed in toImageRow"

	//
	// Check slide type
	//
	slide, err := toSlideType(fromRow.RowType)
	if err != nil || slide != ImageSlide {
		return nil, fmt.Errorf("%s, called for wrong 'rowType' = %s", errorPrefix, fromRow.RowType)
	}

	// image width and height
	height, err := fromRow.ImageHeight.GetSingleValue()
	if err != nil {
		return nil, fmt.Errorf("%s, 'imageHeight' is invalid, %s", errorPrefix, err)
	}

	width, err := fromRow.ImageWidth.GetSingleValue()
	if err != nil {
		return nil, fmt.Errorf("%s, 'imageWidth' is invalid, %s", errorPrefix, err)
	}

	//
	// Check trivial field
	//
	isTrivial := fromRow.Trivial.Value()

	return &ImageRow{
		RowId:         fromRow.RowId,
		IsTrivial:     isTrivial,
		Comment:       fromRow.Comment,
		ModalContents: fromRow.ModalContents,
		ImagePath:     fromRow.FilePath,
		ImageWidth:    width,
		ImageHeight:   height,
		ImageCaption:  fromRow.ImageCaption,
	}, nil
}

/**
 * Function(s) to break down a row to steps
 */
func breakdownImageRow(r *ImageRow, finder *StepIdFinder, prevColumn state.ColumnType) []state.Step {
	// - step creation
	var steps []state.Step

	// cleanup step
	step := imageStep(r, finder)
	steps = append(steps, step)

	return steps
}

/**
 * Function(s) to convert a row to a step
 */
func imageStep(r *ImageRow, StepIdFinder *StepIdFinder) state.Step {
	subId := "ImageStep"
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
			SlideType: state.ImageSlideType,
		},
		ModalFields: state.ModalFields{
			ModalContents: r.ModalContents,
		},
		ImageFields: state.ImageFields{
			ImagePath:    r.ImagePath,
			ImageWidth:   r.ImageWidth,
			ImageHeight:  r.ImageHeight,
			ImageCaption: r.ImageCaption,
		},
	}

	// No tooltip - trivial step and no tooltip to show

	return step
}

func toImageSteps(
	r *Row,
	finder *StepIdFinder,
	prevColumn state.ColumnType,
) ([]state.Step, error) {
	errorPrefix := "failed in toImageSteps"

	specificRow, err := toImageRow(r)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}

	steps := breakdownImageRow(specificRow, finder, prevColumn)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}

	return steps, nil
}
