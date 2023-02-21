package model

import (
	"fmt"
	"reflect"
)

func newTerminal(name string) *Terminal {
	return &Terminal{
		Name: &name,
	}
}

//no pre-condition required, always succeed
func (t *Terminal) typeInCommand(command *ActionCommand) {
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

//no pre-condition required, always succeed
func (t *Terminal) changeCurrentDirectory(cd UpdateTerminal) {
	t.CurrentDirectoryPath = &cd.CurrentDirectoryPath
}

//no pre-condition required, always succeed
func (t *Terminal) writeOutput(command *ActionCommand) {
	node := TerminalNode{
		Content: TerminalOutput{
			Output: &command.UpdateTerminal.Output,
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

func (t *Terminal) isLastCommandExecutable() error {
	lastNode, err := t.getLastNode()
	if err != nil {
		return fmt.Errorf("failed to terminal's last node, %s", err)
	}

	cmd, ok := lastNode.Content.(TerminalCommand)
	if !ok {
		return fmt.Errorf("terminal's last node is not TerminalCommand but %v", reflect.TypeOf(lastNode.Content))
	}

	if cmd.BeforeExecution == nil || *cmd.BeforeExecution == false {
		return fmt.Errorf("terminal's last command is not ready for execution")
	}

	return nil
}
