package input

import (
	"fmt"
	"strings"
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
