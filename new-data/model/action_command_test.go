package model

import (
	"encoding/json"
	"os"
	"testing"
)

func address(s string) *string {
	return &s
}

func TestActionCommandMarshal(t *testing.T) {
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
		t.Run("test_action_command_marshal", func(t *testing.T) {
			compareAfterMarshal(t, e.expectedFile, e.command)
		})
	}
}

func TestActionCommandUnmarshal(t *testing.T) {
	files := []string{
		"testdata/action/command/action_command1.json",
		"testdata/action/command/action_command2.json",
		"testdata/action/command/action_command3.json",
	}

	for _, f := range files {
		t.Run("test_action_command_unmarshal", func(t *testing.T) {
			jsonBytes, err := os.ReadFile(f)
			if err != nil {
				t.Fatalf("failed to read %s", err)
			}

			var cmd ActionCommand
			if err := json.Unmarshal(jsonBytes, &cmd); err != nil {
				t.Fatalf("failed to unmarshal %s", err)
			}
			compareAfterMarshal(t, f, cmd)
		})
	}
}
