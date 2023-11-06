package input

import (
	"fmt"
	"strings"
)

type TerminalTooltip struct {
	Contents string        `json:"contents"`
	Timing   TooltipTiming `json:"timing"`
}

type TerminalCommandRow struct {
	StepId        string           `json:"stepId"`
	IsTrivial     bool             `json:"isTrivial"`
	Comment       string           `json:"comment"`
	Command       string           `json:"command"`
	Tooltip       *TerminalTooltip `json:"tooltip"`
	ModalContents string           `json:"modalContents"`
	TerminalName  string           `json:"terminalName"`
}

type TerminalOutputRow struct {
	StepId        string           `json:"stepId"`
	IsTrivial     bool             `json:"isTrivial"`
	Comment       string           `json:"comment"`
	Output        string           `json:"output"`
	Tooltip       *TerminalTooltip `json:"tooltip"`
	ModalContents string           `json:"modalContents"`
	TerminalName  string           `json:"terminalName"`
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
	// Check instruction fields
	//
	if fromRow.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}
	var terminalName string
	if fromRow.Instruction2 == "" {
		terminalName = "default"
	} else {
		terminalName = fromRow.Instruction2
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
		StepId:        fromRow.StepId,
		IsTrivial:     trivial,
		Comment:       fromRow.Comment,
		Command:       fromRow.Instruction,
		Tooltip:       terminalTooltip,
		ModalContents: fromRow.ModalContents,
		TerminalName:  terminalName,
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
	// Check instruction fields
	//
	if fromRow.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}
	var terminalName string
	if fromRow.Instruction2 == "" {
		terminalName = "default"
	} else {
		terminalName = fromRow.Instruction2
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
		StepId:        fromRow.StepId,
		IsTrivial:     trivial,
		Comment:       fromRow.Comment,
		Output:        fromRow.Instruction,
		Tooltip:       terminalTooltip,
		ModalContents: fromRow.ModalContents,
		TerminalName:  terminalName,
	}, nil
}

/**
 * Functions to generate step from row
 */

func terminalCommandStep(r *TerminalCommandRow, StepIdFinder *StepIdFinder, usedColumns UsedColumns) ResultStep {
	subId := "terminalCommandStep"
	stepId := StepIdFinder.StepIdFor(r.StepId, subId)

	step := ResultStep{
		// fields to make the step searchable for re-generation
		IsFromRow:  true,
		ParentStep: r.StepId,
		SubID:      subId,
		// other fields
		Step:          stepId,
		FocusColumn:   "Terminal",
		Comment:       r.Comment,
		ModalContents: r.ModalContents,
		// Terminal fields
		TerminalType: "command",
		TerminalText: r.Command,
		TerminalName: r.TerminalName,
	}
	if r.Tooltip != nil {
		step.TerminalTooltipContents = r.Tooltip.Contents
		step.TerminalTooltipTiming = r.Tooltip.Timing
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
		Step:          stepId,
		FocusColumn:   "Terminal",
		Comment:       r.Comment,
		ModalContents: r.ModalContents,
		// Terminal fields
		TerminalType: "output",
		TerminalText: r.Output,
		TerminalName: r.TerminalName,
	}
	if r.Tooltip != nil {
		step.TerminalTooltipContents = r.Tooltip.Contents
		step.TerminalTooltipTiming = r.Tooltip.Timing
	}

	step.setColumns(usedColumns)

	return step
}

func moveToTerminalCommandStep(r *TerminalCommandRow, finder *StepIdFinder, usedColumns UsedColumns) ResultStep {
	subId := "moveToTerminalStep"
	stepId := finder.StepIdFor(r.StepId, subId)
	step := ResultStep{
		// fields to make the step searchable for re-generation
		IsFromRow:  true,
		ParentStep: r.StepId,
		SubID:      subId,
		// other fields
		Step:        stepId,
		FocusColumn: "Terminal",
		Comment:     "(move to Terminal)",
		IsTrivial:   true,
		// Terminal fields
		TerminalType: "cd",
		TerminalName: r.TerminalName,
	}
	step.setColumns(usedColumns)

	return step
}

func moveToTerminalOutputStep(r *TerminalOutputRow, finder *StepIdFinder, usedColumns UsedColumns) ResultStep {
	subId := "moveToTerminalStep"
	stepId := finder.StepIdFor(r.StepId, subId)
	step := ResultStep{
		// fields to make the step searchable for re-generation
		IsFromRow:  true,
		ParentStep: r.StepId,
		SubID:      subId,
		// other fields
		Step:        stepId,
		FocusColumn: "Terminal",
		Comment:     "(move to Terminal)",
		IsTrivial:   true,
		// Terminal fields
		TerminalType: "cd",
		TerminalName: r.TerminalName,
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
		Step:          stepId,
		FocusColumn:   "Terminal",
		ModalContents: r.ModalContents,
		// Terminal fields
		TerminalType: "cd",
		TerminalName: r.TerminalName,
		CurrentDir:   currentDir,
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
