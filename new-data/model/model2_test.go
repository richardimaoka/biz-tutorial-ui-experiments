package model

import "testing"

func Test_ActionCd(t *testing.T) {
	result := NewPageState()

	action := ActionTerminal{Command: "cd abc", TerminalName: "default", CurrentDirectory: "abc"}
	if err := result.typeIn(&action); err != nil {
		t.Error(err)
		return
	}
	compareAfterMarshal(t, "testdata/terminal-cd-type-in.json", result)

	if err := result.executeActionTerminal(&action); err != nil {
		t.Error(err)
		return
	}

	compareAfterMarshal(t, "testdata/terminal-cd-executed.json", result)
}

func Test_ActionEcho(t *testing.T) {
	result := NewPageState()

	action := ActionTerminal{Command: "echo hello", TerminalName: "default", Output: "hello"}
	if err := result.typeIn(&action); err != nil {
		t.Error(err)
		return
	}
	compareAfterMarshal(t, "testdata/terminal-echo-type-in.json", result)

	if err := result.executeActionTerminal(&action); err != nil {
		t.Error(err)
		return
	}

	compareAfterMarshal(t, "testdata/terminal-echo-executed.json", result)
}
