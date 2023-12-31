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
 * TerminalRow type(s) and functions
 */
type TerminalRow struct {
	RowId         string           `json:"rowId"`
	IsTrivial     bool             `json:"isTrivial"`
	Comment       string           `json:"comment"`
	Type          TerminalSubType  `json:"type"`
	Text          string           `json:"text"`
	Tooltip       *TerminalTooltip `json:"tooltip"`
	ModalContents string           `json:"modalContents"`
	ModalPosition string           `json:"modalPosition"`
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
	column, err := toColumnType(fromRow.RowType)
	if err != nil || column != TerminalColumn {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.RowType)
	}
	subType, err := toTerminalSubType(fromRow.SubType)
	if err != nil || subType != TerminalCommand {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, fromRow.SubType)
	}

	//
	// Check contents fields
	//
	if fromRow.Contents == "" {
		return nil, fmt.Errorf("%s, 'contents' is empty", errorPrefix)
	}

	//
	// Check terminal name
	//
	terminalName := fromRow.TerminalName
	if terminalName == "" {
		terminalName = "default"
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
	isTrivial := fromRow.Trivial.Value()

	return &TerminalRow{
		RowId:         fromRow.RowId,
		IsTrivial:     isTrivial,
		Comment:       fromRow.Comment,
		ModalContents: fromRow.ModalContents,
		ModalPosition: fromRow.ModalPosition,
		Type:          subType,
		Text:          fromRow.Contents,
		Tooltip:       terminalTooltip,
		TerminalName:  terminalName,
	}, nil
}

func toTerminalOutputRow(fromRow *Row) (*TerminalRow, error) {
	errorPrefix := "failed in toTerminalOutputRow"

	//
	// Check column and type
	//
	column, err := toColumnType(fromRow.RowType)
	if err != nil || column != TerminalColumn {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.RowType)
	}
	subType, err := toTerminalSubType(fromRow.SubType)
	if err != nil || subType != TerminalOutput {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, fromRow.SubType)
	}

	//
	// Check contents fields
	//
	if fromRow.Contents == "" {
		return nil, fmt.Errorf("%s, 'contents' is empty", errorPrefix)
	}

	//
	// Check terminal name
	//
	terminalName := fromRow.TerminalName
	if terminalName == "" {
		terminalName = "default"
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
	isTrivial := fromRow.Trivial.Value()

	return &TerminalRow{
		RowId:         fromRow.RowId,
		IsTrivial:     isTrivial,
		Comment:       fromRow.Comment,
		ModalContents: fromRow.ModalContents,
		Type:          subType,
		Text:          fromRow.Contents,
		Tooltip:       terminalTooltip,
		TerminalName:  terminalName,
	}, nil
}

