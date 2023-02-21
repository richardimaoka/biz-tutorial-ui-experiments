package model

import "testing"

func Test_typeInFailure(t *testing.T) {
	result := NewPageState()
	compareAfterMarshal(t, "testdata/new-page.json", result)

	wrongName := "wrong_name"
	action := ActionTerminal{Command: "cd abc", TerminalName: wrongName, CurrentDirectory: "abc"}
	if err := result.typeIn(&action); err == nil {
		t.Errorf("error expected as terminal with name = %s not found", wrongName)
		return
	}

	// expected page state unchanged from initial page
	compareAfterMarshal(t, "testdata/new-page.json", result)
}

func Test_ActionCd(t *testing.T) {
	result := NewPageState()
	compareAfterMarshal(t, "testdata/new-page.json", result)

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
	compareAfterMarshal(t, "testdata/new-page.json", result)

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

func Test_ActionNoOutput(t *testing.T) {
	result := NewPageState()
	compareAfterMarshal(t, "testdata/new-page.json", result)

	action := ActionTerminal{Command: "sleep 1", TerminalName: "default", Output: ""}
	if err := result.typeIn(&action); err != nil {
		t.Error(err)
		return
	}
	compareAfterMarshal(t, "testdata/terminal-no-output-type-in.json", result)

	if err := result.executeActionTerminal(&action); err != nil {
		t.Error(err)
		return
	}

	compareAfterMarshal(t, "testdata/terminal-no-output-executed.json", result)
}
