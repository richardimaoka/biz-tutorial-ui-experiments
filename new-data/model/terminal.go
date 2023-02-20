package model

import (
	"fmt"
	"reflect"
)

//no pre-condition required, always succeed
func (t *Terminal) typeInCommand(command *ActionCommand) {
	// append terminal node
	trueValue := true
	node := TerminalNode{
		Content: TerminalCommand{
			Command:         &command.Command,
			BeforeExecution: &trueValue,
		},
	}

	// works even if Nodes is nil
	t.Nodes = append(t.Nodes, &node)
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

func (t *Terminal) verifyLastCommand() error {
	lastNode, err := t.getLastNode()
	if err != nil {
		return fmt.Errorf("verifyLastCommand failed, %s", err)
	}

	_, ok := lastNode.Content.(TerminalCommand)
	if !ok {
		return fmt.Errorf("terminal's last node's content is not TerminalCommand but %v", reflect.TypeOf(lastNode.Content))
	}

	return nil
}

func (t *Terminal) markLastCommandExecuted() error {
	lastNode, err := t.getLastNode()
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	lastCommand, ok := lastNode.Content.(TerminalCommand)
	if !ok {
		return fmt.Errorf("terminal's last node's content is not TerminalCommand but %v", reflect.TypeOf(lastNode.Content))
	}

	lastNode.Content = lastCommand.toExecutedCommand()
	return nil
}
