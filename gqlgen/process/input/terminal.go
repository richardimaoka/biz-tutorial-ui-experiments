package input

import (
	"fmt"
	"strings"
)

type TerminalTooltip struct {
	Contents string        `json:"contents"`
	Timing   TooltipTiming `json:"timing"`
}

type TerminalRow struct {
	StepId        string           `json:"stepId"`
	IsTrivial     bool             `json:"isTrivial"`
	Comment       string           `json:"comment"`
	Type          string           `json:"type"`
	Text          string           `json:"text"`
	Tooltip       *TerminalTooltip `json:"tooltip"`
	ModalContents string           `json:"modalContents"`
	TerminalName  string           `json:"terminalName"`
}

/**
 * Function(s) to convert a row to a more specific row
 */

func toTerminalRow(fromRow *Row) (*TerminalRow, error) {
	errorPrefix := "failed to convert to TerminalOutput"

	//
	// Check column and type
	//
	if strings.ToLower(fromRow.Column) != TerminalType {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.Column)
	}
	if strings.ToLower(fromRow.Type) != CommandSubType && strings.ToLower(fromRow.Type) != OutputSubType {
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

	return &TerminalRow{
		StepId:        fromRow.StepId,
		IsTrivial:     trivial,
		Comment:       fromRow.Comment,
		ModalContents: fromRow.ModalContents,
		Type:          fromRow.Type,
		Text:          fromRow.Instruction,
		Tooltip:       terminalTooltip,
		TerminalName:  terminalName,
	}, nil
}

/**
 * Function(s) to convert a row to a step
 */
func terminalCommandStep(r *TerminalRow, StepIdFinder *StepIdFinder, usedColumns UsedColumns) ResultStep {
	subId := "terminalCommandStep"
	stepId := StepIdFinder.StepIdFor(r.StepId, subId)

	step := ResultStep{
		// fields to make the step searchable for re-generation
		IsFromRow:  true,
		ParentStep: r.StepId,
		SubID:      subId,
		// Other fields
		Step:          stepId,
		Comment:       r.Comment,
		FocusColumn:   "Terminal",
		ModalContents: r.ModalContents,
		// Terminal fields
		TerminalType: "command",
		TerminalText: r.Text,
		TerminalName: r.TerminalName,
	}
	if r.Tooltip != nil {
		step.TerminalTooltipContents = r.Tooltip.Contents
		step.TerminalTooltipTiming = r.Tooltip.Timing
	}

	step.setColumns(usedColumns)

	return step
}

func terminalOutputStep(r *TerminalRow, finder *StepIdFinder, usedColumns UsedColumns) ResultStep {
	subId := "terminalOutputStep"
	stepId := finder.StepIdFor(r.StepId, subId)
	step := ResultStep{
		// Fields to make the step searchable for re-generation
		IsFromRow:  true,
		ParentStep: r.StepId,
		SubID:      subId,
		// Other fields
		Step:          stepId,
		IsTrivial:     r.IsTrivial,
		Comment:       r.Comment,
		FocusColumn:   "Terminal",
		ModalContents: r.ModalContents,
		// Terminal fields
		TerminalType: "output",
		TerminalText: r.Text,
		TerminalName: r.TerminalName,
	}
	if r.Tooltip != nil {
		step.TerminalTooltipContents = r.Tooltip.Contents
		step.TerminalTooltipTiming = r.Tooltip.Timing
	}

	step.setColumns(usedColumns)

	return step
}

func moveToTerminalStep(r *TerminalRow, finder *StepIdFinder, usedColumns UsedColumns) ResultStep {
	subId := "moveToTerminalStep"
	stepId := finder.StepIdFor(r.StepId, subId)
	step := ResultStep{
		// Fields to make the step searchable for re-generation
		IsFromRow:  true,
		ParentStep: r.StepId,
		SubID:      subId,
		// Other fields
		Step:          stepId,
		IsTrivial:     true, // always trivial
		Comment:       "(move to Terminal)",
		FocusColumn:   "Terminal",
		ModalContents: r.ModalContents,
		// Terminal fields
		TerminalType: "cd",
		TerminalName: r.TerminalName,
	}
	// No tooltip - move step should be trivial and no tooltip to show

	step.setColumns(usedColumns)

	return step
}

