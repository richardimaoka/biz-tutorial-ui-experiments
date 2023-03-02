package model2

import "testing"

func TestTerminal(t *testing.T) {
	terminal := NewTerminal("default")
	compareAfterMarshal(t, "testdata/terminal/new-terminal.json", terminal)
}

func TestTerminalTypein1(t *testing.T) {
	terminal := NewTerminal("default")
	if err := terminal.TypeInCommand("mkdir abc"); err != nil {
		t.Fatalf("no error expected, but %s", err)
	}

	compareAfterMarshal(t, "testdata/terminal/type-in-command1.json", terminal)
}

func TestTerminalTypein2(t *testing.T) {
	terminal := NewTerminal("default")
	if err := terminal.TypeInCommand("mkdir abc"); err != nil {
		t.Fatalf("no error expected, but %s", err)
	}
	if err := terminal.MarkLastCommandExecuted("mkdir abc"); err != nil {
		t.Fatalf("no error expected, but %s", err)
	}
	if err := terminal.TypeInCommand("mkdir cde"); err != nil {
		t.Fatalf("no error expected, but %s", err)
	}

	compareAfterMarshal(t, "testdata/terminal/type-in-command2.json", terminal)
}

func TestTerminalTypeinFail(t *testing.T) {

	terminal := NewTerminal("default")
	if err := terminal.TypeInCommand("mkdir abc"); err != nil {
		t.Fatalf("no error expected, but %s", err)
	}
	if err := terminal.TypeInCommand("mkdir cde"); err == nil {
		t.Fatalf("error expected")
	}

	compareAfterMarshal(t, "testdata/terminal/type-in-command1.json", terminal)
}

func TestTerminalWriteOutput1(t *testing.T) {
	terminal := NewTerminal("default")
	if err := terminal.TypeInCommand("echo abc"); err != nil {
		t.Fatalf("no error expected, but %s", err)
	}
	if err := terminal.MarkLastCommandExecuted("echo abc"); err != nil {
		t.Fatalf("no error expected, but %s", err)
	}
	if err := terminal.WriteOutput("abc"); err != nil {
		t.Fatalf("no error expected, but %s", err)
	}

	compareAfterMarshal(t, "testdata/terminal/write-output.json", terminal)
}
