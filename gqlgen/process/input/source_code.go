package input

import (
	"fmt"
	"strings"
)

type SourceCodeTooltip struct {
	Contents           string             `json:"contents"`
	LineNumber         int                `json:"lineNumber"`
	PositionPreference PositionPreference `json:"positionPreference"`
	Timing             TooltipTiming      `json:"timing"`
}

type SourceCodeCommit struct {
	StepId          string             `json:"stepId"`
	Comment         string             `json:"comment"`
	Commit          string             `json:"commit"`
	Tooltip         *SourceCodeTooltip `json:"tooltip"`
	TypingAnimation bool               `json:"typingAnimation"`
	ShowDiff        bool               `json:"showDiff"`
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

	typingAnimation, err := strToBool(ab.Instruction2)
	if err != nil {
		return nil, fmt.Errorf("%s, 'instruction2' is invalid, %s", errorPrefix, err)
	}

	showDiff, err := strToBool(ab.Instruction3)
	if err != nil {
		return nil, fmt.Errorf("%s, 'instruction3' is invalid, %s", errorPrefix, err)
	}

	//
	// Check tooltip fields
	//
	tooltip, err := toSourceCodeTooltip(ab)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}

	return &SourceCodeCommit{
		StepId:          ab.StepId,
		Comment:         ab.Comment,
		Commit:          ab.Instruction,
		Tooltip:         tooltip,
		TypingAnimation: typingAnimation,
		ShowDiff:        showDiff,
	}, nil
}

type SourceCodeOpen struct {
	Comment  string
	FilePath string
	Tooltip  *SourceCodeTooltip
}

type SourceCodeError struct {
	Comment  string
	FilePath string
	Tooltip  *SourceCodeTooltip
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

	tooltipPosition, err := toPositionPreference(ab.TooltipTiming)
	if err != nil {
		return nil, fmt.Errorf("'tooltipPosition' field is wrong, %s", err)
	}

	return &SourceCodeTooltip{
		Contents:           contents,
		Timing:             tooltipTiming,
		PositionPreference: tooltipPosition,
	}, nil
}
