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

//TODO: directry translate ActionCommand to Terminal by invoking Termnal's methods
//TODO: when action_command_test.go is finished, combine AddFile, UpdateFile, .... into ActionCommand
