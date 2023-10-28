package input

import (
	"fmt"
	"strings"
)

type TerminalTooltip struct {
	Contents string        `json:"contents"`
	Timing   TooltipTiming `json:"timing"`
}

type TerminalCommand struct {
	StepId  string           `json:"stepId"`
	Trivial bool             `json:"trivial"`
	Comment string           `json:"comment"`
	Command string           `json:"command"`
	Tooltip *TerminalTooltip `json:"tooltip"`
}

type TerminalOutput struct {
	StepId  string           `json:"stepId"`
	Trivial bool             `json:"trivial"`
	Comment string           `json:"comment"`
	Output  string           `json:"output"`
	Tooltip *TerminalTooltip `json:"tooltip"`
}

func toTerminalCommand(ab *Abstract) (*TerminalCommand, error) {
	errorPrefix := "failed to convert to TerminalCommand"

	//
	// Check column and type
	//
	if strings.ToLower(ab.Column) != "terminal" {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, ab.Column)
	}
	if ab.Type != "" && strings.ToLower(ab.Type) != "command" {
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
	terminalTooltip, err := toTerminalTooltip(ab)
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

	return &TerminalCommand{
		StepId:  ab.StepId,
		Trivial: trivial,
		Comment: ab.Comment,
		Command: ab.Instruction,
		Tooltip: terminalTooltip,
	}, nil
}

func toTerminalOutput(ab *Abstract) (*TerminalOutput, error) {
	errorPrefix := "failed to convert to TerminalOutput"

	//
	// Check column and type
	//
	if strings.ToLower(ab.Column) != "terminal" {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, ab.Column)
	}
	if ab.Type != "" && strings.ToLower(ab.Type) != "output" {
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
	terminalTooltip, err := toTerminalTooltip(ab)
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

	return &TerminalOutput{
		StepId:  ab.StepId,
		Trivial: trivial,
		Comment: ab.Comment,
		Output:  ab.Instruction,
		Tooltip: terminalTooltip,
	}, nil
}

func toTerminalTooltip(ab *Abstract) (*TerminalTooltip, error) {
	if ab.Tooltip == "" {
		return nil, nil
	}
	contents := ab.Tooltip

	tooltipTiming, err := toTooltipTiming(ab.TooltipTiming)
	if err != nil {
		return nil, fmt.Errorf("'tooltipTiming' field is wrong, %s", err)
	}

	return &TerminalTooltip{
		Contents: contents,
		Timing:   tooltipTiming,
	}, nil
}
