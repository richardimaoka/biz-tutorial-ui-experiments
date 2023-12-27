package input

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

type ImageFileSuffix = string

const (
	JPG  ImageFileSuffix = "jpg"
	JPEG ImageFileSuffix = "jpeg"
	GIF  ImageFileSuffix = "gif"
	PNG  ImageFileSuffix = "png"
)

func toImageFileSuffix(s string) (ImageFileSuffix, error) {
	switch strings.ToLower(s) {
	case JPG:
		return JPG, nil
	case JPEG:
		return JPEG, nil
	case GIF:
		return GIF, nil
	case PNG:
		return PNG, nil
	case "":
		return PNG, nil
	default:
		return "", fmt.Errorf("ImageFileSuffix value = '%s' is invalid", s)
	}
}

/**
 * BrowserSubType type(s) and functions
 */
type BrowserSubType string

const (
	// Lower cases since they are from manual entries
	BrowserSingle   BrowserSubType = "single"
	BrowserNumSeq   BrowserSubType = "numseq"
	BrowserSequence BrowserSubType = "sequence"
)

func toBrowserSubType(s string) (BrowserSubType, error) {
	lower := strings.ToLower(s)

	switch lower {
	case string(BrowserSingle):
		return BrowserSingle, nil
	case string(BrowserNumSeq):
		return BrowserNumSeq, nil
	case string(BrowserSequence):
		return BrowserSequence, nil
	default:
		return "", fmt.Errorf("'%s' is an invalid browser sub type", s)
	}
}

/**
 * Browser row type(s) and functions
 */

type BrowserRow struct {
	RowId          string   `json:"rowId"`
	IsTrivial      bool     `json:"isTrivial"`
	Comment        string   `json:"comment"`
	ModalContents  string   `json:"modalContents"`
	ImageFileNames []string `json:"imageFileNames"`
}

type BrowserSingleRow struct {
	RowId           string `json:"rowId"`
	IsTrivial       bool   `json:"isTrivial"`
	Comment         string `json:"comment"`
	ModalContents   string `json:"modalContents"`
	ImageFilePath   string `json:"imageFilePath"`
	ImageFileWidth  int    `json:"imageFileWidth"`
	ImageFileHeight int    `json:"imageFileHeight"`
}

var BrowserNumSeqPattern *regexp.Regexp = regexp.MustCompile(`\[[0-9]+\]`)

type BrowserNumSeqRow struct {
	StepId          string `json:"rowId"`
	IsTrivial       bool   `json:"isTrivial"`
	Comment         string `json:"comment"`
	ModalContents   string `json:"modalContents"`
	ImageBaseName   string `json:"imageFileBaseName"`
	ImageFileSuffix string `json:"imageFileSuffix"`
	NumImages       int    `json:"numImages"`
}

type BrowserSequenceRow struct {
	StepId         string   `json:"rowId"`
	IsTrivial      bool     `json:"isTrivial"`
	Comment        string   `json:"comment"`
	ModalContents  string   `json:"modalContents"`
	ImageFileNames []string `json:"imageFileNames"`
}

func toBrowserSingleRow(fromRow *Row) (*BrowserSingleRow, error) {
	errorPrefix := "failed in toBrowserSingleRow"

	//
	// Check column and type
	//
	column, err := toColumnType(fromRow.RowType)
	if err != nil || column != BrowserColumn {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.RowType)
	}
	subType, err := toBrowserSubType(fromRow.SubType)
	if err != nil || subType != BrowserSingle {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, fromRow.SubType)
	}

	//
	// Check filepath
	//
	if fromRow.FilePath == "" {
		return nil, fmt.Errorf("%s, 'filePath' is empty", errorPrefix)
	}
	imageFileName := fromRow.FilePath

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
	trivial, err := strToBool(fromRow.Trivial)
	if err != nil {
		return nil, fmt.Errorf("%s, 'trivial' is invalid, %s", errorPrefix, err)
	}

	return &BrowserSingleRow{
		RowId:           fromRow.RowId,
		IsTrivial:       trivial,
		Comment:         fromRow.Comment,
		ModalContents:   fromRow.ModalContents,
		ImageFilePath:   imageFileName,
		ImageFileWidth:  width,
		ImageFileHeight: height,
	}, nil
}

