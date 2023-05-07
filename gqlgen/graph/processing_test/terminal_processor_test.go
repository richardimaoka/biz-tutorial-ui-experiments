package processing_test

import (
	"fmt"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/processing"
)

// TODO: rewrite tests with table driven tests
// TODO: private-ise the methods exept Transition
// TODO: move the test to processing_test
//
// type TerminalEffect struct {
// 	SeqNo            int     `json:"seqNo"`
// 	TerminalName     string  `json:"terminalName"`
// 	Command          string  `json:"command"`
// 	Output           *string `json:"output"`           //if zero value, no output after execution
// 	CurrentDirectory *string `json:"currentDirectory"` //if zero value, current directory is not changed after execution
// }
// var cases := []TerminalEffect{
// 	{ SeqNo: 0, TerminalName: "default", Command: "echo abc", Output: nil, CurrentDirectory: nil },
// }

// test case for TerminalProcessor's WriteCommand method

func Test_Terminal(t *testing.T) {
	type TestCase struct {
		ExpectedFile   string
		TerminalEffect processing.TerminalEffect
	}

	testCases := []TestCase{
		{"testdata/terminal/terminal1-1.json", processing.TerminalEffect{SeqNo: 0, TerminalName: "default", Command: "mkdir abc", Output: nil, CurrentDirectory: nil}},
		// {"testdata/terminal/terminal1-2.json", processing.TerminalEffect{SeqNo: 1, TerminalName: "default", Command: "mkdir abc", Output: nil, CurrentDirectory: internal.Address("def")}},
		// {"testdata/terminal/terminal1-3.json", processing.TerminalEffect{SeqNo: 2, TerminalName: "default", Command: "echo abc", Output: internal.Address("abc"), CurrentDirectory: nil}},
		// {"testdata/terminal/terminal1-4.json", processing.TerminalEffect{SeqNo: 3, TerminalName: "default", Command: "mkdir def", Output: nil, CurrentDirectory: nil}},
		// {"testdata/terminal/terminal1-5.json", processing.TerminalEffect{SeqNo: 4, TerminalName: "default", Command: "mkdir abc", Output: nil, CurrentDirectory: nil}},
	}

	terminal := processing.NewTerminalProcessor("default")
	for _, c := range testCases {
		t.Run("tt.descriptio", func(t *testing.T) {
			step := fmt.Sprintf("%03d", c.TerminalEffect.SeqNo)
			terminal.Transition(c.TerminalEffect, step)

			if *update {
				internal.WriteGoldenFile(t, c.ExpectedFile, terminal.ToGraphQLTerminal())
			}

			internal.CompareAfterMarshal(t, c.ExpectedFile, terminal.ToGraphQLTerminal())
		})
	}
}

func TestTerminal(t *testing.T) {
	terminal := processing.NewTerminalProcessor("default")

	result := terminal.ToGraphQLTerminal()
	internal.CompareAfterMarshal(t, "testdata/terminal/new-terminal.json", result)
}

func TestTerminal_WriteCommand1(t *testing.T) {
	terminal := processing.NewTerminalProcessor("default")
	terminal.WriteCommand("mkdir abc")
	result := terminal.ToGraphQLTerminal()
	internal.CompareAfterMarshal(t, "testdata/terminal/write-command1.json", result)
}

func TestTerminal_WriteCommand2(t *testing.T) {
	terminal := processing.NewTerminalProcessor("default")
	terminal.WriteCommand("mkdir abc")
	terminal.WriteCommand("mkdir def")
	result := terminal.ToGraphQLTerminal()
	internal.CompareAfterMarshal(t, "testdata/terminal/write-command2.json", result)
}

func TestTerminal_WriteOutput(t *testing.T) {
	terminal := processing.NewTerminalProcessor("default")
	terminal.WriteCommand("echo abc")
	terminal.WriteOutput("abc")
	result := terminal.ToGraphQLTerminal()
	internal.CompareAfterMarshal(t, "testdata/terminal/write-output1.json", result)
}

func TestTerminal_ChangeCurrentDirectory1(t *testing.T) {
	terminal := processing.NewTerminalProcessor("default")
	terminal.WriteCommand("cd hello")
	terminal.ChangeCurrentDirectory("hello")
	result := terminal.ToGraphQLTerminal()
	internal.CompareAfterMarshal(t, "testdata/terminal/cd1.json", result)
}

func TestTerminal_ChangeCurrentDirectory2(t *testing.T) {
	terminal := processing.NewTerminalProcessor("default")
	terminal.WriteCommand("cd hello/world/thunder")
	terminal.ChangeCurrentDirectory("hello/world/thunder")
	result := terminal.ToGraphQLTerminal()
	internal.CompareAfterMarshal(t, "testdata/terminal/cd2.json", result)
}

func TestTerminal_Mutation1(t *testing.T) {
	terminal := processing.NewTerminalProcessor("default")
	terminal.WriteCommand("echo abc")
	terminal.WriteOutput("abc")
	terminal.WriteCommand("cd hello/world/thunder")
	terminal.ChangeCurrentDirectory("hello/world/thunder")

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
