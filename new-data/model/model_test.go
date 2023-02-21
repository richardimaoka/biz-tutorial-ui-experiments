package model

import (
	"testing"
)

func Test_NewPageState(t *testing.T) {
	result := NewPageState()
	compareAfterMarshal(t, "testdata/new-page.json", result)
}

func Test_typeInCommandSuccess(t *testing.T) {
	result := NewPageState()

	cmd := ActionCommand{Command: "mkdir abc", TerminalName: "default"}
	if err := result.typeInTerminalCommand(&cmd); err != nil {
		t.Error(err)
		return
	}

	compareAfterMarshal(t, "testdata/type-in-command.json", result)
}

func Test_typeInCommandFailure(t *testing.T) {
	result := NewPageState()

	wrongname := "wrongname"
	cmd := ActionCommand{Command: "mkdir abc", TerminalName: wrongname}
	if err := result.typeInTerminalCommand(&cmd); err == nil {
		t.Errorf("error expected as terminal with name = %s not found", wrongname)
	}

	// expected page state unchanged from initial page
	compareAfterMarshal(t, "testdata/new-page.json", result)
}

func Test_runTerminalCommandSuccess(t *testing.T) {
	result := NewPageState()

	cmd := ActionCommand{Command: "sleep 1", TerminalName: "default"}
	if err := result.typeInTerminalCommand(&cmd); err != nil {
		t.Error(err)
		return
	}
	if err := result.runTerminalCommand(&cmd); err != nil {
		t.Error(err)
		return
	}

	compareAfterMarshal(t, "testdata/run-terminal-command.json", result)
}

func Test_runTerminalCommandSuccess2(t *testing.T) {
	cmd := ActionCommand{
		Command:      "mkdir abc",
		TerminalName: "default",
		UpdateSourceCode: UpdateSourceCode{
			AddDirectories: []AddDirectory{
				{FilePathString: "abc"},
			},
		},
	}

	result := NewPageState()
	if err := result.typeInTerminalCommand(&cmd); err != nil {
		t.Error(err)
		return
	}
	if err := result.runTerminalCommand(&cmd); err != nil {
		t.Error(err)
		return
	}

	compareAfterMarshal(t, "testdata/run-terminal-command2.json", result)
}

func Test_runTerminalCommandSuccess3(t *testing.T) {
	cmd := ActionCommand{
		Command:      "echo hello",
		TerminalName: "default",
		UpdateTerminal: UpdateTerminal{
			Output: "hello",
		},
	}

	result := NewPageState()
	err := result.typeInTerminalCommand(&cmd)
	if err != nil {
		t.Error(err)
		return
	}
	err = result.runTerminalCommand(&cmd)
	if err != nil {
		t.Error(err)
		return
	}

	compareAfterMarshal(t, "testdata/run-terminal-command3.json", result)
}

func Test_ChangeDirectory(t *testing.T) {
	terminal := newTerminal("default")
	cd := UpdateTerminal{
		CurrentDirectoryPath: "hello",
	}

	if terminal.CurrentDirectoryPath != nil && *terminal.CurrentDirectoryPath == cd.CurrentDirectoryPath {
		t.Errorf("terminal's current directory is already same as `cd` target = %s", *terminal.CurrentDirectoryPath)
	}

	terminal.changeCurrentDirectory(cd)
	if *terminal.CurrentDirectoryPath != cd.CurrentDirectoryPath {
		t.Errorf("terminal's current directory is not changed from %s to %s", *terminal.CurrentDirectoryPath, cd.CurrentDirectoryPath)
	}
}

func Test_calcNextStep(t *testing.T) {
	cases := []struct {
		Description string
		CurrentStep string
		Expected    string
	}{
		{Description: "next to 000 is 001", CurrentStep: "000", Expected: "001"},
		{Description: "next to 001 is 002", CurrentStep: "001", Expected: "002"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			result, err := calcNextStep(test.CurrentStep)
			if err != nil {
				t.Errorf("failed to calc next step for %s, %s", result, err)
			}

			if result != test.Expected {
				t.Errorf("expected %s, but result %s", test.Expected, result)
			}
		})
	}
}
