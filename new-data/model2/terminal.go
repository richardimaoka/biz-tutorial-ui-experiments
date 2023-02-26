package model2

import (
	"fmt"
	"reflect"
)

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
func (t *Terminal) writeOutput(output string) {
	node := TerminalNode{
		Content: TerminalOutput{
			Output: &output,
		},
	}

	// works even if Nodes is nil
	t.Nodes = append(t.Nodes, &node)
}

// assuming TypeInCommand() is called earlier
func (t *Terminal) Execute(action ActionTerminal) error {
	//action.validate()
	//t.validate(action)
	if err := t.markLastCommandExecuted(); err != nil {
		return err
	}

	if action.Output != "" {
		t.writeOutput(action.Output)
	}

	if action.CurrentDirectory != "" {
		t.ChangeCurrentDirectory(action.CurrentDirectory)
	}

	return nil
}
