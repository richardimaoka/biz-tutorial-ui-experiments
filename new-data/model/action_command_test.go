package model

import (
	"testing"
)

func address(s string) *string {
	return &s
}

func TestActionCommands(t *testing.T) {
	type Entry struct {
		command      ActionCommand
		expectedFile string
	}

	entries := []Entry{
		{expectedFile: "testdata/action/command/action_command1.json", command: ActionCommand{TerminalName: "default", Command: "mkdir hello"}},
		{expectedFile: "testdata/action/command/action_command2.json", command: ActionCommand{TerminalName: "default", Command: "echo abc", Output: address("abc")}},
		{expectedFile: "testdata/action/command/action_command3.json", command: ActionCommand{TerminalName: "another", Command: "cd hello/world", CurrentDirectory: address("hello/world")}},
	}

	for _, e := range entries {
		t.Run("test_action_command_marshaling", func(t *testing.T) {
			compareAfterMarshal(t, e.expectedFile, e.command)
		})
	}
}

func TestActionCommandTerminal1(t *testing.T) {
	cmd := ActionCommand{TerminalName: "default", Command: "mkdir hello"}
	terminal := NewTerminal("default")

	if err := terminal.TypeInCommand(cmd.Command); err != nil {
		t.Fatal(err)
	}
	if err := terminal.MarkLastCommandExecuted(cmd.Command); err != nil {
		t.Fatal(err)
	}

	compareAfterMarshal(t, "testdata/action/command/terminal1.json", terminal)
}

func TestActionCommandTerminal2(t *testing.T) {
	output := "abc"
	cmd := ActionCommand{TerminalName: "default", Command: "echo abc", Output: &output}
	terminal := NewTerminal("default")

	if err := terminal.TypeInCommand(cmd.Command); err != nil {
		t.Fatal(err)
	}
	if err := terminal.MarkLastCommandExecuted(cmd.Command); err != nil {
		t.Fatal(err)
	}
	if err := terminal.WriteOutput(*cmd.Output); err != nil {
		t.Fatal(err)
	}

	compareAfterMarshal(t, "testdata/action/command/terminal2.json", terminal)
}

func TestActionCommandTerminal3(t *testing.T) {
	changeDirectory := "hello/world"
	cmd := ActionCommand{TerminalName: "another", Command: "cd hello/world", CurrentDirectory: &changeDirectory}
	terminal := NewTerminal("another")

	if err := terminal.TypeInCommand(cmd.Command); err != nil {
		t.Fatal(err)
	}
	if err := terminal.MarkLastCommandExecuted(cmd.Command); err != nil {
		t.Fatal(err)
	}
	terminal.ChangeCurrentDirectory(*cmd.CurrentDirectory)

	compareAfterMarshal(t, "testdata/action/command/terminal3.json", terminal)
}

//TODO: when action_command_test.go is finished, combine AddFile, UpdateFile, .... into ActionCommand
