package processing

import (
	"testing"
)

// test case for TerminalProcessor's WriteCommand method
func TestTerminal2(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	result := terminal.ToTerminal()
	compareAfterMarshal(t, "testdata/terminal2/new-terminal.json", result)
}

func TestTerminal2_WriteCommand(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	terminal.WriteCommand("mkdir abc")
	result := terminal.ToTerminal()
	compareAfterMarshal(t, "testdata/terminal2/write-command1.json", result)

	terminal.WriteCommand("mkdir def")
	result = terminal.ToTerminal()
	compareAfterMarshal(t, "testdata/terminal2/write-command2.json", result)
}

func TestTerminal2_WriteOutput(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	terminal.WriteCommand("echo abc")
	terminal.WriteOutput("abc")
	result := terminal.ToTerminal()
	compareAfterMarshal(t, "testdata/terminal2/write-output1.json", result)
}

func TestTerminal2_ChangeCurrentDirectory1(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	terminal.WriteCommand("cd hello")
	terminal.ChangeCurrentDirectory("hello")
	result := terminal.ToTerminal()
	compareAfterMarshal(t, "testdata/terminal2/cd1.json", result)
}

func TestTerminal2_ChangeCurrentDirectory2(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	terminal.WriteCommand("cd hello/world/thunder")
	terminal.ChangeCurrentDirectory("hello/world/thunder")
	result := terminal.ToTerminal()
	compareAfterMarshal(t, "testdata/terminal2/cd2.json", result)
}

func TestTerminal2_Clone(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	terminal.WriteCommand("echo abc")
	terminal.WriteOutput("abc")

	terminalOriginal := terminal.Clone()

	terminal.WriteCommand("echo def")
	terminal.WriteOutput("def")
	terminal.WriteCommand("cd hello/world/thunder")
	terminal.ChangeCurrentDirectory("hello/world/thunder")
	result := terminalOriginal.ToTerminal()
	compareAfterMarshal(t, "testdata/terminal2/clone.json", result)
}
