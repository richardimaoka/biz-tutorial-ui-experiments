package state

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type TerminalColumn struct {
	terminal Terminal
}

func NewTerminalColumn() *TerminalColumn {
	return &TerminalColumn{}
}

func (t *TerminalColumn) MarkAllExecuted() {
	t.terminal.MarkAllExecuted()
}

func (t *TerminalColumn) Transition(elemType TerminalElementType, text string) {
	switch elemType {
	case TerminalTypeCommand:
		t.terminal.WriteCommand(text)
	case TerminalTypeOutput:
		t.terminal.WriteOutput(text)
	}
}

func (t *TerminalColumn) Cd(dir string) {
	t.terminal.ChangeCurrentDirectory(dir)
}

func (c *TerminalColumn) ToGraphQLTerminalColumn() *model.TerminalColumn {
	return &model.TerminalColumn{
		Terminal: c.terminal.ToGraphQLTerminal(),
	}
}