func toTerminalOpenRow(fromRow *Row) (*TerminalRow, error) {
	errorPrefix := "failed in toTerminalOpenRow"

	//
	// Check column and type
	//
	column, err := toColumnType(fromRow.RowType)
	if err != nil || column != TerminalColumn {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.RowType)
	}
	subType, err := toTerminalSubType(fromRow.SubType)
	if err != nil {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, fromRow.SubType)
	}

	//
	// Check terminal name
	//
	terminalName := fromRow.TerminalName
	if terminalName == "" {
		terminalName = "default"
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
	isTrivial := fromRow.Trivial.Value()

	return &TerminalRow{
		RowId:         fromRow.RowId,
		IsTrivial:     isTrivial,
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
func terminalCommandStep(r *TerminalRow, StepIdFinder *StepIdFinder) state.Step {
	subId := "terminalCommandStep"
	stepId := StepIdFinder.StepIdFor(r.RowId, subId)

	step := state.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: state.FromRowFields{
			IsFromRow:   true,
			ParentRowId: r.RowId,
			SubID:       subId,
		},
		IntrinsicFields: state.IntrinsicFields{
			StepId:      stepId,
			Comment:     r.Comment,
			Mode:        state.HandsonMode,
			FocusColumn: state.TerminalColumnType,
		},
		AnimationFields: state.AnimationFields{
			IsTrivial: r.IsTrivial,
		},
		ModalFields: state.ModalFields{
			ModalContents: r.ModalContents,
		},
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

func terminalCommandExecutedStep(r *TerminalRow, finder *StepIdFinder) state.Step {
	subId := "terminalCommandExecutedStep"
	stepId := finder.StepIdFor(r.RowId, subId)

	step := state.Step{
		// Fields to make the step searchable for re-generation
		FromRowFields: state.FromRowFields{
			IsFromRow:   true,
			ParentRowId: r.RowId,
			SubID:       subId,
		},
		IntrinsicFields: state.IntrinsicFields{
			StepId:      stepId,
			Comment:     r.Comment,
			Mode:        state.HandsonMode,
			FocusColumn: state.TerminalColumnType,
		},
		AnimationFields: state.AnimationFields{
			IsTrivial: true,
		},
		TerminalFields: state.TerminalFields{
			TerminalStepType: state.TerminalCommandExecuted,
			TerminalName:     r.TerminalName,
		},
	}
	return step
}

func terminalCdStep(r *TerminalRow, StepIdFinder *StepIdFinder) state.Step {
	currentDir := strings.TrimPrefix(r.Text, "cd ")

	subId := "terminalCdStep"
	stepId := StepIdFinder.StepIdFor(r.RowId, subId)

	step := state.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: state.FromRowFields{
			IsFromRow:   true,
			ParentRowId: r.RowId,
			SubID:       subId,
		},
		IntrinsicFields: state.IntrinsicFields{
			StepId:      stepId,
			Comment:     "",
			Mode:        state.HandsonMode,
			FocusColumn: state.TerminalColumnType,
		},
		AnimationFields: state.AnimationFields{
			IsTrivial: true, //always true
		},
		// No ModalFields, as it is a trivial step
		TerminalFields: state.TerminalFields{
			CurrentDir:       currentDir,
			TerminalStepType: state.TerminalCd,
			TerminalName:     r.TerminalName,
		},
	}

	// No tooltip - trivial step and no tooltip to show

	return step
}

func terminalOutputStep(r *TerminalRow, finder *StepIdFinder) state.Step {
	subId := "terminalOutputStep"
	stepId := finder.StepIdFor(r.RowId, subId)

	step := state.Step{
		// Fields to make the step searchable for re-generation
		FromRowFields: state.FromRowFields{
			IsFromRow:   true,
			ParentRowId: r.RowId,
			SubID:       subId,
		},
		IntrinsicFields: state.IntrinsicFields{
			StepId:      stepId,
			Comment:     r.Comment,
			Mode:        state.HandsonMode,
			FocusColumn: state.TerminalColumnType,
		},
		AnimationFields: state.AnimationFields{
			IsTrivial: r.IsTrivial,
		},
		ModalFields: state.ModalFields{
			ModalContents: r.ModalContents,
		},
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

func moveToTerminalStep(r *TerminalRow, finder *StepIdFinder) state.Step {
	subId := "moveToTerminalStep"
	stepId := finder.StepIdFor(r.RowId, subId)

	step := state.Step{
		// fields to make the step searchable for re-generation
		FromRowFields: state.FromRowFields{
			IsFromRow:   true,
			ParentRowId: r.RowId,
			SubID:       subId,
		},
		IntrinsicFields: state.IntrinsicFields{
			StepId:      stepId,
			Comment:     "(move to Terminal)",
			Mode:        state.HandsonMode,
			FocusColumn: state.TerminalColumnType,
		},
		AnimationFields: state.AnimationFields{
			IsTrivial: true, //always true
		},
		// No ModalFields, as it is a trivial step
		TerminalFields: state.TerminalFields{
			TerminalStepType: state.TerminalMove,
			TerminalName:     r.TerminalName,
		},
	}
	// No tooltip - move step should be trivial and no tooltip to show

	return step
}

// func terminalCleanUpStep(r *TerminalRow, StepIdFinder *StepIdFinder) state.Step {
// 	subId := "terminalCleanupStep"
// 	stepId := StepIdFinder.StepIdFor(r.RowId, subId)

// 	step := state.Step{
// 		// fields to make the step searchable for re-generation
// 		FromRowFields: state.FromRowFields{
// 			IsFromRow:   true,
// 			ParentRowId: r.RowId,
// 			SubID:       subId,
// 		},
// 		IntrinsicFields: state.IntrinsicFields{
// 			StepId:      stepId,
// 			Comment:     "",
// 			Mode:        state.HandsonMode,
// 			FocusColumn: state.TerminalColumnType,
// 		},
// 		AnimationFields: state.AnimationFields{
// 			IsTrivial: true, //always true
// 		},
// 		// No ModalFields, as it is a trivial step
// 		TerminalFields: state.TerminalFields{
// 			TerminalStepType: state.TerminalCd,
// 		},
// 	}

// 	// No tooltip - trivial step and no tooltip to show

// 	return step
// }

/**
 * Function(s) to break down a row to steps
 */
func breakdownTerminalRow(r *TerminalRow, finder *StepIdFinder, prevColumn state.ColumnType) []state.Step {
	// - step creation
	var steps []state.Step

	// insert move-to-terminal step if current column != "Terminal", and this is not the very first step
	if prevColumn != state.TerminalColumnType && prevColumn != state.NoColumnType {
		moveToTerminalStep := moveToTerminalStep(r, finder)
		steps = append(steps, moveToTerminalStep)
	}

	if r.Type == TerminalCommand {
		// command step
		cmdStep := terminalCommandStep(r, finder)
		steps = append(steps, cmdStep)

		// cd step
		if strings.HasPrefix(r.Text, "cd ") {
			cdStep := terminalCdStep(r, finder)
			steps = append(steps, cdStep)
		} else {
			executedStep := terminalCommandExecutedStep(r, finder)
			steps = append(steps, executedStep)
		}
	} else if r.Type == TerminalOutput {
		outputStep := terminalOutputStep(r, finder)
		steps = append(steps, outputStep)
	}

	// cleanup step
	// step := terminalCleanUpStep(r, finder)
	// steps = append(steps, step)

	return steps
}

func toTerminalSteps(
	r *Row,
	finder *StepIdFinder,
	prevColumn state.ColumnType,
) ([]state.Step, error) {
	subType, err := toTerminalSubType(r.SubType)
	if err != nil {
		return nil, fmt.Errorf("toTerminalSubType failed, %s", err)
	}

	switch subType {
	case TerminalCommand:
		// row -> specific row
		terminalRow, err := toTerminalCommandRow(r)
		if err != nil {
			return nil, fmt.Errorf("toTerminalSteps failed, %s", err)
		}

		// specific row -> step
		steps := breakdownTerminalRow(terminalRow, finder, prevColumn)
		if err != nil {
			return nil, fmt.Errorf("toTerminalSteps failed, %s", err)
		}
		return steps, nil

	case TerminalOutput:
		// row -> specific row
		terminalRow, err := toTerminalOutputRow(r)
		if err != nil {
			return nil, fmt.Errorf("toTerminalSteps failed, %s", err)
		}

		// specific row -> step
		steps := breakdownTerminalRow(terminalRow, finder, prevColumn)
		if err != nil {
			return nil, fmt.Errorf("toTerminalSteps failed, %s", err)
		}
		return steps, nil

	case TerminalOpen:
		// row -> specific row
		terminalRow, err := toTerminalOpenRow(r)
		if err != nil {
			return nil, fmt.Errorf("toTerminalSteps failed, %s", err)
		}

		// specific row -> step
		steps := breakdownTerminalRow(terminalRow, finder, prevColumn)
		if err != nil {
			return nil, fmt.Errorf("toTerminalSteps failed, %s", err)
		}
		return steps, nil

	default:
		return nil, fmt.Errorf("toTerminalSteps failed, type = '%s' not implemented", r.SubType)
	}
}
