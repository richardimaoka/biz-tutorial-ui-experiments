package input

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/csvfield"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/gitwrap"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

/**
 * SourceCodeSubType type(s) and functions
 */
type SourceCodeSubType string

const (
	// Lower cases since they are from manual entries
	SourceCommit SourceCodeSubType = "commit"
	SourceOpen   SourceCodeSubType = "open"
	SourceError  SourceCodeSubType = "source error"
	FileTree     SourceCodeSubType = "filetree"
)

func toSourceCodeSubType(s string) (SourceCodeSubType, error) {
	lower := strings.ToLower(s)

	switch lower {
	case string(SourceCommit):
		return SourceCommit, nil
	case string(SourceOpen):
		return SourceOpen, nil
	case string(SourceError):
		return SourceError, nil
	case string(FileTree):
		return FileTree, nil
	default:
		return "", fmt.Errorf("'%s' is an invalid source code sub type", s)
	}
}

/**
 * SourceTooltip type(s) and functions
 */

type SourceTooltipTiming string

const (
	SOURCE_TOOLTIP_START SourceTooltipTiming = "START"
	SOURCE_TOOLTIP_END   SourceTooltipTiming = "END"
)

type SourceTooltip struct {
	Contents   string              `json:"contents"`
	LineNumber int                 `json:"lineNumber"`
	Timing     SourceTooltipTiming `json:"timing"`
	// TODO: if IsAppend == true, Timing must be START.
	// So, probably the timing doesn't need to be controlled from outside?
	IsAppend bool `json:"isAppend"`
}

func toSourceTooltipTiming(s string) (SourceTooltipTiming, error) {
	switch strings.ToUpper(s) {
	case string(SOURCE_TOOLTIP_START):
		return SOURCE_TOOLTIP_START, nil
	case string(SOURCE_TOOLTIP_END):
		return SOURCE_TOOLTIP_END, nil
	case "": // default value is different from termianl tooltip
		return SOURCE_TOOLTIP_END, nil
	default:
		return "", fmt.Errorf("SourceTooltipTiming value = '%s' is invalid", s)
	}
}

func (t SourceTooltipTiming) toState() state.SourceCodeTooltipTiming {
	switch t {
	case SOURCE_TOOLTIP_START:
		return state.SOURCE_TOOLTIP_START
	case SOURCE_TOOLTIP_END:
		return state.SOURCE_TOOLTIP_END
	default:
		panic(fmt.Sprintf("SourceToolTipTiming has an invalid value = '%s'", t))
	}
}

func toSourceTooltip(fromRow *Row) (*SourceTooltip, error) {
	if fromRow.Tooltip == "" {
		return nil, nil
	}

	contents := fromRow.Tooltip

	tooltipTiming, err := toSourceTooltipTiming(fromRow.TooltipTiming)
	if err != nil {
		return nil, fmt.Errorf("'tooltipTiming' field is wrong, %s", err)
	}

	tooltipLine, err := fromRow.TooltipLine.GetInt()
	if err != nil {
		return nil, fmt.Errorf("'tooltipLine' is invalid, %s", err)
	}

	var isAppend bool
	isAppendFromRow := strings.ToUpper(fromRow.TooltipAppend)
	if isAppendFromRow == "TRUE" {
		isAppend = true
	} else if isAppendFromRow == "FALSE" {
		isAppend = false
	} else if isAppendFromRow == "" {
		isAppend = false
	} else {
		return nil, fmt.Errorf("'tooltipAppend' = '%s', is an invalid value. It has to be either 'TRUE', 'FALSE' or ''(empty)", fromRow.TooltipAppend)
	}

	return &SourceTooltip{
		Contents:   contents,
		Timing:     tooltipTiming,
		LineNumber: tooltipLine,
		IsAppend:   isAppend,
	}, nil
}

/**
 * Source row type(s) and functions
 */
type SourceCommitRow struct {
	RowId               string         `json:"rowId"`
	IsTrivial           bool           `json:"isTrivial"`
	Comment             string         `json:"comment"`
	Commit              string         `json:"commit"`
	DefaultOpenFilePath string         `json:"defaultOpenFilePath"`
	Tooltip             *SourceTooltip `json:"tooltip"`
	TypingAnimation     bool           `json:"typingAnimation"`
}

type SourceOpenRow struct {
	RowId     string         `json:"rowId"`
	IsTrivial bool           `json:"isTrivial"`
	Comment   string         `json:"comment"`
	FilePath  string         `json:"filePath"`
	Tooltip   *SourceTooltip `json:"tooltip"`
}

type SourceErrorRow struct {
	RowId     string         `json:"rowId"`
	IsTrivial bool           `json:"isTrivial"`
	Comment   string         `json:"comment"`
	FilePath  string         `json:"filePath"`
	Tooltip   *SourceTooltip `json:"tooltip"`
}

