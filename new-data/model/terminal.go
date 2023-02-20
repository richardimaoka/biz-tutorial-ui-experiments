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

func (t *Terminal) getLastCommand() (*TerminalCommand, error) {
	if len(t.Nodes) == 0 {
		return nil, fmt.Errorf("terminal has zero nodes")
	}

	lastNode := t.Nodes[len(t.Nodes)-1]
	if lastNode == nil {
		return nil, fmt.Errorf("terminal' last node = nil")
	}

	//content is interface, possibly nil
	content := lastNode.Content
	lastCommand, ok := content.(TerminalCommand)
	if !ok {
		return nil, fmt.Errorf("terminal's last node's content is not TerminalCommand but %v", reflect.TypeOf(lastNode.Content))
	}

	return &lastCommand, nil
}
