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
	CommandSubType TerminalSubType = "command"
	OutputSubType  TerminalSubType = "output"
)

func toTerminalSubType(s string) (TerminalSubType, error) {
	lower := strings.ToLower(s)

	switch lower {
	case string(CommandSubType):
		return CommandSubType, nil
	case string(OutputSubType):
		return OutputSubType, nil
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

/**
 * TerminalRow type(s) and functions
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

func toTerminalRow(fromRow *Row) (*TerminalRow, error) {
	errorPrefix := "failed to convert to TerminalOutput"

	//
	// Check column and type
	//
	if strings.ToLower(fromRow.Column) != TerminalType {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.Column)
	}
	subType, err := toTerminalSubType(fromRow.Type)
	if err != nil {
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

/**
 * Function(s) to convert a row to a step
 */
func terminalCommandStep(r *TerminalRow, StepIdFinder *StepIdFinder, usedColumns UsedColumns) result.Step {
	subId := "terminalCommandStep"
	stepId := StepIdFinder.StepIdFor(r.StepId, subId)

	step := result.Step{
		// fields to make the step searchable for re-generation
		IsFromRow:  true,
		ParentStep: r.StepId,
		SubID:      subId,
		// Other fields
		StepId:        stepId,
		Comment:       r.Comment,
		ModalContents: r.ModalContents,
		ColumnFields:  resultColumns(result.TerminalColumn, usedColumns),
		TerminalFields: result.TerminalFields{
			TerminalType: result.TerminalCommand,
			TerminalText: r.Text,
			TerminalName: r.TerminalName,
		},
	}

	if r.Tooltip != nil {
		step.TerminalTooltipContents = r.Tooltip.Contents
		step.TerminalTooltipTiming = r.Tooltip.Timing
	}

	return step
}

func terminalOutputStep(r *TerminalRow, finder *StepIdFinder, usedColumns UsedColumns) result.Step {
	subId := "terminalOutputStep"
	stepId := finder.StepIdFor(r.StepId, subId)

	step := result.Step{
		// Fields to make the step searchable for re-generation
		IsFromRow:  true,
		ParentStep: r.StepId,
		SubID:      subId,
		// Other fields
		StepId:        stepId,
		IsTrivial:     r.IsTrivial,
		Comment:       r.Comment,
		ModalContents: r.ModalContents,
		ColumnFields:  resultColumns(result.TerminalColumn, usedColumns),
		TerminalFields: result.TerminalFields{
			TerminalType: result.TerminalOutput,
			TerminalText: r.Text,
			TerminalName: r.TerminalName,
		},
	}
	if r.Tooltip != nil {
		step.TerminalTooltipContents = r.Tooltip.Contents
		step.TerminalTooltipTiming = r.Tooltip.Timing
	}

	return step
}

func moveToTerminalStep(r *TerminalRow, finder *StepIdFinder, usedColumns UsedColumns) result.Step {
	subId := "moveToTerminalStep"
	stepId := finder.StepIdFor(r.StepId, subId)

	step := result.Step{
		// Fields to make the step searchable for re-generation
		IsFromRow:  true,
		ParentStep: r.StepId,
		SubID:      subId,
		// Other fields
		StepId:        stepId,
		IsTrivial:     true, // always trivial
		Comment:       "(move to Terminal)",
		ModalContents: r.ModalContents,
		ColumnFields:  resultColumns(result.TerminalColumn, usedColumns),
		TerminalFields: result.TerminalFields{
			TerminalType: result.TerminalMove,
			TerminalName: r.TerminalName,
		},
	}
	// No tooltip - move step should be trivial and no tooltip to show

	return step
}

func terminalCdStep(r *TerminalRow, StepIdFinder *StepIdFinder, usedColumns UsedColumns) result.Step {
	currentDir := strings.TrimPrefix(r.Text, "cd ")

	subId := "terminalCdStep"
	stepId := StepIdFinder.StepIdFor(r.StepId, subId)

	step := result.Step{
		// Fields to make the step searchable for re-generation
		IsFromRow:  true,
		ParentStep: r.StepId,
		SubID:      subId,
		// other fields
		StepId:        stepId,
		IsTrivial:     true, // always trivial
		Comment:       "",
		ModalContents: r.ModalContents,
		// Columns
		ColumnFields: resultColumns(result.TerminalColumn, usedColumns),
		// Terminal fields
		TerminalFields: result.TerminalFields{
			CurrentDir:   currentDir,
			TerminalType: result.TerminalCd,
			TerminalName: r.TerminalName,
		},
	}

	// No tooltip - cd command should be trivial and no tooltip to show

	return step
}

/**
 * Function(s) to break down a row to steps
 */
func breakdownTerminalRow(
	r *TerminalRow,
	finder *StepIdFinder,
	prevColumns ColumnInfo,
) ([]result.Step, ColumnInfo) {

	// - step creation
	var steps []result.Step

	// insert move-to-terminal step if current column != "Terminal"
	if prevColumns.Focus != result.TerminalColumn && prevColumns.Focus != result.NoColumn {
		moveToTerminalStep := moveToTerminalStep(r, finder, prevColumns.AllUsed)
		steps = append(steps, moveToTerminalStep)
	}

	if r.Type == CommandSubType {
		// command step
		cmdStep := terminalCommandStep(r, finder, prevColumns.AllUsed)
		steps = append(steps, cmdStep)

		// cd step
		if strings.HasPrefix(r.Text, "cd ") {
			cmdStep := terminalCdStep(r, finder, prevColumns.AllUsed)
			steps = append(steps, cmdStep)
		}
	} else if r.Type == OutputSubType {
		outputStep := terminalOutputStep(r, finder, prevColumns.AllUsed)
		steps = append(steps, outputStep)
	}

	currentColumns := ColumnInfo{
		AllUsed: appendIfNotExists(prevColumns.AllUsed, result.TerminalColumn),
		Focus:   result.TerminalColumn,
	}
	return steps, currentColumns
}

func toTerminalSteps(
	r *Row,
	finder *StepIdFinder,
	prevColumns ColumnInfo,
) ([]result.Step, ColumnInfo, error) {
	terminalRow, err := toTerminalRow(r)
	if err != nil {
		return nil, prevColumns, err
	}

	breakdowns, currentColumns := breakdownTerminalRow(terminalRow, finder, prevColumns)
	return breakdowns, currentColumns, nil
}
