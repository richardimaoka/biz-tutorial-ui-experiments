package processing

import "testing"

func TestTerminal_MutationCommand(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	terminal.WriteCommand("mkdir abc")
	terminal.WriteCommand("mkdir def")
	result := terminal.ToTerminal()

	// after terminal is materialized to GraphQL object, mutation should have no effect
	terminal.elements[0].(*terminalCommand).promptExpression = "mutated-command"
	terminal.elements[0].(*terminalCommand).promptSymbol = 'G'
	terminal.elements[0].(*terminalCommand).command = "mkdir ghi"

	compareAfterMarshal(t, "testdata/terminal/write-command2.json", result)
}

func TestTerminalElement_MutationOutput(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	terminal.WriteCommand("echo abc")
	terminal.WriteOutput("abc")

	result := terminal.ToTerminal()

	// after terminal is materialized to GraphQL object, mutation should have no effect
	terminal.elements[1].(*terminalOutput).output = "def"

	compareAfterMarshal(t, "testdata/terminal/write-output1.json", result)
}
