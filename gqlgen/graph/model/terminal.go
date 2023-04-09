package model

import (
	"fmt"
	"reflect"
)

func (t *Terminal) getLastNode() (*TerminalNode, error) {
	if len(t.Nodes) == 0 {
		return nil, fmt.Errorf("getLastNode failed, terminal has zero nodes")
	}

	lastNode := t.Nodes[len(t.Nodes)-1]
	if lastNode == nil {
		return nil, fmt.Errorf("getLastNode failed, terminal' last node = nil")
	}

	return lastNode, nil
}

func (t *Terminal) canTypeInCommand() error {
	if len(t.Nodes) == 0 {
		return nil //allow typing in initial command
	}

	lastNode, err := t.getLastNode()
	if err != nil {
		return fmt.Errorf("cannot type in command, failed to find terminal's last node, %s", err)
	}

	switch v := lastNode.Content.(type) {
	case TerminalCommand:
		if *v.BeforeExecution {
			return fmt.Errorf("cannot type in command, last command has beforeExecution = true")
		}
		return nil
	case TerminalOutput:
		return nil
	default:
		return nil
	}
}

func (t *Terminal) canExecuteCommand(command string) error {
	lastNode, err := t.getLastNode()
	if err != nil {
		return fmt.Errorf("cannot execute command, failed get to terminal's last node, %s", err)
	}

	cmd, ok := lastNode.Content.(TerminalCommand)
	if !ok {
		return fmt.Errorf("cannot execute command, terminal's last node is not TerminalCommand but %v", reflect.TypeOf(lastNode.Content))
	}

	if *cmd.Command != command {
		return fmt.Errorf("cannot execute command, terminal's last command = %s, not %s", *cmd.Command, command)
	}

	if cmd.BeforeExecution == nil || *cmd.BeforeExecution == false {
		return fmt.Errorf("cannot execute command, terminal's last command is not ready for execution")
	}

	return nil
}

func (t *Terminal) typeInCommand(command string) {
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

func (t *Terminal) changeCurrentDirectory(filePath string) {
	t.CurrentDirectory = &filePath
}

func (t *Terminal) writeOutput(output string) {
	node := TerminalNode{
		Content: TerminalOutput{
			Output: &output,
		},
	}

	// works even if Nodes is nil
	t.Nodes = append(t.Nodes, &node)
}

// precondition: canMarkLastCommandExecuted(command) is called
func (t *Terminal) markCommandExecuted(command string) {
	lastNode, err := t.getLastNode()
	if err != nil {
		fmt.Println("**********************************this happeneddd unexpectedly!!!********")
		return //if canMarkLastCommandExecuted(command) is successfully called, this should never happen
	}

	falseValue := false
	lastNode.Content = TerminalCommand{Command: &command, BeforeExecution: &falseValue}
}

func (t *Terminal) executeCommand(command string, output, currentDirectory *string) {
	t.markCommandExecuted(command)
	if currentDirectory != nil {
		t.changeCurrentDirectory(*currentDirectory)
	}
	if output != nil {
		t.writeOutput(*output)
	}
}

// public methods

func NewTerminal(name string) *Terminal {
	return &Terminal{
		Name: &name,
	}
}

func (t *Terminal) TypeInCommand(command string) error {
	if err := t.canTypeInCommand(); err != nil {
		return fmt.Errorf("TypeInCommand failed, %s", err)
	}

	t.typeInCommand(command)
	return nil
}

func (t *Terminal) ExecuteCommand(command string, currentDirectory, output *string) error {
	if err := t.canExecuteCommand(command); err != nil {
		return fmt.Errorf("ExecuteCommand failed, %s", err)
	}

	t.executeCommand(command, currentDirectory, output)
	return nil
}