type FileTreeRow struct {
	RowId     string        `json:"rowId"`
	IsTrivial csvfield.Bool `json:"isTrivial"`
	Comment   string        `json:"comment"`
}

func toSourceCommitRow(fromRow *Row) (*SourceCommitRow, error) {
	errorPrefix := "failed to convert to SourceCodeCommit"

	//
	// Check column and type
	//
	column, err := toColumnType(fromRow.RowType)
	if err != nil || column != SourceColumn {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.RowType)
	}
	subType, err := toSourceCodeSubType(fromRow.SubType)
	if err != nil || subType != SourceCommit {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, fromRow.SubType)
	}

	//
	// Check contents fields
	//
	if fromRow.Contents == "" {
		return nil, fmt.Errorf("%s, 'contents' is empty", errorPrefix)
	}
	commit := fromRow.Contents

	defaultOpenFilePath := fromRow.FilePath
	typingAnimation := bool(fromRow.TypingAnimation)

	//
	// Check tooltip field
	//
	tooltip, err := toSourceTooltip(fromRow)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}

	//
	// Check isTrivial field
	//
	isTrivial := fromRow.Trivial.Value()

	return &SourceCommitRow{
		RowId:               fromRow.RowId,
		IsTrivial:           isTrivial,
		Comment:             fromRow.Comment,
		Commit:              commit,
		DefaultOpenFilePath: defaultOpenFilePath,
		Tooltip:             tooltip,
		TypingAnimation:     typingAnimation,
	}, nil
}

func toSourceOpenRow(fromRow *Row) (*SourceOpenRow, error) {
	errorPrefix := "failed to convert to SourceCodeOpen"

	//
	// Check column and type
	//
	column, err := toColumnType(fromRow.RowType)
	if err != nil || column != SourceColumn {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.RowType)
	}
	subType, err := toSourceCodeSubType(fromRow.SubType)
	if err != nil || subType != SourceOpen {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, fromRow.SubType)
	}

	//
	// Check contents
	//
	if fromRow.Contents == "" {
		return nil, fmt.Errorf("%s, 'contents' is empty", errorPrefix)
	}
	filePath := fromRow.Contents

	//
	// Check tooltip fields
	//
	tooltip, err := toSourceTooltip(fromRow)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}

	//
	// Check isTrivial field
	//
	trivial := fromRow.Trivial.Value()

	return &SourceOpenRow{
		RowId:     fromRow.RowId,
		IsTrivial: trivial,
		Comment:   fromRow.Comment,
		FilePath:  filePath,
		Tooltip:   tooltip,
	}, nil
}

func toSourceErrorRow(fromRow *Row) (*SourceErrorRow, error) {
	errorPrefix := "failed to convert to SourceCodeError"

	//
	// Check column and type
	//
	column, err := toColumnType(fromRow.RowType)
	if err != nil || column != SourceColumn {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.RowType)
	}
	subType, err := toSourceCodeSubType(fromRow.SubType)
	if err != nil || subType != SourceError {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, fromRow.SubType)
	}

	//
	// Check contents
	//
	if fromRow.Contents == "" {
		return nil, fmt.Errorf("%s, 'contents' is empty", errorPrefix)
	}
	filepath := fromRow.Contents

	//
	// Check tooltip fields
	//
	tooltip, err := toSourceTooltip(fromRow)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}
	if tooltip == nil {
		return nil, fmt.Errorf("%s, source code error needs the error detail in 'tooltip'", errorPrefix)
	}

	//
	// Check trivial field
	//
	isTrivial := fromRow.Trivial.Value()

	return &SourceErrorRow{
		RowId:     fromRow.RowId,
		IsTrivial: isTrivial,
		Comment:   fromRow.Comment,
		FilePath:  filepath,
		Tooltip:   tooltip,
	}, nil
}

func toFileTreeRow(fromRow *Row) (*FileTreeRow, error) {
	errorPrefix := "failed to convert to FileTreeRow"

	//
	// Check column and type
	//
	column, err := toColumnType(fromRow.RowType)
	if err != nil || column != SourceColumn {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.RowType)
	}
	subType, err := toSourceCodeSubType(fromRow.SubType)
	if err != nil || subType != FileTree {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, fromRow.SubType)
	}

	return &FileTreeRow{
		RowId:     fromRow.RowId,
		Comment:   fromRow.Comment,
		IsTrivial: fromRow.Trivial,
	}, nil
}

/**
 * Function(s) to convert a row to a step
 */
func fileTreeStep(r *FileTreeRow, StepIdFinder *StepIdFinder) state.Step {
	subId := "fileTreeStep"
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
			Comment:     "(file tree)",
			Mode:        state.HandsonMode,
			FocusColumn: state.SourceColumnType,
		},
		AnimationFields: state.AnimationFields{
			IsTrivial: true,
		},
		// No ModalFields, as it is a trivial step
		SourceFields: state.SourceFields{
			SourceStepType: state.FileTree,
		},
	}

	// No tooltip - trivial step and no tooltip to show

	return step
}

