package model

import "testing"

func Test_typeInCommand(t *testing.T) {
	result := NewPageState()

	cmd := ActionTerminal{Command: "cd abc", TerminalName: "default", CurrentDirectory: "abc"}
	if err := result.typeInTerminalCommand(&cmd); err != nil {
		t.Error(err)
		return
	}

	compareAfterMarshal(t, "testdata/type-in-command.json", result)
}
