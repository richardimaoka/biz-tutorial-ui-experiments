package model2

import (
	"testing"
)

func TestMarshalActionCommand(t *testing.T) {
	cmd := ActionCommand{ActionType: "TerminalCommand", TerminalName: "default"}

	compareAfterMarshal(t, "testdata/action/command/action_command0.json", cmd)

	if cmd.Output != "" {
		t.Fatalf("cmd.Output expected empty string but = %s", cmd.Output)
	}

	if cmd.CurrentDirectory != "" {
		t.Fatalf("cmd.CurrentDirectory expected empty string but = %s", cmd.CurrentDirectory)
	}
}
