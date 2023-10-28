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

type TerminalOutput struct {
	StepId  string           `json:"stepId"`
	Comment string           `json:"comment"`
	Output  string           `json:"output"`
	Tooltip *TerminalTooltip `json:"tooltip"`
}

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

func toTerminalCommand(ab *Abstract) (*TerminalCommand, error) {
	if ab.Instruction == "" {
		return nil, fmt.Errorf("failed to convert to TerminalCommand, 'instruction' was empty")
	}

	var tooltip *TerminalTooltip
	if ab.Tooltip != "" {
		contents := ab.Tooltip

		tooltipTiming, err := toTooltipTiming(ab.TooltipTiming)
		if err != nil {
			return nil, fmt.Errorf("failed to convert to TerminalCommand, 'tooltipTiming' field is wrong, %s", err)
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
