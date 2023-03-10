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

func (t *Terminal) canMarkLastCommandExecuted(command string) (*TerminalNode, error) {
	lastNode, err := t.getLastNode()
	if err != nil {
		return nil, fmt.Errorf("failed get to terminal's last node, %s", err)
	}

	cmd, ok := lastNode.Content.(TerminalCommand)
	if !ok {
		return nil, fmt.Errorf("terminal's last node is not TerminalCommand but %v", reflect.TypeOf(lastNode.Content))
	}

	if *cmd.Command != command {
		return nil, fmt.Errorf("terminal's last command = %s, not %s", *cmd.Command, command)
	}

	if cmd.BeforeExecution == nil || *cmd.BeforeExecution == false {
		return nil, fmt.Errorf("terminal's last command is not ready for execution")
	}

	return lastNode, nil
}

func (t *Terminal) canWriteOutput() error {
	if len(t.Nodes) == 0 {
		return nil // allow writing initial output
	}

	lastNode, err := t.getLastNode()
	if err != nil {
		return fmt.Errorf("cannot write output, failed to find terminal's last node, %s", err)
	}

	switch v := lastNode.Content.(type) {
	case TerminalCommand:
		if *v.BeforeExecution {
			return fmt.Errorf("cannot write output, last command has beforeExecution = true")
		}
		return nil
	case TerminalOutput:
		return nil
	default:
		return nil
	}
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

func (t *Terminal) writeOutput(output string) {
	node := TerminalNode{
		Content: TerminalOutput{
			Output: &output,
		},
	}

	// works even if Nodes is nil
	t.Nodes = append(t.Nodes, &node)
}

// public methods

func NewTerminal(name string) *Terminal {
	return &Terminal{
		Name: &name,
	}
}

//no pre-condition required, always succeed
func (t *Terminal) ChangeCurrentDirectory(filePath string) {
	t.CurrentDirectory = &filePath
}

func (t *Terminal) TypeInCommand(command string) error {
	if err := t.canTypeInCommand(); err != nil {
		return fmt.Errorf("TypeInCommand failed, %s", err)
	}

	t.typeInCommand(command)
	return nil
}

func (t *Terminal) WriteOutput(output string) error {
	if err := t.canWriteOutput(); err != nil {
		return fmt.Errorf("WriteOutput failed, %s", err)
	}

	t.writeOutput(output)
	return nil
}

func (t *Terminal) MarkLastCommandExecuted(command string) error {
	lastNode, err := t.canMarkLastCommandExecuted(command)
	if err != nil {
		return fmt.Errorf("MarkLastCommandExecuted failed, %s", err)
	}

	lastNode.markCommandExecuted(command)
	return nil
}
