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
}