func toBrowserNumSeqRow(fromRow *Row) (*BrowserRow, error) {
	errorPrefix := "failed in toBrowserNumSeqRow"

	//
	// Check column and type
	//
	column, err := toColumnType(fromRow.RowType)
	if err != nil || column != BrowserColumn {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.RowType)
	}
	subType, err := toBrowserSubType(fromRow.SubType)
	if err != nil || subType != BrowserNumSeq {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, fromRow.SubType)
	}

	//
	// Check instruction
	//
	if fromRow.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}

	// extract num from (e.g.) filename[30].png
	numFiles, err := numInSqBracket(fromRow.Instruction)
	if err != nil {
		return nil, fmt.Errorf("%s, 'instruction' is in wrong form, %s", errorPrefix, err)
	}

	baseName, err := fileBaseName(fromRow.Instruction)
	if err != nil {
		return nil, fmt.Errorf("%s, 'instruction' is invalid, %s", errorPrefix, err)
	}
	baseName = removeSqBrackets(baseName)

	suffix, err := fileSuffix(fromRow.Instruction, fromRow)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}

	var imageFileNames []string
	for i := 0; i < numFiles; i++ {
		imageFileNames = append(imageFileNames, fmt.Sprintf("%s-%03d.%s", baseName, i, suffix))
	}

	//
	// Check trivial field
	//
	trivial, err := strToBool(fromRow.Trivial)
	if err != nil {
		return nil, fmt.Errorf("%s, 'trivial' is invalid, %s", errorPrefix, err)
	}

	return &BrowserRow{
		RowId:          fromRow.RowId,
		IsTrivial:      trivial,
		Comment:        fromRow.Comment,
		ImageFileNames: imageFileNames,
	}, nil
}

func toBrowserSequenceRow(fromRow *Row) (*BrowserRow, error) {
	errorPrefix := "failed in toBrowserSequenceRow"

	//
	// Check column and type
	//
	column, err := toColumnType(fromRow.RowType)
	if err != nil || column != BrowserColumn {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.RowType)
	}
	subType, err := toBrowserSubType(fromRow.SubType)
	if err != nil || subType != BrowserSequence {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, fromRow.SubType)
	}

	//
	// Check instruction
	//
	if fromRow.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}

	splitByComma := strings.Split(fromRow.Instruction, ",")
	var imgFiles []string
	for _, v := range splitByComma {
		baseName, err := fileBaseName(v)
		if err != nil {
			return nil, fmt.Errorf("%s, 'instruction' is invalid, %s", errorPrefix, err)
		}
		suffix, err := fileSuffix(v, fromRow)
		if err != nil {
			return nil, fmt.Errorf("%s, %s", errorPrefix, err)
		}
		imgFiles = append(imgFiles, fmt.Sprintf("%s.%s", baseName, suffix))
	}

	//
	// Check trivial field
	//
	trivial, err := strToBool(fromRow.Trivial)
	if err != nil {
		return nil, fmt.Errorf("%s, 'trivial' is invalid, %s", errorPrefix, err)
	}

	return &BrowserRow{
		RowId:          fromRow.RowId,
		IsTrivial:      trivial,
		Comment:        fromRow.Comment,
		ImageFileNames: imgFiles,
	}, nil
}

func removeSqBrackets(s string) string {
	split := strings.Split(s, "[")
	return split[0]
}

// Get file base name from given file-name candidate
func fileBaseName(s string) (string, error) {
	if s == "" {
		return "", fmt.Errorf("file name is empty")
	}

	splitByDot := strings.Split(s, ".")
	if len(splitByDot) == 1 {
		return s, nil
	} else if len(splitByDot) == 2 {
		return splitByDot[0], nil
	} else {
		return "", fmt.Errorf("file name has too many dots")
	}
}

// Get file suffix from 1) given file-name candidate, 2) if not there, try finding from Row
func fileSuffix(fileNameCandidate string, fromRow *Row) (string, error) {
	splitByDot := strings.Split(fileNameCandidate, ".")
	if len(splitByDot) == 1 {
		// if 'instruction' doesn't have a '.', then the suffix must be in 'instruction2'
		suffix, err := toImageFileSuffix(fromRow.Instruction2)
		if err != nil {
			return "", fmt.Errorf("file name = '%s' has a wrong suffix, %s", fileNameCandidate, err)
		}
		return suffix, nil
	} else if len(splitByDot) == 2 {
		// if 'instruction' has a '.', then the suffix follows that .
		suffix, err := toImageFileSuffix(splitByDot[1])
		if err != nil {
			return "", fmt.Errorf("'instruction2' has a wrong suffix, %s", err)
		}
		return suffix, nil
	} else {
		return "", fmt.Errorf("file name = '%s' has too many dots", fileNameCandidate)
	}
}

// extract number from square bracket
// s to be form of (e.g.) 'filename[42]'
func numInSqBracket(s string) (int, error) {
	found := BrowserNumSeqPattern.FindString(s)
	if found == "" {
		return 0, fmt.Errorf("the string doesn't have the form '[${num}]'")
	}

	// found = (e.g.) '[42]'
	num, err := strconv.Atoi(found[1 : len(found)-1]) // remove '[' and ']' from 'found'
	if err != nil {
		return 0, err
	}
	if num < 1 {
		return 0, fmt.Errorf("number in '[${num}]' is a negative number = %d", num)
	}

	return num, nil
}

/**
 * Function(s) to convert a row to a step
 */
