package model2

import (
	"testing"
)

func TestActionCommand1(t *testing.T) {
	cmd := ActionCommand{TerminalName: "default", Command: "mkdir hello"}
	compareAfterMarshal(t, "testdata/action/command/action_command1.json", cmd)
}

func TestActionCommand2(t *testing.T) {
	output := "abc"
	cmd := ActionCommand{TerminalName: "default", Command: "echo abc", Output: &output}
	compareAfterMarshal(t, "testdata/action/command/action_command2.json", cmd)
}

func TestActionCommand3(t *testing.T) {
	changeDirectory := "hello/world"
	cmd := ActionCommand{TerminalName: "another", Command: "cd hello/world", CurrentDirectory: &changeDirectory}
	compareAfterMarshal(t, "testdata/action/command/action_command3.json", cmd)
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
	terminal := NewTerminal("default")

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
