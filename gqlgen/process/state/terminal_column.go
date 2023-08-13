package state

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

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

func (t *TerminalColumn) Process(step, terminalType, terminalText, currentDir string) error {
	if terminalText == "" {
		t.MarkAllExecuted()
	} else {
		terminalType, err := ToTerminalElementType(terminalType)
		if err != nil {
			return fmt.Errorf("ToGraphQLPages failed at step = %s to convert terminal type, %s", step, err)
		}
		t.Transition(terminalType, terminalText)
	}

	if currentDir != "" {
		t.Cd(currentDir)
	}

	return nil
}

func (c *TerminalColumn) ToGraphQLTerminalColumn() *model.TerminalColumn {
	return &model.TerminalColumn{
		Terminal: c.terminal.ToGraphQLTerminal(),
	}
}
