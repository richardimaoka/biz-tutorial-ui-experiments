package model2

import (
	"fmt"
	"reflect"
)

type ActionTerminal struct {
	Command          string
	TerminalName     string
	Output           string //if "", no output after execution
	CurrentDirectory string //if "", current directory is not changed after execution
}

//no pre-condition required, always succeed
func (t *Terminal) typeIn(action ActionTerminal) {
	trueValue := true
	node := TerminalNode{
		Content: TerminalCommand{
			Command:         &action.Command,
			BeforeExecution: &trueValue,
		},
	}

	// works even if Nodes is nil
	t.Nodes = append(t.Nodes, &node)
}

//no pre-condition required, always succeed
func (t *Terminal) changeDirectory(action ActionTerminal) {
	t.CurrentDirectoryPath = &action.CurrentDirectory
}

//no pre-condition required, always succeed
func (t *Terminal) writeOutput(action ActionTerminal) {
	node := TerminalNode{
		Content: TerminalOutput{
			Output: &action.Output,
		},
	}

	// works even if Nodes is nil
	t.Nodes = append(t.Nodes, &node)
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

// assuming typeIn() is called earlier
func (t *Terminal) Execute(action ActionTerminal) error {
	//action.validate()
	//t.validate(action)
	if err := t.markLastCommandExecuted(); err != nil {
		return err
	}

	if action.Output != "" {
		t.writeOutput(action)
	}

	if action.CurrentDirectory != "" {
		t.changeDirectory(action)
	}

	return nil
}
