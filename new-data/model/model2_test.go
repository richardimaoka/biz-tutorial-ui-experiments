package model

import "testing"

func Test_typeInCommand(t *testing.T) {
	result := NewPageState()

	action := ActionTerminal{Command: "cd abc", TerminalName: "default", CurrentDirectory: "abc"}
	if err := result.typeIn(&action); err != nil {
		t.Error(err)
		return
	}

	compareAfterMarshal(t, "testdata/type-in-command.json", result)
}
