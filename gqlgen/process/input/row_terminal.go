package input

import (
	"fmt"
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
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

type TerminalTooltipTiming string

const (
	TERMINAL_TOOLTIP_START TerminalTooltipTiming = "START"
	TERMINAL_TOOLTIP_END   TerminalTooltipTiming = "END"
)

type TerminalTooltip struct {
	Contents string                `json:"contents"`
	Timing   TerminalTooltipTiming `json:"timing"`
}

func toTerminalTooltipTiming(s string) (TerminalTooltipTiming, error) {
	switch strings.ToUpper(s) {
	case string(TERMINAL_TOOLTIP_START):
		return TERMINAL_TOOLTIP_START, nil
	case string(TERMINAL_TOOLTIP_END):
		return TERMINAL_TOOLTIP_END, nil
	case "": // default value is different from source tooltip
		return TERMINAL_TOOLTIP_START, nil
	default:
		return "", fmt.Errorf("TerminalTooltipTiming value = '%s' is invalid", s)
	}
}

func (t TerminalTooltipTiming) toState() state.TerminalTooltipTiming {
	switch t {
	case TERMINAL_TOOLTIP_START:
		return state.TERMINAL_TOOLTIP_START
	case TERMINAL_TOOLTIP_END:
		return state.TERMINAL_TOOLTIP_END
	default:
		panic(fmt.Sprintf("TerminalTooltipTiming has an invalid value = '%s'", t))
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
func terminalCommandStep(r *TerminalRow, StepIdFinder *StepIdFinder, currentColumns state.ColumnFields) state.Step {
	subId := "terminalCommandStep"
	stepId := StepIdFinder.StepIdFor(r.StepId, subId)

	step := state.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: state.FromRowFields{
			IsFromRow:  true,
			ParentStep: r.StepId,
			SubID:      subId,
		},
		IntrinsicFields: state.IntrinsicFields{
			StepId:  stepId,
			Comment: r.Comment,
		},
		AnimationFields: state.AnimationFields{
			IsTrivial: r.IsTrivial,
		},
		ModalFields: state.ModalFields{
			ModalContents: r.ModalContents,
		},
		ColumnFields: currentColumns,
		TerminalFields: state.TerminalFields{
			TerminalStepType: state.TerminalCommand,
			TerminalText:     r.Text,
			TerminalName:     r.TerminalName,
		},
	}

	if r.Tooltip != nil {
		step.TerminalTooltipContents = r.Tooltip.Contents
		step.TerminalTooltipTiming = r.Tooltip.Timing.toState()
	}

	return step
}

func terminalOutputStep(r *TerminalRow, finder *StepIdFinder, currentColumns state.ColumnFields) state.Step {
	subId := "terminalOutputStep"
	stepId := finder.StepIdFor(r.StepId, subId)

	step := state.Step{
		// Fields to make the step searchable for re-generation
		FromRowFields: state.FromRowFields{
			IsFromRow:  true,
			ParentStep: r.StepId,
			SubID:      subId,
		},
		IntrinsicFields: state.IntrinsicFields{
			StepId:  stepId,
			Comment: r.Comment,
		},
		AnimationFields: state.AnimationFields{
			IsTrivial: r.IsTrivial,
		},
		ModalFields: state.ModalFields{
			ModalContents: r.ModalContents,
		},
		ColumnFields: currentColumns,
		TerminalFields: state.TerminalFields{
			TerminalStepType: state.TerminalOutput,
			TerminalText:     r.Text,
			TerminalName:     r.TerminalName,
		},
	}
	if r.Tooltip != nil {
		step.TerminalTooltipContents = r.Tooltip.Contents
		step.TerminalTooltipTiming = r.Tooltip.Timing.toState()
	}

	return step
}

func moveToTerminalStep(r *TerminalRow, finder *StepIdFinder, currentColumns state.ColumnFields) state.Step {
	subId := "moveToTerminalStep"
	stepId := finder.StepIdFor(r.StepId, subId)

	step := state.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: state.FromRowFields{
			IsFromRow:  true,
			ParentStep: r.StepId,
			SubID:      subId,
		},
		IntrinsicFields: state.IntrinsicFields{
			StepId:  stepId,
			Comment: "(move to Terminal)",
		},
		AnimationFields: state.AnimationFields{
			IsTrivial: true, //always true
		},
		// No ModalFields, as it is a trivial step
		ColumnFields: currentColumns,
		TerminalFields: state.TerminalFields{
			TerminalStepType: state.TerminalMove,
			TerminalName:     r.TerminalName,
		},
	}
	// No tooltip - move step should be trivial and no tooltip to show

	return step
}

func terminalCdStep(r *TerminalRow, StepIdFinder *StepIdFinder, currentColumns state.ColumnFields) state.Step {
	currentDir := strings.TrimPrefix(r.Text, "cd ")

	subId := "terminalCdStep"
	stepId := StepIdFinder.StepIdFor(r.StepId, subId)

	step := state.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: state.FromRowFields{
			IsFromRow:  true,
			ParentStep: r.StepId,
			SubID:      subId,
		},
		IntrinsicFields: state.IntrinsicFields{
			StepId:  stepId,
			Comment: "",
		},
		AnimationFields: state.AnimationFields{
			IsTrivial: true, //always true
		},
		// No ModalFields, as it is a trivial step
		ColumnFields: currentColumns,
		TerminalFields: state.TerminalFields{
			CurrentDir:       currentDir,
			TerminalStepType: state.TerminalCd,
			TerminalName:     r.TerminalName,
		},
	}

	// No tooltip - trivial step and no tooltip to show

	return step
}

func terminalCleanUpStep(r *TerminalRow, StepIdFinder *StepIdFinder, currentColumns state.ColumnFields) state.Step {
	subId := "terminalCleanupStep"
	stepId := StepIdFinder.StepIdFor(r.StepId, subId)

	step := state.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: state.FromRowFields{
			IsFromRow:  true,
			ParentStep: r.StepId,
			SubID:      subId,
		},
		IntrinsicFields: state.IntrinsicFields{
			StepId:  stepId,
			Comment: "",
		},
		AnimationFields: state.AnimationFields{
			IsTrivial: true, //always true
		},
		// No ModalFields, as it is a trivial step
		ColumnFields: currentColumns,
		TerminalFields: state.TerminalFields{
			TerminalStepType: state.TerminalCd,
		},
	}

	// No tooltip - trivial step and no tooltip to show

	return step
}

/**
 * Function(s) to break down a row to steps
 */
func breakdownTerminalRow(r *TerminalRow, finder *StepIdFinder, prevColumns *ColumnInfo) []state.Step {
	// - step creation
	var steps []state.Step
	currentColumns := resultColumns(state.TerminalColumnType, prevColumns.AllUsed)

	// insert move-to-terminal step if current column != "Terminal", and this is not the very first step
	if prevColumns.Focus != state.TerminalColumnType && prevColumns.Focus != state.NoColumnType {
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

	// cleanup step
	step := terminalCleanUpStep(r, finder, currentColumns)
	steps = append(steps, step)

	return steps
}

func toTerminalSteps(
	r *Row,
	finder *StepIdFinder,
	prevColumns *ColumnInfo,
) ([]state.Step, *ColumnInfo, error) {
	// current columns update
	currentColumns := &ColumnInfo{
		AllUsed: appendIfNotExists(prevColumns.AllUsed, state.TerminalColumnType),
		Focus:   state.TerminalColumnType,
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
