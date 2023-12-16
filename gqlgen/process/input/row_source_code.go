package input

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
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
	StepId          string         `json:"stepId"`
	IsTrivial       bool           `json:"isTrivial"`
	Comment         string         `json:"comment"`
	Commit          string         `json:"commit"`
	Tooltip         *SourceTooltip `json:"tooltip"`
	TypingAnimation bool           `json:"typingAnimation"`
	ShowDiff        bool           `json:"showDiff"`
}

type SourceOpenRow struct {
	StepId    string         `json:"stepId"`
	IsTrivial bool           `json:"isTrivial"`
	Comment   string         `json:"comment"`
	FilePath  string         `json:"filePath"`
	Tooltip   *SourceTooltip `json:"tooltip"`
}

type SourceErrorRow struct {
	StepId    string         `json:"stepId"`
	IsTrivial bool           `json:"isTrivial"`
	Comment   string         `json:"comment"`
	FilePath  string         `json:"filePath"`
	Tooltip   *SourceTooltip `json:"tooltip"`
}

type FileTreeRow struct {
	StepId  string `json:"stepId"`
	Comment string `json:"comment"`
}

func toSourceCommitRow(fromRow *Row) (*SourceCommitRow, error) {
	errorPrefix := "failed to convert to SourceCodeCommit"

	//
	// Check column and type
	//
	column, err := toColumnType(fromRow.Column)
	if err != nil || column != SourceColumn {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.Column)
	}
	subType, err := toSourceCodeSubType(fromRow.Type)
	if err != nil || subType != SourceCommit {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, fromRow.Type)
	}

	//
	// Check instruction fields
	//
	if fromRow.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}
	commit := fromRow.Instruction

	typingAnimation, err := strToBool(fromRow.Instruction2)
	if err != nil {
		return nil, fmt.Errorf("%s, 'instruction2' is invalid, %s", errorPrefix, err)
	}

	showDiff, err := strToBool(fromRow.Instruction3)
	if err != nil {
		return nil, fmt.Errorf("%s, 'instruction3' is invalid, %s", errorPrefix, err)
	}

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
	isTrivial, err := strToBool(fromRow.Trivial)
	if err != nil {
		return nil, fmt.Errorf("%s, 'trivial' is invalid, %s", errorPrefix, err)
	}

	return &SourceCommitRow{
		StepId:          fromRow.StepId,
		IsTrivial:       isTrivial,
		Comment:         fromRow.Comment,
		Commit:          commit,
		Tooltip:         tooltip,
		TypingAnimation: typingAnimation,
		ShowDiff:        showDiff,
	}, nil
}

func toSourceOpenRow(fromRow *Row) (*SourceOpenRow, error) {
	errorPrefix := "failed to convert to SourceCodeOpen"

	//
	// Check column and type
	//
	column, err := toColumnType(fromRow.Column)
	if err != nil || column != SourceColumn {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.Column)
	}
	subType, err := toSourceCodeSubType(fromRow.Type)
	if err != nil || subType != SourceOpen {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, fromRow.Type)
	}

	//
	// Check instruction
	//
	if fromRow.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}
	filePath := fromRow.Instruction

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
	isTrivial, err := strToBool(fromRow.Trivial)
	if err != nil {
		return nil, fmt.Errorf("%s, 'trivial' is invalid, %s", errorPrefix, err)
	}

	return &SourceOpenRow{
		StepId:    fromRow.StepId,
		IsTrivial: isTrivial,
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
	column, err := toColumnType(fromRow.Column)
	if err != nil || column != SourceColumn {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.Column)
	}
	subType, err := toSourceCodeSubType(fromRow.Type)
	if err != nil || subType != SourceError {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, fromRow.Type)
	}

	//
	// Check instruction
	//
	if fromRow.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}
	filepath := fromRow.Instruction

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
	isTrivial, err := strToBool(fromRow.Trivial)
	if err != nil {
		return nil, fmt.Errorf("%s, 'trivial' is invalid, %s", errorPrefix, err)
	}

	return &SourceErrorRow{
		StepId:    fromRow.StepId,
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
	column, err := toColumnType(fromRow.Column)
	if err != nil || column != SourceColumn {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.Column)
	}
	subType, err := toSourceCodeSubType(fromRow.Type)
	if err != nil || subType != FileTree {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, fromRow.Type)
	}

	return &FileTreeRow{
		StepId:  fromRow.StepId,
		Comment: fromRow.Comment,
	}, nil
}

/**
 * Function(s) to convert a row to a step
 */
