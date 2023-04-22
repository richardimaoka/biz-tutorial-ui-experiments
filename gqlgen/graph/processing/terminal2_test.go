package processing

import (
	"testing"
)

/*
   write command
	 write command x 2
	 write command, then output
	 change current directory
	 clone, mutate and check if origianl not mutaed



*/

// test case for Terminal2's WriteCommand method
func TestTerminal2(t *testing.T) {
	terminal := NewTerminal2("default")
	result := terminal.ToTerminal()
	compareAfterMarshal(t, "testdata/terminal2/new-terminal.json", result)
}

func TestTerminal2_WriteCommand(t *testing.T) {
	terminal := NewTerminal2("default")
	terminal.WriteCommand("mkdir abc")
	result := terminal.ToTerminal()
	compareAfterMarshal(t, "testdata/terminal2/write-command1.json", result)

	terminal.WriteCommand("mkdir def")
	result = terminal.ToTerminal()
	compareAfterMarshal(t, "testdata/terminal2/write-command2.json", result)
}

func TestTerminal2_WriteOutput(t *testing.T) {
	terminal := NewTerminal2("default")
	terminal.WriteCommand("echo abc")
	terminal.WriteOutput("abc")
	result := terminal.ToTerminal()
	compareAfterMarshal(t, "testdata/terminal2/write-output1.json", result)
}

func TestTerminal2_ChangeCurrentDirectory1(t *testing.T) {
	terminal := NewTerminal2("default")
	terminal.WriteCommand("cd hello")
	terminal.ChangeCurrentDirectory("hello")
	result := terminal.ToTerminal()
	compareAfterMarshal(t, "testdata/terminal2/cd1.json", result)
}

func TestTerminal2_ChangeCurrentDirectory2(t *testing.T) {
	terminal := NewTerminal2("default")
	terminal.WriteCommand("cd hello/world/thunder")
	terminal.ChangeCurrentDirectory("hello/world/thunder")
	result := terminal.ToTerminal()
	compareAfterMarshal(t, "testdata/terminal2/cd2.json", result)
}
