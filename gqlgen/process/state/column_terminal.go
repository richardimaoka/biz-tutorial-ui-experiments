package state

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

/**
 * Terminal Column type and methods
 */

type TerminalColumn struct {
	terminals []*Terminal
}

func NewTerminalColumn() *TerminalColumn {
	return &TerminalColumn{}
}

func (c *TerminalColumn) getOrCreateTerminal(name string) *Terminal {
	for _, t := range c.terminals {
		if t.terminalName == name {
			return t
		}
	}

	terminal := NewTerminal(name)
	c.terminals = append(c.terminals, terminal)

	return terminal
}

func (c *TerminalColumn) TerminalCommand(
	stepId string,
	fields *TerminalFields,
) error {
	terminal := c.getOrCreateTerminal(fields.TerminalName)
	terminal.WriteCommand(stepId, fields.TerminalText)
	if fields.TerminalTooltipContents == "" {
		terminal.ClearTooltip()
	} else {
		terminal.SetTooltip(fields.TerminalTooltipContents, TERMINAL_TOOLTIP_START)
	}
	return nil
}

func (c *TerminalColumn) TerminalCommandExecute(
	fields *TerminalFields,
) error {
	terminal := c.getOrCreateTerminal(fields.TerminalName)
	err := terminal.ExecuteLastCommand()
	if err != nil {
		return fmt.Errorf("TerminalColumn TerminalCommandExecute() failed, %s", err)
	}

	return nil
}

func (c *TerminalColumn) TerminalOutput(
	stepId string,
	fields *TerminalFields,
) error {
	terminal := c.getOrCreateTerminal(fields.TerminalName)
	terminal.WriteOutput(stepId, fields.TerminalText)
	if fields.TerminalTooltipContents == "" {
		terminal.ClearTooltip()
	} else {
		terminal.SetTooltip(fields.TerminalTooltipContents, TERMINAL_TOOLTIP_START)
	}
	return nil
}

func (c *TerminalColumn) TerminalCd(
	name string,
	fields *TerminalFields,
) error {
	terminal := c.getOrCreateTerminal(name)
	terminal.ChangeCurrentDirectory(fields.CurrentDir)
	return nil
}

func (c *TerminalColumn) Update(stepId string, fields *TerminalFields) error {
	var err error

	switch fields.TerminalStepType {
	case TerminalOutput:
		err = c.TerminalOutput(stepId, fields)
	case TerminalCommand:
		err = c.TerminalCommand(stepId, fields)
	case TerminalCommandExecuted:
		err = c.TerminalCommandExecute(fields)
	case TerminalCd:
		err = c.TerminalCd(fields.TerminalName, fields)
	case TerminalMove:
		// no update is needed, just changing FocusColumn is fine
	case TerminalOpen:
		// no update is needed, just changing FocusColumn is fine
	default:
		err = fmt.Errorf("type = '%s' is not implemented", fields.TerminalStepType)
	}

	if err != nil {
		return fmt.Errorf("TerminalColumn Update() failed, %s", err)
	}

	return nil
}

func (c *TerminalColumn) ToGraphQL() *model.TerminalColumn {
	var terminals []*model.Terminal
	for _, t := range c.terminals {
		terminals = append(terminals, t.ToGraphQL())
	}

	return &model.TerminalColumn{
		Terminals: terminals,
	}
}

func (c *TerminalColumn) ToGraphQLColumnWrapper() *model.ColumnWrapper {
	return &model.ColumnWrapper{
		Column:            c.ToGraphQL(),
		ColumnName:        "Terminal",
		ColumnDisplayName: stringRef("terminal"),
	}
}
