package model2

import (
	"fmt"
	"reflect"
)

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

func (t *Terminal) isLastCommandExecutable(command string) error {
	lastNode, err := t.getLastNode()
	if err != nil {
		return fmt.Errorf("failed to terminal's last node, %s", err)
	}

	cmd, ok := lastNode.Content.(TerminalCommand)
	if !ok {
		return fmt.Errorf("terminal's last node is not TerminalCommand but %v", reflect.TypeOf(lastNode.Content))
	}

	if *cmd.Command != command {
		return fmt.Errorf("terminal's last command = %s, not %s", *cmd.Command, command)
	}

	if cmd.BeforeExecution == nil || *cmd.BeforeExecution == false {
		return fmt.Errorf("terminal's last command is not ready for execution")
	}

	return nil
}

// public methods

func NewTerminal(name string) *Terminal {
	return &Terminal{
		Name: &name,
	}
}

//no pre-condition required, always succeed
func (t *Terminal) TypeInCommand(command string) {
	trueValue := true
	node := TerminalNode{
		Content: TerminalCommand{
			Command:         &command,
			BeforeExecution: &trueValue,
		},
	}

	// works even if Nodes is nil
	t.Nodes = append(t.Nodes, &node)
}

//no pre-condition required, always succeed
func (t *Terminal) ChangeCurrentDirectory(filePath string) {
	t.CurrentDirectoryPath = &filePath
}

//no pre-condition required, always succeed
func (t *Terminal) WriteOutput(output string) {
	node := TerminalNode{
		Content: TerminalOutput{
			Output: &output,
		},
	}

	// works even if Nodes is nil
	t.Nodes = append(t.Nodes, &node)
}

func (t *Terminal) MarkLastCommandExecuted(command string) error {
	if err := t.isLastCommandExecutable(command); err != nil {
		return fmt.Errorf("MarkLastCommandExecuted failed, %s", err)
	}

	lastNode, err := t.getLastNode()
	if err != nil {
		return fmt.Errorf("MarkLastCommandExecuted failed, %s", err)
	}

	falseValue := false
	lastNode.Content = TerminalCommand{Command: &command, BeforeExecution: &falseValue}
	return nil
}
