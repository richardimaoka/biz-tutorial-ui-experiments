package input

import (
	"fmt"
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/result"
)

/**
 * TerminalSubType type(s) and functions
 */
type TerminalSubType string

const (
	// Lower cases since they are from manual entries
	TerminalCommand TerminalSubType = "command"
	TerminalOutput  TerminalSubType = "output"
	TerminalOpen    TerminalSubType = "open"
)

func toTerminalSubType(s string) (TerminalSubType, error) {
	lower := strings.ToLower(s)

	switch lower {
	case string(TerminalCommand):
		return TerminalCommand, nil
	case string(TerminalOutput):
		return TerminalOutput, nil
	case string(TerminalOpen):
		return TerminalOpen, nil
	default:
		return "", fmt.Errorf("'%s' is an invalid terminal sub type", s)
	}
}

/**
 * TerminalTooltip type(s) and functions
 */
type TerminalTooltip struct {
	Contents string        `json:"contents"`
	Timing   TooltipTiming `json:"timing"`
}

func toTerminalTooltipTiming(s string) (TooltipTiming, error) {
	switch strings.ToUpper(s) {
	case START:
		return START, nil
	case END:
		return END, nil
	case "": // default value is different from source tooltip
		return START, nil
	default:
		return "", fmt.Errorf("TooltipTiming value = '%s' is invalid", s)
	}
}

func toTerminalTooltip(fromRow *Row) (*TerminalTooltip, error) {
	// if tooltip is empty, then return no tooltip
	if fromRow.Tooltip == "" {
		return nil, nil
	}

	contents := fromRow.Tooltip

	tooltipTiming, err := toTerminalTooltipTiming(fromRow.TooltipTiming)
	if err != nil {
		return nil, fmt.Errorf("'tooltipTiming' field is wrong, %s", err)
	}

	return &TerminalTooltip{
		Contents: contents,
		Timing:   tooltipTiming,
	}, nil
}

/**
 * Terminaljrow type(s) and functions
 */
type TerminalRow struct {
	StepId        string           `json:"stepId"`
	IsTrivial     bool             `json:"isTrivial"`
	Comment       string           `json:"comment"`
	Type          TerminalSubType  `json:"type"`
	Text          string           `json:"text"`
	Tooltip       *TerminalTooltip `json:"tooltip"`
	ModalContents string           `json:"modalContents"`
	TerminalName  string           `json:"terminalName"`
}

/**
 * Function(s) to convert a row to a more specific row
 */

func toTerminalCommandRow(fromRow *Row) (*TerminalRow, error) {
	errorPrefix := "failed in toTerminalCommandRow"

	//
	// Check column and type
	//
	column, err := toColumnType(fromRow.Column)
	if err != nil || column != TerminalColumn {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.Column)
	}
	subType, err := toTerminalSubType(fromRow.Type)
	if err != nil || subType != TerminalCommand {
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
		Type:          subType,
		Text:          fromRow.Instruction,
		Tooltip:       terminalTooltip,
		TerminalName:  terminalName,
	}, nil
}

func toTerminalOutputRow(fromRow *Row) (*TerminalRow, error) {
	errorPrefix := "failed in toTerminalOutputRow"

	//
	// Check column and type
	//
	column, err := toColumnType(fromRow.Column)
	if err != nil || column != TerminalColumn {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.Column)
	}
	subType, err := toTerminalSubType(fromRow.Type)
	if err != nil || subType != TerminalOutput {
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
		Type:          subType,
		Text:          fromRow.Instruction,
		Tooltip:       terminalTooltip,
		TerminalName:  terminalName,
	}, nil
}

func toTerminalOpenRow(fromRow *Row) (*TerminalRow, error) {
	errorPrefix := "failed in toTerminalOpenRow"

	//
	// Check column and type
	//
	column, err := toColumnType(fromRow.Column)
	if err != nil || column != TerminalColumn {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.Column)
	}
	subType, err := toTerminalSubType(fromRow.Type)
	if err != nil {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, fromRow.Type)
	}

	//
	// Check instruction fields
	//
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
		Type:          subType,
		Tooltip:       terminalTooltip,
		TerminalName:  terminalName,
	}, nil
}