func fileTreeStep(r *FileTreeRow, StepIdFinder *StepIdFinder, usedColumns UsedColumns) state.Step {
	subId := "fileTreeStep"
	stepId := StepIdFinder.StepIdFor(r.StepId, subId)

	step := state.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: state.FromRowFields{
			IsFromRow:  true,
			ParentStep: r.StepId,
			SubID:      subId,
		},
		IntrinsicFields: state.IntrinsicFields{
			StepId:  stepId,
			Comment: "(file tree)",
		},
		AnimationFields: state.AnimationFields{
			IsTrivial: true,
		},
		// No ModalFields, as it is a trivial step
		ColumnFields: resultColumns(state.SourceColumnType, usedColumns),
		SourceFields: state.SourceFields{
			SourceStepType: state.FileTree,
		},
	}

	// No tooltip - trivial step and no tooltip to show

	return step
}

func openFileStep(r *SourceOpenRow, StepIdFinder *StepIdFinder, usedColumns UsedColumns, filePath string) state.Step {
	subId := "openFileStep"
	stepId := StepIdFinder.StepIdFor(r.StepId, subId)

	step := state.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: state.FromRowFields{
			IsFromRow:  true,
			ParentStep: r.StepId,
			SubID:      subId,
		},
		IntrinsicFields: state.IntrinsicFields{
			StepId:  stepId,
			Comment: r.Comment,
		},
		AnimationFields: state.AnimationFields{
			IsTrivial: r.IsTrivial,
		},
		// No ModalFields, as it is a trivial step
		ColumnFields: resultColumns(state.SourceColumnType, usedColumns),
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

func sourceCommitStep(r *SourceCommitRow, StepIdFinder *StepIdFinder, usedColumns UsedColumns, filePath string) state.Step {
	subId := fmt.Sprintf("sourceCommitStep-%s", filePath)
	stepId := StepIdFinder.StepIdFor(r.StepId, subId)

	step := state.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: state.FromRowFields{
			IsFromRow:  true,
			ParentStep: r.StepId,
			SubID:      subId,
		},
		IntrinsicFields: state.IntrinsicFields{
			StepId:  stepId,
			Comment: r.Comment,
		},
		AnimationFields: state.AnimationFields{
			IsTrivial: r.IsTrivial,
		},
		// No ModalFields, as it is a trivial step
		ColumnFields: resultColumns(state.SourceColumnType, usedColumns),
		SourceFields: state.SourceFields{
			SourceStepType:      state.SourceCommit,
			Commit:              r.Commit,
			DefaultOpenFilePath: filePath,
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

func openSourceErrorStep(r *SourceErrorRow, StepIdFinder *StepIdFinder, usedColumns UsedColumns, filePath string) state.Step {
	subId := "openSourceErrorStep"
	stepId := StepIdFinder.StepIdFor(r.StepId, subId)

	step := state.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: state.FromRowFields{
			IsFromRow:  true,
			ParentStep: r.StepId,
			SubID:      subId,
		},
		IntrinsicFields: state.IntrinsicFields{
			StepId:  stepId,
			Comment: r.Comment,
		},
		AnimationFields: state.AnimationFields{
			IsTrivial: r.IsTrivial,
		},
		// No ModalFields, as it is a trivial step
		ColumnFields: resultColumns(state.SourceColumnType, usedColumns),
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

func moveToSourceCodeStep(parentStepId string, StepIdFinder *StepIdFinder, usedColumns UsedColumns) state.Step {
	subId := fmt.Sprintf("moveToSourceCodeStep")
	stepId := StepIdFinder.StepIdFor(parentStepId, subId)

	step := state.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: state.FromRowFields{
			IsFromRow:  true,
			ParentStep: parentStepId,
			SubID:      subId,
		},
		IntrinsicFields: state.IntrinsicFields{
			StepId:  stepId,
			Comment: ("(move to source code)"),
		},
		AnimationFields: state.AnimationFields{
			IsTrivial: true,
		},
		// No ModalFields, as it is a trivial step
		ColumnFields: resultColumns(state.SourceColumnType, usedColumns),
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
	prevColumns *ColumnInfo,
	repo *git.Repository,
	prevCommit string,
) ([]state.Step, error) {
	// - step creation
	var steps []state.Step

	// insert move-to-terminal step if current column != "Source Code", and this is not the very first step
	if prevColumns.Focus != state.SourceColumnType && prevColumns.Focus != state.NoColumnType {
		step := moveToSourceCodeStep(r.StepId, finder, prevColumns.AllUsed)
		steps = append(steps, step)
	}

	// find files from commit
	filePaths, err := filesBetweenCommits(repo, prevCommit, r.Commit)
	if err != nil {
		return nil, fmt.Errorf("breakdownSourceCommitRow failed, %s", err)
	}

	// open file steps
	for i, filePath := range filePaths {
		step := sourceCommitStep(r, finder, prevColumns.AllUsed, filePath)
		steps = append(steps, step)
		if i == 5 {
			break
		}
	}

	return steps, nil
}

func breakdownSourceOpenRow(
	r *SourceOpenRow,
	finder *StepIdFinder,
	prevColumns *ColumnInfo,
	repo *git.Repository,
) ([]state.Step, error) {
	// - step creation
	var steps []state.Step

	// insert move-to-terminal step if current column != "Source Code", and this is not the very first step
	if prevColumns.Focus != state.SourceColumnType && prevColumns.Focus != state.NoColumnType {
		step := moveToSourceCodeStep(r.StepId, finder, prevColumns.AllUsed)
		steps = append(steps, step)
	}

	// open file step
	step := openFileStep(r, finder, prevColumns.AllUsed, r.FilePath)
	steps = append(steps, step)

	return steps, nil
}

func breakdownSourceErrorRow(
	r *SourceErrorRow,
	finder *StepIdFinder,
	prevColumns *ColumnInfo,
	repo *git.Repository,
) ([]state.Step, error) {
	// - step creation
	var steps []state.Step

	// insert move-to-terminal step if current column != "Source Code", and this is not the very first step
	if prevColumns.Focus != state.SourceColumnType && prevColumns.Focus != state.NoColumnType {
		step := moveToSourceCodeStep(r.StepId, finder, prevColumns.AllUsed)
		steps = append(steps, step)
	}

	// open file step
	step := openSourceErrorStep(r, finder, prevColumns.AllUsed, r.FilePath)
	steps = append(steps, step)

	return steps, nil
}

func breakdownFileTreeRow(
	r *FileTreeRow,
	finder *StepIdFinder,
	prevColumns *ColumnInfo,
	repo *git.Repository,
) ([]state.Step, error) {
	// - step creation
	var steps []state.Step

	// insert move-to-terminal step if current column != "Source Code", and this is not the very first step
	if prevColumns.Focus != state.SourceColumnType && prevColumns.Focus != state.NoColumnType {
		step := moveToSourceCodeStep(r.StepId, finder, prevColumns.AllUsed)
		steps = append(steps, step)
	}

	// open file step
	step := fileTreeStep(r, finder, prevColumns.AllUsed)
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
	prevColumns *ColumnInfo,
	repo *git.Repository,
	prevCommit string,
) ([]state.Step, *ColumnInfo, error) {
	// current columns update
	currentColumns := &ColumnInfo{
		AllUsed: appendIfNotExists(prevColumns.AllUsed, state.SourceColumnType),
		Focus:   state.SourceColumnType,
	}

	subType, err := toSourceCodeSubType(r.Type)
	if err != nil {
		return nil, nil, fmt.Errorf("toSourceSteps failed, %s", err)
	}

	switch subType {
	case SourceCommit:
		// row -> specific row
		row, err := toSourceCommitRow(r)
		if err != nil {
			return nil, nil, fmt.Errorf("toSourceSteps failed, %s", err)
		}

		// specific row -> step
		steps, err := breakdownSourceCommitRow(row, finder, prevColumns, repo, prevCommit)
		if err != nil {
			return nil, nil, fmt.Errorf("toSourceSteps failed, %s", err)
		}
		return steps, currentColumns, nil

	case SourceOpen:
		// row -> specific row
		row, err := toSourceOpenRow(r)
		if err != nil {
			return nil, nil, fmt.Errorf("toSourceSteps failed, %s", err)
		}

		// specific row -> step
		steps, err := breakdownSourceOpenRow(row, finder, prevColumns, repo)
		if err != nil {
			return nil, nil, fmt.Errorf("toSourceSteps failed, %s", err)
		}
		return steps, currentColumns, nil

	case SourceError:
		// row -> specific row
		row, err := toSourceErrorRow(r)
		if err != nil {
			return nil, nil, fmt.Errorf("toSourceSteps failed, %s", err)
		}

		// specific row -> step
		steps, err := breakdownSourceErrorRow(row, finder, prevColumns, repo)
		if err != nil {
			return nil, nil, fmt.Errorf("toSourceSteps failed, %s", err)
		}
		return steps, currentColumns, nil

	case FileTree:
		// row -> specific row
		row, err := toFileTreeRow(r)
		if err != nil {
			return nil, nil, fmt.Errorf("toSourceSteps failed, %s", err)
		}

		// specific row -> step
		steps, err := breakdownFileTreeRow(row, finder, prevColumns, repo)
		if err != nil {
			return nil, nil, fmt.Errorf("toSourceSteps failed, %s", err)
		}
		return steps, currentColumns, nil

	default:
		return nil, nil, fmt.Errorf("toSourceSteps failed, type = '%s' not implemented", r.Type)
	}
}