func openFileStep(r *SourceOpenRow, StepIdFinder *StepIdFinder, filePath string) state.Step {
	subId := "openFileStep"
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
			FocusColumn: state.SourceColumnType,
		},
		AnimationFields: state.AnimationFields{
			IsTrivial: r.IsTrivial,
		},
		// No ModalFields, as it is a trivial step
		SourceFields: state.SourceFields{
			SourceStepType:      state.SourceOpen,
			DefaultOpenFilePath: filePath,
		},
	}
	if r.Tooltip != nil {
		step.SourceTooltipContents = r.Tooltip.Contents
		step.SourceTooltipTiming = r.Tooltip.Timing.toState()
		step.SourceTooltipLineNumber = r.Tooltip.LineNumber
		step.SourceTooltipIsAppend = r.Tooltip.IsAppend
	}

	return step
}

func sourceCommitStep(r *SourceCommitRow, StepIdFinder *StepIdFinder) state.Step {
	subId := "sourceCommitStep"
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
			FocusColumn: state.SourceColumnType,
		},
		AnimationFields: state.AnimationFields{
			IsTrivial: r.IsTrivial,
		},
		// No ModalFields, as it is a trivial step
		SourceFields: state.SourceFields{
			SourceStepType:      state.SourceCommit,
			Commit:              r.Commit,
			DefaultOpenFilePath: r.DefaultOpenFilePath,
			TypingAnimation:     true,
		},
	}
	if r.Tooltip != nil {
		step.SourceTooltipContents = r.Tooltip.Contents
		step.SourceTooltipTiming = r.Tooltip.Timing.toState()
		step.SourceTooltipLineNumber = r.Tooltip.LineNumber
		step.SourceTooltipIsAppend = r.Tooltip.IsAppend
	}

	return step
}

func openSourceErrorStep(r *SourceErrorRow, StepIdFinder *StepIdFinder, filePath string) state.Step {
	subId := "openSourceErrorStep"
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
			FocusColumn: state.SourceColumnType,
		},
		AnimationFields: state.AnimationFields{
			IsTrivial: r.IsTrivial,
		},
		// No ModalFields, as it is a trivial step
		SourceFields: state.SourceFields{
			SourceStepType:      state.SourceMove,
			DefaultOpenFilePath: filePath,
		},
	}
	if r.Tooltip != nil {
		step.SourceTooltipContents = r.Tooltip.Contents
		step.SourceTooltipTiming = r.Tooltip.Timing.toState()
		step.SourceTooltipLineNumber = r.Tooltip.LineNumber
		step.SourceTooltipIsAppend = r.Tooltip.IsAppend
	}

	return step
}

func moveToSourceCodeStep(parentRowId string, StepIdFinder *StepIdFinder) state.Step {
	subId := fmt.Sprintf("moveToSourceCodeStep")
	stepId := StepIdFinder.StepIdFor(parentRowId, subId)

	step := state.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: state.FromRowFields{
			IsFromRow:   true,
			ParentRowId: parentRowId,
			SubID:       subId,
		},
		IntrinsicFields: state.IntrinsicFields{
			StepId:      stepId,
			Comment:     ("(move to source code)"),
			Mode:        state.HandsonMode,
			FocusColumn: state.SourceColumnType,
		},
		AnimationFields: state.AnimationFields{
			IsTrivial: true,
		},
		// No ModalFields, as it is a trivial step
		SourceFields: state.SourceFields{
			SourceStepType: state.SourceMove,
		},
	}

	// No tooltip - trivial step and no tooltip to show

	return step
}

/**
 * Function(s) to break down a row to steps
 */
func breakdownSourceCommitRow(
	r *SourceCommitRow,
	finder *StepIdFinder,
	prevColumn state.ColumnType,
) ([]state.Step, error) {
	// - step creation
	var steps []state.Step

	// insert move-to-terminal step if current column != "Source Code", and this is not the very first step
	if prevColumn != state.SourceColumnType && prevColumn != state.NoColumnType {
		step := moveToSourceCodeStep(r.RowId, finder)
		steps = append(steps, step)
	}

	step := sourceCommitStep(r, finder)
	steps = append(steps, step)

	// // find files from commit
	// filePaths, err := filesBetweenCommits(repo, prevCommit, r.Commit)
	// if err != nil {
	// 	return nil, fmt.Errorf("breakdownSourceCommitRow failed, %s", err)
	// }

	// // open file steps
	// for i, filePath := range filePaths {
	// 	step := sourceCommitStep(r, finder, prevColumns.AllUsed, filePath)
	// 	steps = append(steps, step)
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	return steps, nil
}