func openBrowserStep(r *BrowserSingleRow, StepIdFinder *StepIdFinder) state.Step {
	subId := "openBrowserStep"
	stepId := StepIdFinder.StepIdFor(r.RowId, subId)

	step := state.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: state.FromRowFields{
			IsFromRow:   true,
			ParentRowId: r.RowId,
			SubID:       subId,
		},
		IntrinsicFields: state.IntrinsicFields{
			StepId:      stepId,
			Comment:     r.Comment,
			Mode:        state.HandsonMode,
			FocusColumn: state.BrowserColumnType,
		},
		AnimationFields: state.AnimationFields{
			IsTrivial: r.IsTrivial,
		},
		ModalFields: state.ModalFields{
			ModalContents: r.ModalContents,
		},
		BrowserFields: state.BrowserFields{
			BrowserStepType:    state.BrowserOpen,
			BrowserImagePath:   r.ImageFilePath,
			BrowserImageWidth:  r.ImageFileWidth,
			BrowserImageHeight: r.ImageFileHeight,
		},
	}

	return step
}

func openBrowserStepN(r *BrowserRow, StepIdFinder *StepIdFinder, nthFile int) state.Step {
	subId := fmt.Sprintf("openBrowserStep-%03d", nthFile)
	stepId := StepIdFinder.StepIdFor(r.RowId, subId)

	imageFileName := r.ImageFileNames[nthFile]

	step := state.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: state.FromRowFields{
			IsFromRow:   true,
			ParentRowId: r.RowId,
			SubID:       subId,
		},
		IntrinsicFields: state.IntrinsicFields{
			StepId:      stepId,
			Comment:     r.Comment,
			Mode:        state.HandsonMode,
			FocusColumn: state.BrowserColumnType,
		},
		AnimationFields: state.AnimationFields{
			IsTrivial: r.IsTrivial,
		},
		ModalFields: state.ModalFields{
			ModalContents: r.ModalContents,
		},
		BrowserFields: state.BrowserFields{
			BrowserStepType:  state.BrowserOpen,
			BrowserImagePath: imageFileName,
		},
	}

	return step
}

func moveToBrowserStep(rowId string, finder *StepIdFinder) state.Step {
	subId := "moveToBrowserStep"
	stepId := finder.StepIdFor(rowId, subId)

	step := state.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: state.FromRowFields{
			IsFromRow:   true,
			ParentRowId: rowId,
			SubID:       subId,
		},
		IntrinsicFields: state.IntrinsicFields{
			StepId:      stepId,
			Comment:     "(move to Browser)",
			Mode:        state.HandsonMode,
			FocusColumn: state.BrowserColumnType,
		},
		AnimationFields: state.AnimationFields{
			IsTrivial: true, //always true
		},
		// No ModalFields, as it is a trivial step
		BrowserFields: state.BrowserFields{
			BrowserStepType: state.BrowserMove,
		},
	}
	// No tooltip - move step should be trivial and no tooltip to show

	return step
}

/**
 * Function(s) to break down a row to steps
 */
func breakdownBrowserRow(
	r *BrowserSingleRow,
	finder *StepIdFinder,
	prevColumn state.ColumnType,
) ([]state.Step, error) {
	// - step creation
	var steps []state.Step

	// insert move-to-terminal step if current column != "Browser", and this is not the very first step
	if prevColumn != state.BrowserColumnType && prevColumn != state.NoColumnType {
		step := moveToBrowserStep(r.RowId, finder)
		steps = append(steps, step)
	}

	// open browser step
	step := openBrowserStep(r, finder)
	steps = append(steps, step)

	return steps, nil
}

/**
 * Function to turn a row into steps
 */

func toBrowserSteps(
	r *Row,
	finder *StepIdFinder,
	prevColumn state.ColumnType,
) ([]state.Step, error) {
	subType, err := toBrowserSubType(r.SubType)
	if err != nil {
		return nil, fmt.Errorf("toBrowserSubType failed, %s", err)
	}

	switch subType {
	case BrowserSingle:
		// row -> specific row
		row, err := toBrowserSingleRow(r)
		if err != nil {
			return nil, fmt.Errorf("toBrowserSteps failed, %s", err)
		}

		// specific row -> step
		steps, err := breakdownBrowserRow(row, finder, prevColumn)
		if err != nil {
			return nil, fmt.Errorf("toBrowserSteps failed, %s", err)
		}
		return steps, nil

	// case BrowserNumSeq:
	// 	// row -> specific row
	// 	row, err := toBrowserNumSeqRow(r)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("toBrowserSteps failed, %s", err)
	// 	}

	// 	// specific row -> step
	// 	steps, err := breakdownBrowserRow(row, finder, prevColumn)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("toBrowserSteps failed, %s", err)
	// 	}
	// 	return steps, nil

	// case BrowserSequence:
	// 	// row -> specific row
	// 	row, err := toBrowserSequenceRow(r)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("toBrowserSteps failed, %s", err)
	// 	}

	// 	// specific row -> step
	// 	steps, err := breakdownBrowserRow(row, finder, prevColumn)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("toBrowserSteps failed, %s", err)
	// 	}
	// 	return steps, nil

	default:
		return nil, fmt.Errorf("toBrowserSteps failed, type = '%s' not implemented", r.SubType)
	}

}
