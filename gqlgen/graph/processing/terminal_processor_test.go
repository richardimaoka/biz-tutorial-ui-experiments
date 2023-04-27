package processing

import (
	"testing"
)

// test case for TerminalProcessor's WriteCommand method
func TestTerminal(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	result := terminal.ToGraphQLTerminal()
	compareAfterMarshal(t, "testdata/terminal/new-terminal.json", result)
}

func TestTerminal_WriteCommand1(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	terminal.WriteCommand("mkdir abc")
	result := terminal.ToGraphQLTerminal()
	compareAfterMarshal(t, "testdata/terminal/write-command1.json", result)
}

func TestTerminal_WriteCommand2(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	terminal.WriteCommand("mkdir abc")
	terminal.WriteCommand("mkdir def")
	result := terminal.ToGraphQLTerminal()
	compareAfterMarshal(t, "testdata/terminal/write-command2.json", result)
}

func TestTerminal_WriteOutput(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	terminal.WriteCommand("echo abc")
	terminal.WriteOutput("abc")
	result := terminal.ToGraphQLTerminal()
	compareAfterMarshal(t, "testdata/terminal/write-output1.json", result)
}

func TestTerminal_ChangeCurrentDirectory1(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	terminal.WriteCommand("cd hello")
	terminal.ChangeCurrentDirectory("hello")
	result := terminal.ToGraphQLTerminal()
	compareAfterMarshal(t, "testdata/terminal/cd1.json", result)
}

func TestTerminal_ChangeCurrentDirectory2(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	terminal.WriteCommand("cd hello/world/thunder")
	terminal.ChangeCurrentDirectory("hello/world/thunder")
	result := terminal.ToGraphQLTerminal()
	compareAfterMarshal(t, "testdata/terminal/cd2.json", result)
}

func TestTerminal_Clone1(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	terminal.WriteCommand("echo abc")
	terminal.WriteOutput("abc")

	terminalOriginal := terminal.Clone()

	terminal.WriteCommand("echo def")
	terminal.WriteOutput("def")
	terminal.WriteCommand("cd hello/world/thunder")
	terminal.ChangeCurrentDirectory("hello/world/thunder")

	result := terminalOriginal.ToGraphQLTerminal()
	compareAfterMarshal(t, "testdata/terminal/clone.json", result)
}

func TestTerminal_Clone2(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	terminal.WriteCommand("echo abc")
	terminal.WriteOutput("abc")

	terminalOriginal := terminal.Clone()

	terminal.currentDirectory = "mutated/current/dir"
	terminal.terminalName = "mutated terminal name"
	terminal.elements[0].(*terminalCommandProcessor).promptExpression = "mutated-expression"
	terminal.elements[0].(*terminalCommandProcessor).promptSymbol = 'X'
	terminal.elements[0].(*terminalCommandProcessor).command = "mutated-command"
	terminal.elements[1].(*terminalOutputProcessor).output = "mutated-output"

	result := terminalOriginal.ToGraphQLTerminal()
	compareAfterMarshal(t, "testdata/terminal/clone.json", result)
}

func TestTerminal_Mutation(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	terminal.WriteCommand("echo abc")
	terminal.WriteOutput("abc")
	terminal.WriteCommand("cd hello/world/thunder")
	terminal.ChangeCurrentDirectory("hello/world/thunder")

	result := terminal.ToGraphQLTerminal()

	terminal.currentDirectory = "mutated/current/dir"
	terminal.terminalName = "mutated terminal name"
	terminal.elements[0].(*terminalCommandProcessor).promptExpression = "mutated-expression"
	terminal.elements[0].(*terminalCommandProcessor).promptSymbol = 'X'
	terminal.elements[0].(*terminalCommandProcessor).command = "mutated-command"
	terminal.elements[1].(*terminalOutputProcessor).output = "mutated-output"

	compareAfterMarshal(t, "testdata/terminal/mutated.json", result)
}