func breakdownSourceOpenRow(
	r *SourceOpenRow,
	finder *StepIdFinder,
	prevColumn state.ColumnType,
) ([]state.Step, error) {
	// - step creation
	var steps []state.Step

	// insert move-to-terminal step if current column != "Source Code", and this is not the very first step
	if prevColumn != state.SourceColumnType && prevColumn != state.NoColumnType {
		step := moveToSourceCodeStep(r.RowId, finder)
		steps = append(steps, step)
	}

	// open file step
	step := openFileStep(r, finder, r.FilePath)
	steps = append(steps, step)

	return steps, nil
}

func breakdownSourceErrorRow(
	r *SourceErrorRow,
	finder *StepIdFinder,
	prevColumn state.ColumnType,
) ([]state.Step, error) {
	// - step creation
	var steps []state.Step

	// insert move-to-terminal step if current column != "Source Code", and this is not the very first step
	if prevColumn != state.SourceColumnType && prevColumn != state.NoColumnType {
		step := moveToSourceCodeStep(r.RowId, finder)
		steps = append(steps, step)
	}

	// open file step
	step := openSourceErrorStep(r, finder, r.FilePath)
	steps = append(steps, step)

	return steps, nil
}

func breakdownFileTreeRow(
	r *FileTreeRow,
	finder *StepIdFinder,
	prevColumn state.ColumnType,
) ([]state.Step, error) {
	// - step creation
	var steps []state.Step

	// insert move-to-terminal step if current column != "Source Code", and this is not the very first step
	if prevColumn != state.SourceColumnType && prevColumn != state.NoColumnType {
		step := moveToSourceCodeStep(r.RowId, finder)
		steps = append(steps, step)
	}

	// open file step
	step := fileTreeStep(r, finder)
	steps = append(steps, step)

	return steps, nil
}

/**
 * Helper function
 */
func filesBetweenCommits(repo *git.Repository, fromCommit, toCommit string) ([]string, error) {
	if fromCommit == "" {
		// initial commit
		files, err := gitwrap.GetCommitFiles(repo, toCommit)
		if err != nil {
			return nil, err
		}

		var filePaths []string
		for _, v := range files {
			filePaths = append(filePaths, v.Name)
		}

		return filePaths, nil
	} else {
		// non-initial commit, get added/updated/renamed files
		files, err := gitwrap.GetPatchFiles(repo, fromCommit, toCommit)
		if err != nil {
			return nil, err
		}

		var filePaths []string
		for _, v := range files {
			filePaths = append(filePaths, v.Path())
		}

		return filePaths, nil
	}
}

/**
 * Function to turn a row into steps
 */
func toSourceSteps(
	r *Row,
	finder *StepIdFinder,
	prevColumn state.ColumnType,
) ([]state.Step, error) {
	subType, err := toSourceCodeSubType(r.SubType)
	if err != nil {
		return nil, fmt.Errorf("toSourceSteps failed, %s", err)
	}

	switch subType {
	case SourceCommit:
		// row -> specific row
		row, err := toSourceCommitRow(r)
		if err != nil {
			return nil, fmt.Errorf("toSourceSteps failed, %s", err)
		}

		// specific row -> step
		steps, err := breakdownSourceCommitRow(row, finder, prevColumn)
		if err != nil {
			return nil, fmt.Errorf("toSourceSteps failed, %s", err)
		}
		return steps, nil

	case SourceOpen:
		// row -> specific row
		row, err := toSourceOpenRow(r)
		if err != nil {
			return nil, fmt.Errorf("toSourceSteps failed, %s", err)
		}

		// specific row -> step
		steps, err := breakdownSourceOpenRow(row, finder, prevColumn)
		if err != nil {
			return nil, fmt.Errorf("toSourceSteps failed, %s", err)
		}
		return steps, nil

	case SourceError:
		// row -> specific row
		row, err := toSourceErrorRow(r)
		if err != nil {
			return nil, fmt.Errorf("toSourceSteps failed, %s", err)
		}

		// specific row -> step
		steps, err := breakdownSourceErrorRow(row, finder, prevColumn)
		if err != nil {
			return nil, fmt.Errorf("toSourceSteps failed, %s", err)
		}
		return steps, nil

	case FileTree:
		// row -> specific row
		row, err := toFileTreeRow(r)
		if err != nil {
			return nil, fmt.Errorf("toSourceSteps failed, %s", err)
		}

		// specific row -> step
		steps, err := breakdownFileTreeRow(row, finder, prevColumn)
		if err != nil {
			return nil, fmt.Errorf("toSourceSteps failed, %s", err)
		}
		return steps, nil

	default:
		return nil, fmt.Errorf("toSourceSteps failed, type = '%s' not implemented", r.SubType)
	}
}