/**
 * Function(s) to convert a row to a step
 */
func terminalCommandStep(r *TerminalRow, StepIdFinder *StepIdFinder, currentColumns result.ColumnFields) result.Step {
	subId := "terminalCommandStep"
	stepId := StepIdFinder.StepIdFor(r.StepId, subId)

	step := result.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: result.FromRowFields{
			IsFromRow:  true,
			ParentStep: r.StepId,
			SubID:      subId,
		},
		IntrinsicFields: result.IntrinsicFields{
			StepId:  stepId,
			Comment: r.Comment,
		},
		AnimationFields: result.AnimationFields{
			IsTrivial: r.IsTrivial,
		},
		ModalFields: result.ModalFields{
			ModalContents: r.ModalContents,
		},
		ColumnFields: currentColumns,
		TerminalFields: result.TerminalFields{
			TerminalStepType: result.TerminalCommand,
			TerminalText:     r.Text,
			TerminalName:     r.TerminalName,
		},
	}

	if r.Tooltip != nil {
		step.TerminalTooltipContents = r.Tooltip.Contents
		step.TerminalTooltipTiming = r.Tooltip.Timing
	}

	return step
}

func terminalOutputStep(r *TerminalRow, finder *StepIdFinder, currentColumns result.ColumnFields) result.Step {
	subId := "terminalOutputStep"
	stepId := finder.StepIdFor(r.StepId, subId)

	step := result.Step{
		// Fields to make the step searchable for re-generation
		FromRowFields: result.FromRowFields{
			IsFromRow:  true,
			ParentStep: r.StepId,
			SubID:      subId,
		},
		IntrinsicFields: result.IntrinsicFields{
			StepId:  stepId,
			Comment: r.Comment,
		},
		AnimationFields: result.AnimationFields{
			IsTrivial: r.IsTrivial,
		},
		ModalFields: result.ModalFields{
			ModalContents: r.ModalContents,
		},
		ColumnFields: currentColumns,
		TerminalFields: result.TerminalFields{
			TerminalStepType: result.TerminalOutput,
			TerminalText:     r.Text,
			TerminalName:     r.TerminalName,
		},
	}
	if r.Tooltip != nil {
		step.TerminalTooltipContents = r.Tooltip.Contents
		step.TerminalTooltipTiming = r.Tooltip.Timing
	}

	return step
}

func moveToTerminalStep(r *TerminalRow, finder *StepIdFinder, currentColumns result.ColumnFields) result.Step {
	subId := "moveToTerminalStep"
	stepId := finder.StepIdFor(r.StepId, subId)

	step := result.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: result.FromRowFields{
			IsFromRow:  true,
			ParentStep: r.StepId,
			SubID:      subId,
		},
		IntrinsicFields: result.IntrinsicFields{
			StepId:  stepId,
			Comment: "(move to Terminal)",
		},
		AnimationFields: result.AnimationFields{
			IsTrivial: true, //always true
		},
		// No ModalFields, as it is a trivial step
		ColumnFields: currentColumns,
		TerminalFields: result.TerminalFields{
			TerminalStepType: result.TerminalMove,
			TerminalName:     r.TerminalName,
		},
	}
	// No tooltip - move step should be trivial and no tooltip to show

	return step
}

