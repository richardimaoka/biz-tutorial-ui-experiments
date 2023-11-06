package input

import (
	"fmt"
	"strings"
)

type SourceTooltipRow struct {
	Contents   string        `json:"contents"`
	LineNumber int           `json:"lineNumber"`
	Timing     TooltipTiming `json:"timing"`
}

type SourceCommitRow struct {
	StepId          string            `json:"stepId"`
	Trivial         bool              `json:"trivial"`
	Comment         string            `json:"comment"`
	Commit          string            `json:"commit"`
	Tooltip         *SourceTooltipRow `json:"tooltip"`
	TypingAnimation bool              `json:"typingAnimation"`
	ShowDiff        bool              `json:"showDiff"`
}

type SourceOpenRow struct {
	StepId   string            `json:"stepId"`
	Trivial  bool              `json:"trivial"`
	Comment  string            `json:"comment"`
	FilePath string            `json:"filePath"`
	Tooltip  *SourceTooltipRow `json:"tooltip"`
}

type SourceErrorRow struct {
	StepId   string            `json:"stepId"`
	Trivial  bool              `json:"trivial"`
	Comment  string            `json:"comment"`
	FilePath string            `json:"filePath"`
	Tooltip  *SourceTooltipRow `json:"tooltip"`
}

func toSourceCommitRow(fromRow *Row) (*SourceCommitRow, error) {
	errorPrefix := "failed to convert to SourceCodeCommit"

	//
	// Check column and type
	//
	if strings.ToLower(fromRow.Column) != "source" {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.Column)
	}
	if fromRow.Type != "" && strings.ToLower(fromRow.Type) != "commit" {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, fromRow.Type)
	}

	//
	// Check instruction fields
	//
	if fromRow.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}

	//
	// Check tooltip fields
	//
	tooltip, err := toSourceTooltipRow(fromRow)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}

	typingAnimation, err := strToBool(fromRow.Instruction2)
	if err != nil {
		return nil, fmt.Errorf("%s, 'instruction2' is invalid, %s", errorPrefix, err)
	}

	showDiff, err := strToBool(fromRow.Instruction3)
	if err != nil {
		return nil, fmt.Errorf("%s, 'instruction3' is invalid, %s", errorPrefix, err)
	}

	//
	// Check trivial field
	//
	trivial, err := strToBool(fromRow.Trivial)
	if err != nil {
		return nil, fmt.Errorf("%s, 'trivial' is invalid, %s", errorPrefix, err)
	}

	return &SourceCommitRow{
		StepId:          fromRow.StepId,
		Trivial:         trivial,
		Comment:         fromRow.Comment,
		Commit:          fromRow.Instruction,
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
	if fromRow.Type != "" && strings.ToLower(fromRow.Type) != "commit" {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, fromRow.Type)
	}

	//
	// Check instruction fields
	//
	if fromRow.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}

	//
	// Check tooltip fields
	//
	tooltip, err := toSourceTooltipRow(fromRow)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}

	//
	// Check trivial field
	//
	trivial, err := strToBool(fromRow.Trivial)
	if err != nil {
		return nil, fmt.Errorf("%s, 'trivial' is invalid, %s", errorPrefix, err)
	}

	return &SourceOpenRow{
		StepId:   fromRow.StepId,
		Trivial:  trivial,
		Comment:  fromRow.Comment,
		FilePath: fromRow.Instruction,
		Tooltip:  tooltip,
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
	if fromRow.Type != "" && strings.ToLower(fromRow.Type) != "commit" {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, fromRow.Type)
	}

	//
	// Check instruction fields
	//
	if fromRow.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}

	//
	// Check tooltip fields
	//
	tooltip, err := toSourceTooltipRow(fromRow)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}
	if tooltip == nil {
		return nil, fmt.Errorf("%s, source code error needs the error detail in 'tooltip'", errorPrefix)
	}

	//
	// Check trivial field
	//
	trivial, err := strToBool(fromRow.Trivial)
	if err != nil {
		return nil, fmt.Errorf("%s, 'trivial' is invalid, %s", errorPrefix, err)
	}

	return &SourceErrorRow{
		StepId:   fromRow.StepId,
		Trivial:  trivial,
		Comment:  fromRow.Comment,
		FilePath: fromRow.Instruction,
		Tooltip:  tooltip,
	}, nil
}

func toSourceTooltipRow(fromRow *Row) (*SourceTooltipRow, error) {
	if fromRow.Tooltip == "" {
		return nil, nil
	}

	contents := fromRow.Tooltip

	tooltipTiming, err := toTooltipTiming(fromRow.TooltipTiming)
	if err != nil {
		return nil, fmt.Errorf("'tooltipTiming' field is wrong, %s", err)
	}

	if fromRow.TooltipLine == 0 {
		return nil, fmt.Errorf("'tooltipLine' cannot be 0")
	}

	return &SourceTooltipRow{
		Contents:   contents,
		Timing:     tooltipTiming,
		LineNumber: fromRow.TooltipLine,
	}, nil
}
