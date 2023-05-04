package processing

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/internal"
)

// test case for TerminalProcessor's WriteCommand method
func TestTerminal(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	result := terminal.ToGraphQLTerminal()
	internal.CompareAfterMarshal(t, "testdata/terminal/new-terminal.json", result)
}

func TestTerminal_WriteCommand1(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	terminal.WriteCommand("mkdir abc")
	result := terminal.ToGraphQLTerminal()
	internal.CompareAfterMarshal(t, "testdata/terminal/write-command1.json", result)
}

func TestTerminal_WriteCommand2(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	terminal.WriteCommand("mkdir abc")
	terminal.WriteCommand("mkdir def")
	result := terminal.ToGraphQLTerminal()
	internal.CompareAfterMarshal(t, "testdata/terminal/write-command2.json", result)
}

func TestTerminal_WriteOutput(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	terminal.WriteCommand("echo abc")
	terminal.WriteOutput("abc")
	result := terminal.ToGraphQLTerminal()
	internal.CompareAfterMarshal(t, "testdata/terminal/write-output1.json", result)
}

func TestTerminal_ChangeCurrentDirectory1(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	terminal.WriteCommand("cd hello")
	terminal.ChangeCurrentDirectory("hello")
	result := terminal.ToGraphQLTerminal()
	internal.CompareAfterMarshal(t, "testdata/terminal/cd1.json", result)
}

func TestTerminal_ChangeCurrentDirectory2(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	terminal.WriteCommand("cd hello/world/thunder")
	terminal.ChangeCurrentDirectory("hello/world/thunder")
	result := terminal.ToGraphQLTerminal()
	internal.CompareAfterMarshal(t, "testdata/terminal/cd2.json", result)
}

func TestTerminal_Mutation(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	terminal.WriteCommand("echo abc")
	terminal.WriteOutput("abc")
	terminal.WriteCommand("cd hello/world/thunder")
	terminal.ChangeCurrentDirectory("hello/world/thunder")

	// once materialized GraphQL model...
	materialized := terminal.ToGraphQLTerminal()
	internal.CompareAfterMarshal(t, "testdata/terminal/materialized.json", materialized)

	// ...mutation afterwards should have no effect
	terminal.ChangeCurrentDirectory("mutated/current/dir")
	terminal.WriteCommand("mutation extra command")
	terminal.WriteOutput("mutation extra output")

	// materialized GraphQL model should not be affected
	internal.CompareAfterMarshal(t,
		"testdata/terminal/materialized.json",
		materialized) // by changing this to terminal.ToGraphQLTerminal(), you can confirm mutation actually updated terminal
}

// test terminalProcessor.Clone() and terminalElementProcessor.Clone() as the former calls latter
func TestTerminal_Clone(t *testing.T) {
	terminal := NewTerminalProcessor("default")
	terminal.WriteCommand("echo abc")
	terminal.WriteOutput("abc")

	// once cloned...
	terminalOriginal := terminal.Clone()
	internal.CompareAfterMarshal(t, "testdata/terminal/cloned.json", terminalOriginal.ToGraphQLTerminal())

	// ...mutation afterwards should have no effect
	terminal.WriteCommand("echo def")
	terminal.WriteOutput("def")
	terminal.WriteCommand("cd hello/world/thunder")
	terminal.ChangeCurrentDirectory("hello/world/thunder")

	// cloned terminal is not affected
	internal.CompareAfterMarshal(t,
		"testdata/terminal/cloned.json",
		terminalOriginal.ToGraphQLTerminal(), // by changing this to terminal.ToGraphQLTerminal(), you can confirm mutation actually updated terminal
	)
}
