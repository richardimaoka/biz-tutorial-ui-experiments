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
	name string,
	command string,
	tooltipContents string,
) {
	terminal := c.getOrCreateTerminal(name)
	terminal.WriteCommand(stepId, command)
	if tooltipContents == "" {
		terminal.ClearTooltip()
	} else {
		terminal.SetTooltip(tooltipContents, TERMINAL_TOOLTIP_START)
	}
}

func (c *TerminalColumn) TerminalOutput(
	stepId string,
	name string,
	output string,
	tooltipContents string,
) {
	terminal := c.getOrCreateTerminal(name)
	terminal.WriteOutput(stepId, output)
	if tooltipContents == "" {
		terminal.ClearTooltip()
	} else {
		terminal.SetTooltip(tooltipContents, TERMINAL_TOOLTIP_START)
	}
}

func (c *TerminalColumn) TerminalCd(
	name string,
	dirPath string,
) {
	terminal := c.getOrCreateTerminal(name)
	terminal.ChangeCurrentDirectory(dirPath)
}

func (c *TerminalColumn) Update(stepId string, fields *TerminalFields) error {
	switch fields.TerminalStepType {
	case TerminalCommand:
		c.TerminalCommand(stepId, fields.TerminalName, fields.TerminalText, fields.TerminalTooltipContents)
	case TerminalOutput:
		c.TerminalOutput(stepId, fields.TerminalName, fields.TerminalText, fields.TerminalTooltipContents)
	case TerminalCd:
		c.TerminalCd(fields.TerminalName, fields.CurrentDir)
	case TerminalMove:
		// no update is needed, just changing FocusColumn is fine
	case TerminalOpen:
		// no update is needed, just changing FocusColumn is fine
	default:
		return fmt.Errorf("Update failed, type = '%s' is not implemented", fields.TerminalStepType)
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
