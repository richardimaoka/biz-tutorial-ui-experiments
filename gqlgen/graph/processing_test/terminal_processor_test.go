package processing_test

import (
	"fmt"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/processing"
)

func Test_NewTerminalProcessor(t *testing.T) {
	terminal := processing.NewTerminalProcessor("default")
	result := terminal.ToGraphQLTerminal()
	internal.CompareAfterMarshal(t, "testdata/terminal/new-terminal.json", result)
}

func Test_TerminalEffect(t *testing.T) {
	type TestCase struct {
		ExpectedFile   string
		TerminalEffect processing.TerminalEffect
	}

	runTestCases := func(t *testing.T, testCases []TestCase) {
		terminal := processing.NewTerminalProcessor("default")
		for _, c := range testCases {
			step := fmt.Sprintf("%03d", c.TerminalEffect.SeqNo)
			terminal.Transition(c.TerminalEffect, step)

			// if `--update` flag is passed from command line, update the golden file
			if *update {
				internal.WriteGoldenFile(t, c.ExpectedFile, terminal.ToGraphQLTerminal())
			}

			internal.CompareAfterMarshal(t, c.ExpectedFile, terminal.ToGraphQLTerminal())
		}
	}

	t.Run("write_command1", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{"testdata/terminal/terminal1-1.json", processing.TerminalEffect{SeqNo: 0, TerminalName: "default", Command: "mkdir abc", Output: nil, CurrentDirectory: nil}},
		})
	})

	t.Run("write_command2", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{"testdata/terminal/terminal2-1.json", processing.TerminalEffect{SeqNo: 0, TerminalName: "default", Command: "mkdir abc", Output: nil, CurrentDirectory: nil}},
			{"testdata/terminal/terminal2-2.json", processing.TerminalEffect{SeqNo: 0, TerminalName: "default", Command: "mkdir def", Output: nil, CurrentDirectory: nil}},
		})
	})

	t.Run("write_output", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{"testdata/terminal/terminal3-1.json", processing.TerminalEffect{SeqNo: 0, TerminalName: "default", Command: "echo abc", Output: internal.Address("abc"), CurrentDirectory: nil}},
		})
	})

	t.Run("change_current_directory1", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{"testdata/terminal/terminal4-1.json", processing.TerminalEffect{SeqNo: 0, TerminalName: "default", Command: "cd hello", Output: nil, CurrentDirectory: internal.Address("hello")}},
		})
	})

	t.Run("change_current_directory2", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{"testdata/terminal/terminal5-1.json", processing.TerminalEffect{SeqNo: 0, TerminalName: "default", Command: "cd hello/world/thunder", Output: nil, CurrentDirectory: internal.Address("hello/world/thunder")}},
		})
	})
}

func TestTerminal_Mutation1(t *testing.T) {
	terminal := processing.NewTerminalProcessor("default")
	terminal.Transition(processing.TerminalEffect{Command: "echo abc", Output: internal.Address("abc")}, "001")
	terminal.Transition(processing.TerminalEffect{Command: "cd hello/world/thunder", CurrentDirectory: internal.Address("hello/world/thunder")}, "002")

	// once materialized GraphQL model...
	materialized := terminal.ToGraphQLTerminal()
	internal.CompareAfterMarshal(t, "testdata/terminal/materialized.json", materialized)

	// ...mutation to Terminal should have no effect ...
	terminal.ChangeCurrentDirectory("mutated/current/dir")
	terminal.WriteCommand("mutation extra command")
	terminal.WriteOutput("mutation extra output")

	// ...on materialized GraphQL model
	internal.CompareAfterMarshal(t,
		"testdata/terminal/materialized.json",
		materialized)
}

func TestTerminal_Mutation2(t *testing.T) {
	terminal := processing.NewTerminalProcessor("default")
	terminal.WriteCommand("echo abc")
	terminal.WriteOutput("abc")
	terminal.WriteCommand("cd hello/world/thunder")
	terminal.ChangeCurrentDirectory("hello/world/thunder")

	// once materialized GraphQL model...
	materialized := terminal.ToGraphQLTerminal()
	internal.CompareAfterMarshal(t, "testdata/terminal/materialized.json", materialized)

	// ...mutation to materialized GraphQL model should have no effect...
	ptr1 := materialized.Nodes[0].Content.(*model.TerminalCommand).Command
	*ptr1 = "mutation extra command"
	ptr2 := materialized.Nodes[1].Content.(*model.TerminalOutput).Output
	*ptr2 = "mutation extra output"

	// ...on Terminal
	internal.CompareAfterMarshal(t,
		"testdata/terminal/materialized.json",
		terminal.ToGraphQLTerminal())
}

// test terminalProcessor.Clone() and terminalElementProcessor.Clone() as the former calls latter
func TestTerminal_Clone(t *testing.T) {
	terminal := processing.NewTerminalProcessor("default")
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
