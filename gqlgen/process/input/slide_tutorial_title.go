package input

import (
	"fmt"
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/csvfield"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

/**
 * TutorialTitleRow type(s) and functions
 */
type TutorialTitleRow struct {
	RowId         string               `json:"rowId"`
	IsTrivial     bool                 `json:"isTrivial"`
	Comment       string               `json:"comment"`
	Title         string               `json:"title"`
	ImagePaths    string               `json:"imagePaths"`
	ImageWidths   csvfield.CsvMultiInt `json:"imageWidths"`
	ImageHeights  csvfield.CsvMultiInt `json:"imageHeights"`
	ImageCaptions string               `json:"imageCaptions"`
	ModalContents string               `json:"modalContents"`
}

/**
 * Function(s) to convert a row to a more specific row
 */

func toTutorialTitleRow(fromRow *Row) (*TutorialTitleRow, error) {
	errorPrefix := "failed in toTutorialTitleRow"

	//
	// Check slide type
	//
	slide, err := toSlideType(fromRow.RowType)
	if err != nil || slide != TutorialTitleSlide {
		return nil, fmt.Errorf("%s, called for wrong 'rowType' = %s", errorPrefix, fromRow.RowType)
	}

	//
	// Check instruction field
	//
	if fromRow.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}
	tutorialTitle := fromRow.Instruction

	//
	// Check file and image fields
	//
	if fromRow.FilePath != "" {
		delimiter := "\n"
		files := strings.Split(fromRow.FilePath, delimiter)
		imageCaptions := strings.Split(fromRow.ImageCaption, delimiter)

		sameLength := len(files) == fromRow.ImageWidths.Length() &&
			fromRow.ImageWidths.Length() == fromRow.ImageHeights.Length() &&
			fromRow.ImageHeights.Length() == len(imageCaptions)

		if !sameLength {
			return nil, fmt.Errorf("%s, length of filePaths, imageSizes, imageWidths, imageCaptions got different", errorPrefix)
		}
	}

	//
	// Check trivial field
	//
	trivial, err := strToBool(fromRow.Trivial)
	if err != nil {
		return nil, fmt.Errorf("%s, 'trivial' is invalid, %s", errorPrefix, err)
	}

	return &TutorialTitleRow{
		RowId:         fromRow.RowId,
		IsTrivial:     trivial,
		Comment:       fromRow.Comment,
		Title:         tutorialTitle,
		ImagePaths:    fromRow.FilePath,
		ImageWidths:   fromRow.ImageWidths,
		ImageHeights:  fromRow.ImageHeights,
		ImageCaptions: fromRow.ImageCaption,
	}, nil
}

/**
 * Function(s) to break down a row to steps
 */
func breakdownTutotirlaTitleRow(r *TutorialTitleRow, finder *StepIdFinder, prevColumn state.ColumnType) []state.Step {
	// - step creation
	var steps []state.Step

	// cleanup step
	step := tutorialTitleStep(r, finder)
	steps = append(steps, step)

	return steps
}

/**
 * Function(s) to convert a row to a step
 */
func tutorialTitleStep(r *TutorialTitleRow, StepIdFinder *StepIdFinder) state.Step {
	subId := "tutorialTitleStep"
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
			SlideType: state.TutorialTitleSlideType,
		},
		TutorialTitleFields: state.TutorialTitleFields{
			TutorialTitle:           r.Title,
			TutorialTitleImagePaths: r.ImagePaths,
			// TutorialTitleImageSizes:    r.ImageSizes,
			TutorialTitleImageCaptions: r.ImageCaptions,
		},
	}

	// No tooltip - trivial step and no tooltip to show

	return step
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