func terminalCdStep(r *TerminalRow, StepIdFinder *StepIdFinder, usedColumns UsedColumns) ResultStep {
	currentDir := strings.TrimPrefix(r.Text, "cd ")

	subId := "terminalCdStep"
	stepId := StepIdFinder.StepIdFor(r.StepId, subId)

	step := ResultStep{
		// Fields to make the step searchable for re-generation
		IsFromRow:  true,
		ParentStep: r.StepId,
		SubID:      subId,
		// other fields
		Step:          stepId,
		IsTrivial:     true, // always trivial
		Comment:       "",
		FocusColumn:   "Terminal",
		ModalContents: r.ModalContents,
		// Terminal fields
		TerminalType: "cd",
		TerminalName: r.TerminalName,
		CurrentDir:   currentDir,
	}
	// No tooltip - cd command should be trivial and no tooltip to show

	step.setColumns(usedColumns)

	return step
}

func toTerminalResultSteps(fromRow *Row) error {
	_, err := toTerminalRow(fromRow)
	return err
	// rowType := strings.ToLower(fromRow.Type)
	// switch rowType {
	// case CommandSubType:
	// case OutputSubType:
	// 	_, err := toTerminalRow(fromRow)
	// 	return err
	// default:
	// 	return fmt.Errorf("toTerminalRow failed, column = '%s' has wrong type = '%s'", fromRow.Column, fromRow.Type)
	// }
}

/**
 * Function(s) to break down a row to steps
 */
// func toTerminalResultSteps2(
// 	r *TerminalRow,
// 	finder *StepIdFinder,
// 	prevColumns ColumnInfo,
// ) ([]ResultStep, ColumnInfo, error) {
// 	usedColumns := appendIfNotExists(existingColumns, "Terminal")

// 	// - precondition for RoughStep

// 	// check if it's a valid terminal step
// 	if r.Instruction == "" {
// 		return nil, NoColumn, EmptyColumns, fmt.Errorf("step is missing 'instruction', step = '%s', type = '%s'", r.Step, r.Type)
// 	}

// 	// - step creation
// 	var steps []DetailedStep

// 	// insert move-to-terminal step if current column != "Terminal"
// 	if prevColumn != "Terminal" && prevColumn != "" {
// 		moveToTerminalStep := moveToTerminalStep(r, uuidFinder, usedColumns)
// 		steps = append(steps, moveToTerminalStep)
// 	}

// 	// command step
// 	cmdStep := terminalCommandStep(r, uuidFinder, usedColumns)
// 	steps = append(steps, cmdStep)

// 	// cd step
// 	if strings.HasPrefix(r.Instruction, "cd ") {
// 		cmdStep := terminalCdStep(r, uuidFinder, usedColumns)
// 		steps = append(steps, cmdStep)
// 	}

// 	return steps, "Terminal", usedColumns, nil
// }

// func terminalOutputConvert(
// 	s *RoughStep,
// 	uuidFinder *UUIDFinder,
// 	existingColumns UsedColumns,
// ) ([]DetailedStep, CurrentColumn, UsedColumns, error) {
// 	usedColumns := appendIfNotExists(existingColumns, "Terminal")

// 	// - precondition for RoughStep

// 	// - step creation
// 	var steps []DetailedStep

// 	// output step
// 	s.Instruction2 = s.Instruction //TODO: workaround for now
// 	outputStep := terminalOutputStep(s, uuidFinder, usedColumns)
// 	steps = append(steps, outputStep)

// 	return steps, "Terminal", usedColumns, nil
// }
