package model

import "testing"

func Test_typeInCommand(t *testing.T) {
	result := NewPageState()

	action := ActionTerminal{Command: "cd abc", TerminalName: "default", CurrentDirectory: "abc"}
	if err := result.typeIn(&action); err != nil {
		t.Error(err)
		return
	}

	compareAfterMarshal(t, "testdata/terminal-cd-type-in.json", result)
}

func Test_executeCd(t *testing.T) {
	result := NewPageState()

	action := ActionTerminal{Command: "cd abc", TerminalName: "default", CurrentDirectory: "abc"}
	if err := result.typeIn(&action); err != nil {
		t.Error(err)
		return
	}
	if err := result.executeActionTerminal(&action); err != nil {
		t.Error(err)
		return
	}

	compareAfterMarshal(t, "testdata/terminal-cd-executed.json", result)
}
