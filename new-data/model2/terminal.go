package model2

import (
	"fmt"
	"reflect"
)

func newTerminal(name string) *Terminal {
	return &Terminal{
		Name: &name,
	}
}

func (t *Terminal) getLastNode() (*TerminalNode, error) {
	if len(t.Nodes) == 0 {
		return nil, fmt.Errorf("terminal has zero nodes")
	}

	lastNode := t.Nodes[len(t.Nodes)-1]
	if lastNode == nil {
		return nil, fmt.Errorf("terminal' last node = nil")
	}

	return lastNode, nil
}

// pre-condition check = isLastCommandExecutable()
func (t *Terminal) markLastCommandExecuted() error {
	lastNode, err := t.getLastNode()
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	lastCommand, ok := lastNode.Content.(TerminalCommand)
	if !ok {
		return fmt.Errorf("terminal's last node is not TerminalCommand but %v", reflect.TypeOf(lastNode.Content))
	}

	lastNode.Content = lastCommand.toExecutedCommand()
	return nil
}
