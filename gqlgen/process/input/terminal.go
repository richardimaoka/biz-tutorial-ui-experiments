package input

import (
	"fmt"
	"strings"
)

type TerminalTooltipRow struct {
	Contents string        `json:"contents"`
	Timing   TooltipTiming `json:"timing"`
}

type TerminalCommandRow struct {
	StepId  string              `json:"stepId"`
	Trivial bool                `json:"trivial"`
	Comment string              `json:"comment"`
	Command string              `json:"command"`
	Tooltip *TerminalTooltipRow `json:"tooltip"`
}

type TerminalOutputRow struct {
	StepId  string              `json:"stepId"`
	Trivial bool                `json:"trivial"`
	Comment string              `json:"comment"`
	Output  string              `json:"output"`
	Tooltip *TerminalTooltipRow `json:"tooltip"`
}

func toTerminalCommandRow(fromRow *Row) (*TerminalCommandRow, error) {
	errorPrefix := "failed to convert to TerminalCommand"

	//
	// Check column and type
	//
	if strings.ToLower(fromRow.Column) != Terminal {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.Column)
	}
	if strings.ToLower(fromRow.Type) != Command {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, fromRow.Type)
	}

	//
	// Check instruction
	//
	if fromRow.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}

	//
	// Check tooltip fields
	//
	terminalTooltip, err := toTerminalTooltip(fromRow)
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

	return &TerminalCommandRow{
		StepId:  fromRow.StepId,
		Trivial: trivial,
		Comment: fromRow.Comment,
		Command: fromRow.Instruction,
		Tooltip: terminalTooltip,
	}, nil
}

func toTerminalOutputRow(fromRow *Row) (*TerminalOutputRow, error) {
	errorPrefix := "failed to convert to TerminalOutput"

	//
	// Check column and type
	//
	if strings.ToLower(fromRow.Column) != Terminal {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.Column)
	}
	if strings.ToLower(fromRow.Type) != Output {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, fromRow.Type)
	}

	//
	// Check instruction
	//
	if fromRow.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}

	//
	// Check tooltip fields
	//
	terminalTooltip, err := toTerminalTooltip(fromRow)
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

	return &TerminalOutputRow{
		StepId:  fromRow.StepId,
		Trivial: trivial,
		Comment: fromRow.Comment,
		Output:  fromRow.Instruction,
		Tooltip: terminalTooltip,
	}, nil
}

func toTerminalRow(fromRow *Row) error {
	rowType := strings.ToLower(fromRow.Type)

	switch rowType {
	case Command:
		_, err := toTerminalCommandRow(fromRow)
		return err
	case Output:
		_, err := toTerminalOutputRow(fromRow)
		return err
	default:
		return fmt.Errorf("toTerminalRow failed, column = '%s' has wrong type = '%s'", fromRow.Column, fromRow.Type)
	}
}