func terminalCdStep(r *TerminalRow, StepIdFinder *StepIdFinder, currentColumns result.ColumnFields) result.Step {
	currentDir := strings.TrimPrefix(r.Text, "cd ")

	subId := "terminalCdStep"
	stepId := StepIdFinder.StepIdFor(r.StepId, subId)

	step := result.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: result.FromRowFields{
			IsFromRow:  true,
			ParentStep: r.StepId,
			SubID:      subId,
		},
		IntrinsicFields: result.IntrinsicFields{
			StepId:  stepId,
			Comment: "",
		},
		AnimationFields: result.AnimationFields{
			IsTrivial: true, //always true
		},
		// No ModalFields, as it is a trivial step
		ColumnFields: currentColumns,
		TerminalFields: result.TerminalFields{
			CurrentDir:       currentDir,
			TerminalStepType: result.TerminalCd,
			TerminalName:     r.TerminalName,
		},
	}

	// No tooltip - trivial step and no tooltip to show

	return step
}

/**
 * Function(s) to break down a row to steps
 */
func breakdownTerminalRow(r *TerminalRow, finder *StepIdFinder, prevColumns *ColumnInfo) []result.Step {
	// - step creation
	var steps []result.Step
	currentColumns := resultColumns(result.TerminalColumnType, prevColumns.AllUsed)

	// insert move-to-terminal step if current column != "Terminal", and this is not the very first step
	if prevColumns.Focus != result.TerminalColumnType && prevColumns.Focus != result.NoColumnType {
		moveToTerminalStep := moveToTerminalStep(r, finder, currentColumns)
		steps = append(steps, moveToTerminalStep)
	}

	if r.Type == TerminalCommand {
		// command step
		cmdStep := terminalCommandStep(r, finder, currentColumns)
		steps = append(steps, cmdStep)

		// cd step
		if strings.HasPrefix(r.Text, "cd ") {
			cmdStep := terminalCdStep(r, finder, currentColumns)
			steps = append(steps, cmdStep)
		}
	} else if r.Type == TerminalOutput {
		outputStep := terminalOutputStep(r, finder, currentColumns)
		steps = append(steps, outputStep)
	}

	return steps
}

func toTerminalSteps(
	r *Row,
	finder *StepIdFinder,
	prevColumns *ColumnInfo,
) ([]result.Step, *ColumnInfo, error) {
	// current columns update
	currentColumns := &ColumnInfo{
		AllUsed: appendIfNotExists(prevColumns.AllUsed, result.TerminalColumnType),
		Focus:   result.TerminalColumnType,
	}

	subType, err := toTerminalSubType(r.Type)
	if err != nil {
		return nil, nil, fmt.Errorf("toTerminalSubType failed, %s", err)
	}

	switch subType {
	case TerminalCommand:
		// row -> specific row
		terminalRow, err := toTerminalCommandRow(r)
		if err != nil {
			return nil, nil, fmt.Errorf("toTerminalSteps failed, %s", err)
		}

		// specific row -> step
		steps := breakdownTerminalRow(terminalRow, finder, prevColumns)
		if err != nil {
			return nil, nil, fmt.Errorf("toTerminalSteps failed, %s", err)
		}
		return steps, currentColumns, nil

	case TerminalOutput:
		// row -> specific row
		terminalRow, err := toTerminalOutputRow(r)
		if err != nil {
			return nil, nil, fmt.Errorf("toTerminalSteps failed, %s", err)
		}

		// specific row -> step
		steps := breakdownTerminalRow(terminalRow, finder, prevColumns)
		if err != nil {
			return nil, nil, fmt.Errorf("toTerminalSteps failed, %s", err)
		}
		return steps, currentColumns, nil

	case TerminalOpen:
		// row -> specific row
		terminalRow, err := toTerminalOpenRow(r)
		if err != nil {
			return nil, nil, fmt.Errorf("toTerminalSteps failed, %s", err)
		}

		// specific row -> step
		steps := breakdownTerminalRow(terminalRow, finder, prevColumns)
		if err != nil {
			return nil, nil, fmt.Errorf("toTerminalSteps failed, %s", err)
		}
		return steps, currentColumns, nil

	default:
		return nil, nil, fmt.Errorf("toTerminalSteps failed, type = '%s' not implemented", r.Type)
	}
}
