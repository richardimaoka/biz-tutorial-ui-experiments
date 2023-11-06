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

/**
 * Functions to convert to specific row
 */
func toTerminalCommandRow(fromRow *Row) (*TerminalCommandRow, error) {
	errorPrefix := "failed to convert to TerminalCommand"

	//
	// Check column and type
	//
	if strings.ToLower(fromRow.Column) != TerminalType {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.Column)
	}
	if strings.ToLower(fromRow.Type) != CommandSubType {
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
	if strings.ToLower(fromRow.Column) != TerminalType {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.Column)
	}
	if strings.ToLower(fromRow.Type) != OutputSubType {
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

/**
 * Functions to generate step from row
 */
func moveToTerminalStep(parentStepId string, finder *StepIdFinder, usedColumns UsedColumns) ResultStep {
	subId := "moveToTerminalStep"
	stepId := finder.StepIdFor(parentStepId, subId)
	step := ResultStep{
		// fields to make the step searchable for re-generation
		IsFromRow:  true,
		ParentStep: parentStepId,
		SubID:      subId,
		// other fields
		Step:        stepId,
		FocusColumn: "Terminal",
		Comment:     "(move to Terminal)",
	}
	step.setColumns(usedColumns)

	return step
}

func terminalOutputStep(r *TerminalOutputRow, finder *StepIdFinder, usedColumns UsedColumns) ResultStep {
	subId := "terminalOutputStep"
	stepId := finder.StepIdFor(r.StepId, subId)
	step := ResultStep{
		// fields to make the step searchable for re-generation
		IsFromRow:  true,
		ParentStep: r.StepId,
		SubID:      subId,
		// other fields
		Step:         stepId,
		FocusColumn:  "Terminal",
		TerminalType: "output",
		TerminalText: r.Output,
		// ModalText:    r.ModalText,
	}

	if r.Tooltip != nil {
		step.TerminalTooltip = r.Tooltip.Contents
		step.TerminalTooltipTiming = r.Tooltip.Timing
	}

	step.setColumns(usedColumns)

	return step
}

func terminalCommandStep(r *TerminalCommandRow, StepIdFinder *StepIdFinder, usedColumns UsedColumns) ResultStep {
	subId := "terminalCommandStep"
	stepId := StepIdFinder.StepIdFor(r.StepId, subId)

	step := ResultStep{
		// fields to make the step searchable for re-generation
		IsFromRow:  true,
		ParentStep: r.StepId,
		SubID:      subId,
		// other fields
		Step:         stepId,
		FocusColumn:  "Terminal",
		TerminalType: "command",
		TerminalText: r.Command,
		// TerminalName: , // Go zero value is ""
		// ModalText: r.ModalText,
	}

	if r.Tooltip != nil {
		step.TerminalTooltip = r.Tooltip.Contents
		step.TerminalTooltipTiming = r.Tooltip.Timing
	}

	step.setColumns(usedColumns)

	return step
}

func terminalCdStep(r *TerminalCommandRow, StepIdFinder *StepIdFinder, usedColumns UsedColumns) ResultStep {
	currentDir := strings.TrimPrefix(r.Command, "cd ")

	subId := "terminalCdStep"
	stepId := StepIdFinder.StepIdFor(r.StepId, subId)

	step := ResultStep{
		// fields to make the step searchable for re-generation
		IsFromRow:  true,
		ParentStep: r.StepId,
		SubID:      subId,
		// other fields
		Step:         stepId,
		FocusColumn:  "Terminal",
		TerminalType: "cd",
		// TerminalName: r.Instruction3, // Go zero value is ""
		CurrentDir: currentDir, // Go zero value is ""
		// ModalText:    r.ModalText,
	}

	// cd command should be trivial and no tooltip to show

	step.setColumns(usedColumns)

	return step
}

func toTerminalRow(fromRow *Row) error {
	rowType := strings.ToLower(fromRow.Type)

	switch rowType {
	case CommandSubType:
		_, err := toTerminalCommandRow(fromRow)
		return err
	case OutputSubType:
		_, err := toTerminalOutputRow(fromRow)
		return err
	default:
		return fmt.Errorf("toTerminalRow failed, column = '%s' has wrong type = '%s'", fromRow.Column, fromRow.Type)
	}
}
