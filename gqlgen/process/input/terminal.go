package input

import (
	"fmt"
	"strings"
)

type TooltipTiming = string

const (
	BEGINNING TooltipTiming = "beginning"
	END       TooltipTiming = "end"
)

func toTooltipTiming(s string) (TooltipTiming, error) {
	switch strings.ToLower(s) {
	case BEGINNING:
		return BEGINNING, nil
	case END:
		return END, nil
	case "":
		return END, nil
	default:
		return "", fmt.Errorf("TooltipTiming value = '%s' is invalid", s)
	}
}

type TerminalTooltip struct {
	Contents string        `json:"contents"`
	Timing   TooltipTiming `json:"timing"`
}

type TerminalCommand struct {
	StepId  string           `json:"stepId"`
	Comment string           `json:"comment"`
	Command string           `json:"command"`
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
	var tooltip *TerminalTooltip
	if ab.Tooltip != "" {
		contents := ab.Tooltip

		tooltipTiming, err := toTooltipTiming(ab.TooltipTiming)
		if err != nil {
			return nil, fmt.Errorf("%s, 'tooltipTiming' field is wrong, %s", errorPrefix, err)
		}

		tooltip = &TerminalTooltip{
			Contents: contents,
			Timing:   tooltipTiming,
		}
	}

	return &TerminalCommand{
		StepId:  ab.StepId,
		Comment: ab.Comment,
		Command: ab.Instruction,
		Tooltip: tooltip,
	}, nil
}

type TerminalOutput struct {
	StepId  string           `json:"stepId"`
	Comment string           `json:"comment"`
	Output  string           `json:"output"`
	Tooltip *TerminalTooltip `json:"tooltip"`
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
	var tooltip *TerminalTooltip
	if ab.Tooltip != "" {
		contents := ab.Tooltip

		tooltipTiming, err := toTooltipTiming(ab.TooltipTiming)
		if err != nil {
			return nil, fmt.Errorf("%s, 'tooltipTiming' field is wrong, %s", errorPrefix, err)
		}

		tooltip = &TerminalTooltip{
			Contents: contents,
			Timing:   tooltipTiming,
		}
	}

	return &TerminalOutput{
		StepId:  ab.StepId,
		Comment: ab.Comment,
		Output:  ab.Instruction,
		Tooltip: tooltip,
	}, nil
}
