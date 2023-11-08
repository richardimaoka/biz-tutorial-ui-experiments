package input

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/gitwrap"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/result"
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
	default:
		return "", fmt.Errorf("'%s' is an invalid source code sub type", s)
	}
}

/**
 * SourceTooltip type(s) and functions
 */
type SourceTooltip struct {
	Contents   string        `json:"contents"`
	LineNumber int           `json:"lineNumber"`
	Timing     TooltipTiming `json:"timing"`
	// TODO: if IsAppend == true, Timing must be START.
	// So, probably the timing doesn't need to be controlled from outside?
	IsAppend bool `json:"isAppend"`
}

func toSourceTooltipTiming(s string) (TooltipTiming, error) {
	switch strings.ToUpper(s) {
	case START:
		return START, nil
	case END:
		return END, nil
	case "": // default value is different from termianl tooltip
		return END, nil
	default:
		return "", fmt.Errorf("TooltipTiming value = '%s' is invalid", s)
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

	if fromRow.TooltipLine == 0 {
		return nil, fmt.Errorf("'tooltipLine' = %d, cannot be 0 nor empty", fromRow.TooltipLine)
	} else if fromRow.TooltipLine < 0 {
		return nil, fmt.Errorf("'tooltipLine' = %d, but cannot be a negative number", fromRow.TooltipLine)
	}

	var isAppend bool
	isAppendFromRow := strings.ToUpper(fromRow.TooltipAppend)
	if isAppendFromRow == "TRUE" {
		isAppend = true
	} else if isAppendFromRow == "FALSE" {
		isAppend = false
	} else {
		return nil, fmt.Errorf("'tooltipAppend' = '%s', is an invalid value. It has to be either 'TRUE', 'FALSE' or empty", fromRow.TooltipAppend)
	}

	return &SourceTooltip{
		Contents:   contents,
		Timing:     tooltipTiming,
		LineNumber: fromRow.TooltipLine,
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

func toSourceCommitRow(fromRow *Row) (*SourceCommitRow, error) {
	errorPrefix := "failed to convert to SourceCodeCommit"

	//
	// Check column and type
	//
	if strings.ToLower(fromRow.Column) != "source" {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.Column)
	}
	if strings.ToLower(fromRow.Type) != string(SourceCommit) {
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
	if strings.ToLower(fromRow.Column) != "source" {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.Column)
	}
	if strings.ToLower(fromRow.Type) != string(SourceOpen) {
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
	if strings.ToLower(fromRow.Column) != "source" {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.Column)
	}
	if strings.ToLower(fromRow.Type) != string(SourceError) {
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

/**
 * Function(s) to convert a row to a step
 */
func fileTreeStep(r *SourceCommitRow, StepIdFinder *StepIdFinder, usedColumns UsedColumns) result.Step {
	subId := "fileTreeStep"
	stepId := StepIdFinder.StepIdFor(r.StepId, subId)

	step := result.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: result.FromRowFields{
			IsFromRow:  true,
			ParentStep: r.StepId,
			SubID:      subId,
		},
		IntrinsicFields: result.IntrinsicFields{
			StepId:  stepId,
			Comment: "(file tree)",
		},
		AnimationFields: result.AnimationFields{
			IsTrivial: true,
		},
		// No ModalFields, as it is a trivial step
		ColumnFields: resultColumns(result.SourceColumn, usedColumns),
		SourceCodeFields: result.SourceCodeFields{
			Commit:       r.Commit,
			ShowFileTree: true,
		},
	}

	// No tooltip - trivial step and no tooltip to show

	return step
}

func openFileStep(r *SourceOpenRow, StepIdFinder *StepIdFinder, usedColumns UsedColumns, filePath string) result.Step {
	subId := "openFileStep"
	stepId := StepIdFinder.StepIdFor(r.StepId, subId)

	step := result.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: result.FromRowFields{
			IsFromRow:  true,
			ParentStep: r.StepId,
			SubID:      subId,
		},
		IntrinsicFields: result.IntrinsicFields{
			StepId:  stepId,
			Comment: r.Comment,
		},
		AnimationFields: result.AnimationFields{
			IsTrivial: r.IsTrivial,
		},
		// No ModalFields, as it is a trivial step
		ColumnFields: resultColumns(result.SourceColumn, usedColumns),
		SourceCodeFields: result.SourceCodeFields{
			DefaultOpenFilePath: filePath,
			ShowFileTree:        false,
		},
	}
	if r.Tooltip != nil {
		step.SourceCodeTooltipContents = r.Tooltip.Contents
		step.SourceCodeTooltipTiming = r.Tooltip.Timing
		step.SourceCodeTooltipLineNumber = r.Tooltip.LineNumber
		step.SourceCodeTooltipIsAppend = r.Tooltip.IsAppend
	}

	return step
}

func openFileCommitStep(r *SourceCommitRow, StepIdFinder *StepIdFinder, usedColumns UsedColumns, filePath string) result.Step {
	subId := fmt.Sprintf("openFileCommitStep-%s", filePath)
	stepId := StepIdFinder.StepIdFor(r.StepId, subId)

	step := result.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: result.FromRowFields{
			IsFromRow:  true,
			ParentStep: r.StepId,
			SubID:      subId,
		},
		IntrinsicFields: result.IntrinsicFields{
			StepId:  stepId,
			Comment: r.Comment,
		},
		AnimationFields: result.AnimationFields{
			IsTrivial: r.IsTrivial,
		},
		// No ModalFields, as it is a trivial step
		ColumnFields: resultColumns(result.SourceColumn, usedColumns),
		SourceCodeFields: result.SourceCodeFields{
			Commit:              r.Commit,
			DefaultOpenFilePath: filePath,
			ShowFileTree:        false,
			TypingAnimation:     r.TypingAnimation,
		},
	}
	if r.Tooltip != nil {
		step.SourceCodeTooltipContents = r.Tooltip.Contents
		step.SourceCodeTooltipTiming = r.Tooltip.Timing
		step.SourceCodeTooltipLineNumber = r.Tooltip.LineNumber
		step.SourceCodeTooltipIsAppend = r.Tooltip.IsAppend
	}

	return step
}

func openSourceErrorStep(r *SourceErrorRow, StepIdFinder *StepIdFinder, usedColumns UsedColumns, filePath string) result.Step {
	subId := "openSourceErrorStep"
	stepId := StepIdFinder.StepIdFor(r.StepId, subId)

	step := result.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: result.FromRowFields{
			IsFromRow:  true,
			ParentStep: r.StepId,
			SubID:      subId,
		},
		IntrinsicFields: result.IntrinsicFields{
			StepId:  stepId,
			Comment: r.Comment,
		},
		AnimationFields: result.AnimationFields{
			IsTrivial: r.IsTrivial,
		},
		// No ModalFields, as it is a trivial step
		ColumnFields: resultColumns(result.SourceColumn, usedColumns),
		SourceCodeFields: result.SourceCodeFields{
			DefaultOpenFilePath: filePath,
			ShowFileTree:        false,
		},
	}
	if r.Tooltip != nil {
		step.SourceCodeTooltipContents = r.Tooltip.Contents
		step.SourceCodeTooltipTiming = r.Tooltip.Timing
		step.SourceCodeTooltipLineNumber = r.Tooltip.LineNumber
		step.SourceCodeTooltipIsAppend = r.Tooltip.IsAppend
	}

	return step
}

