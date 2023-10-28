package input

import (
	"fmt"
	"strings"
)

type SourceCodeTooltip struct {
	Contents   string        `json:"contents"`
	LineNumber int           `json:"lineNumber"`
	Timing     TooltipTiming `json:"timing"`
}

type SourceCodeCommit struct {
	StepId          string             `json:"stepId"`
	Trivial         bool               `json:"trivial"`
	Comment         string             `json:"comment"`
	Commit          string             `json:"commit"`
	Tooltip         *SourceCodeTooltip `json:"tooltip"`
	TypingAnimation bool               `json:"typingAnimation"`
	ShowDiff        bool               `json:"showDiff"`
}

type SourceCodeOpen struct {
	StepId   string             `json:"stepId"`
	Trivial  bool               `json:"trivial"`
	Comment  string             `json:"comment"`
	FilePath string             `json:"filePath"`
	Tooltip  *SourceCodeTooltip `json:"tooltip"`
}

type SourceCodeError struct {
	StepId   string             `json:"stepId"`
	Trivial  bool               `json:"trivial"`
	Comment  string             `json:"comment"`
	FilePath string             `json:"filePath"`
	Tooltip  *SourceCodeTooltip `json:"tooltip"`
}

func toSourceCodeCommit(ab *Abstract) (*SourceCodeCommit, error) {
	errorPrefix := "failed to convert to SourceCodeCommit"

	//
	// Check column and type
	//
	if strings.ToLower(ab.Column) != "source" {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, ab.Column)
	}
	if ab.Type != "" && strings.ToLower(ab.Type) != "commit" {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, ab.Type)
	}

	//
	// Check instruction fields
	//
	if ab.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}

	//
	// Check tooltip fields
	//
	tooltip, err := toSourceCodeTooltip(ab)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}

	typingAnimation, err := strToBool(ab.Instruction2)
	if err != nil {
		return nil, fmt.Errorf("%s, 'instruction2' is invalid, %s", errorPrefix, err)
	}

	showDiff, err := strToBool(ab.Instruction3)
	if err != nil {
		return nil, fmt.Errorf("%s, 'instruction3' is invalid, %s", errorPrefix, err)
	}

	//
	// Check trivial field
	//
	trivial, err := strToBool(ab.Trivial)
	if err != nil {
		return nil, fmt.Errorf("%s, 'trivial' is invalid, %s", errorPrefix, err)
	}

	return &SourceCodeCommit{
		StepId:          ab.StepId,
		Trivial:         trivial,
		Comment:         ab.Comment,
		Commit:          ab.Instruction,
		Tooltip:         tooltip,
		TypingAnimation: typingAnimation,
		ShowDiff:        showDiff,
	}, nil
}

func toSourceCodeOpen(ab *Abstract) (*SourceCodeOpen, error) {
	errorPrefix := "failed to convert to SourceCodeOpen"

	//
	// Check column and type
	//
	if strings.ToLower(ab.Column) != "source" {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, ab.Column)
	}
	if ab.Type != "" && strings.ToLower(ab.Type) != "commit" {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, ab.Type)
	}

	//
	// Check instruction fields
	//
	if ab.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}

	//
	// Check tooltip fields
	//
	tooltip, err := toSourceCodeTooltip(ab)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}

	//
	// Check trivial field
	//
	trivial, err := strToBool(ab.Trivial)
	if err != nil {
		return nil, fmt.Errorf("%s, 'trivial' is invalid, %s", errorPrefix, err)
	}

	return &SourceCodeOpen{
		StepId:   ab.StepId,
		Trivial:  trivial,
		Comment:  ab.Comment,
		FilePath: ab.Instruction,
		Tooltip:  tooltip,
	}, nil
}

func toSourceCodeError(ab *Abstract) (*SourceCodeError, error) {
	errorPrefix := "failed to convert to SourceCodeError"

	//
	// Check column and type
	//
	if strings.ToLower(ab.Column) != "source" {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, ab.Column)
	}
	if ab.Type != "" && strings.ToLower(ab.Type) != "commit" {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, ab.Type)
	}

	//
	// Check instruction fields
	//
	if ab.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}

	//
	// Check tooltip fields
	//
	tooltip, err := toSourceCodeTooltip(ab)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}
	if tooltip == nil {
		return nil, fmt.Errorf("%s, source code error needs the error detail in 'tooltip'", errorPrefix)
	}

	//
	// Check trivial field
	//
	trivial, err := strToBool(ab.Trivial)
	if err != nil {
		return nil, fmt.Errorf("%s, 'trivial' is invalid, %s", errorPrefix, err)
	}

	return &SourceCodeError{
		StepId:   ab.StepId,
		Trivial:  trivial,
		Comment:  ab.Comment,
		FilePath: ab.Instruction,
		Tooltip:  tooltip,
	}, nil
}

func toSourceCodeTooltip(ab *Abstract) (*SourceCodeTooltip, error) {
	if ab.Tooltip == "" {
		return nil, nil
	}

	contents := ab.Tooltip

	tooltipTiming, err := toTooltipTiming(ab.TooltipTiming)
	if err != nil {
		return nil, fmt.Errorf("'tooltipTiming' field is wrong, %s", err)
	}

	if ab.TooltipLine == 0 {
		return nil, fmt.Errorf("'tooltipLine' cannot be 0")
	}

	return &SourceCodeTooltip{
		Contents:   contents,
		Timing:     tooltipTiming,
		LineNumber: ab.TooltipLine,
	}, nil
}