func moveToSourceCodeStep(parentStepId string, StepIdFinder *StepIdFinder, usedColumns UsedColumns) result.Step {
	subId := fmt.Sprintf("moveToSourceCodeStep")
	stepId := StepIdFinder.StepIdFor(parentStepId, subId)

	step := result.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: result.FromRowFields{
			IsFromRow:  true,
			ParentStep: parentStepId,
			SubID:      subId,
		},
		IntrinsicFields: result.IntrinsicFields{
			StepId:  stepId,
			Comment: ("(move to source code)"),
		},
		AnimationFields: result.AnimationFields{
			IsTrivial: true,
		},
		// No ModalFields, as it is a trivial step
		ColumnFields: resultColumns(result.SourceColumn, usedColumns),
	}

	// No tooltip - trivial step and no tooltip to show

	return step
}

func filesBetweenCommits(repo *git.Repository, fromCommit, toCommit string) ([]string, error) {
	if fromCommit == "" {
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
 * Function(s) to break down a row to steps
 */
func breakdownSourceCommitRow(
	r *SourceCommitRow,
	finder *StepIdFinder,
	prevColumns ColumnInfo,
	repo *git.Repository,
	prevCommit string,
) ([]result.Step, ColumnInfo, error) {
	// - step creation
	var steps []result.Step

	// insert move-to-terminal step if current column != "Source Code", and this is not the very first step
	if prevColumns.Focus != result.SourceColumn && prevColumns.Focus != result.NoColumn {
		step := moveToSourceCodeStep(r.StepId, finder, prevColumns.AllUsed)
		steps = append(steps, step)
	}

	// find files from commit
	filePaths, err := filesBetweenCommits(repo, r.Commit, prevCommit)
	if err != nil {
		return nil, ColumnInfo{}, fmt.Errorf("breakdownSourceCommitRow faield %s", err)
	}

	// open file steps
	for i, filePath := range filePaths {
		step := openFileCommitStep(r, finder, prevColumns.AllUsed, filePath)
		steps = append(steps, step)
		if i == 5 {
			break
		}
	}

	currentColumns := ColumnInfo{
		AllUsed: appendIfNotExists(prevColumns.AllUsed, result.TerminalColumn),
		Focus:   result.SourceColumn,
	}

	return steps, currentColumns, nil
}

func breakdownSourceOpenRow(
	r *SourceOpenRow,
	finder *StepIdFinder,
	prevColumns ColumnInfo,
	repo *git.Repository,
	prevCommit string,
) ([]result.Step, ColumnInfo, error) {
	// - step creation
	var steps []result.Step

	// insert move-to-terminal step if current column != "Source Code", and this is not the very first step
	if prevColumns.Focus != result.SourceColumn && prevColumns.Focus != result.NoColumn {
		step := moveToSourceCodeStep(r.StepId, finder, prevColumns.AllUsed)
		steps = append(steps, step)
	}

	// open file step
	step := openFileStep(r, finder, prevColumns.AllUsed, r.FilePath)
	steps = append(steps, step)

	currentColumns := ColumnInfo{
		AllUsed: appendIfNotExists(prevColumns.AllUsed, result.TerminalColumn),
		Focus:   result.SourceColumn,
	}

	return steps, currentColumns, nil
}

func breakdownSourceErrorRow(
	r *SourceErrorRow,
	finder *StepIdFinder,
	prevColumns ColumnInfo,
	repo *git.Repository,
	prevCommit string,
) ([]result.Step, ColumnInfo, error) {
	// - step creation
	var steps []result.Step

	// insert move-to-terminal step if current column != "Source Code", and this is not the very first step
	if prevColumns.Focus != result.SourceColumn && prevColumns.Focus != result.NoColumn {
		step := moveToSourceCodeStep(r.StepId, finder, prevColumns.AllUsed)
		steps = append(steps, step)
	}

	// open file step
	step := openSourceErrorStep(r, finder, prevColumns.AllUsed, r.FilePath)
	steps = append(steps, step)

	currentColumns := ColumnInfo{
		AllUsed: appendIfNotExists(prevColumns.AllUsed, result.TerminalColumn),
		Focus:   result.SourceColumn,
	}

	return steps, currentColumns, nil
}

/**
 * Function to turn a row into steps
 */
func toSourceSteps(
	r *Row,
	finder *StepIdFinder,
	prevColumns ColumnInfo,
	repo *git.Repository,
	prevCommit string,
) ([]result.Step, ColumnInfo, error) {
	subType, err := toSourceCodeSubType(r.Type)
	if err != nil {
		return nil, prevColumns, fmt.Errorf("toSourceSteps failed, %s", err)
	}

	switch subType {
	case SourceCommit:
		row, err := toSourceCommitRow(r)
		if err != nil {
			return nil, prevColumns, fmt.Errorf("toSourceSteps failed, %s", err)
		}

		steps, currentColumns, err := breakdownSourceCommitRow(row, finder, prevColumns, repo, prevCommit)
		if err != nil {
			return nil, prevColumns, fmt.Errorf("toSourceSteps failed, %s", err)
		}
		return steps, currentColumns, nil

	case SourceOpen:
		row, err := toSourceOpenRow(r)
		if err != nil {
			return nil, prevColumns, fmt.Errorf("toSourceSteps failed, %s", err)
		}

		steps, currentColumns, err := breakdownSourceOpenRow(row, finder, prevColumns, repo, prevCommit)
		if err != nil {
			return nil, prevColumns, fmt.Errorf("toSourceSteps failed, %s", err)
		}
		return steps, currentColumns, nil

	case SourceError:
		row, err := toSourceErrorRow(r)
		if err != nil {
			return nil, prevColumns, fmt.Errorf("toSourceSteps failed, %s", err)
		}

		steps, currentColumns, err := breakdownSourceErrorRow(row, finder, prevColumns, repo, prevCommit)
		if err != nil {
			return nil, prevColumns, fmt.Errorf("toSourceSteps failed, %s", err)
		}
		return steps, currentColumns, nil

	default:
		return nil, prevColumns, fmt.Errorf("toSourceSteps failed, type = '%s' not implemented", r.Type)
	}
}
