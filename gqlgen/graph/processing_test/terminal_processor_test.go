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

func Test_TerminalTransition(t *testing.T) {
	type TestCase struct {
		ExpectedFile   string
		TerminalEffect processing.TerminalEffect
	}

	runTestCases := func(t *testing.T, testCases []TestCase) {
		terminal := processing.NewTerminalProcessor("default")
		for _, c := range testCases {
			step := fmt.Sprintf("%03d", c.TerminalEffect.SeqNo)
			terminal.Transition(step, c.TerminalEffect)

			// if `-update` flag is passed from command line, update the golden file
			if *update {
				internal.WriteGoldenFile(t, c.ExpectedFile, terminal.ToGraphQLTerminal())
			}

			internal.CompareAfterMarshal(t, c.ExpectedFile, terminal.ToGraphQLTerminal())
		}
	}

	t.Run("write_command1", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{"testdata/terminal/transition1-1.json", processing.TerminalEffect{SeqNo: 0, TerminalName: "default", Command: "mkdir abc", Output: nil, CurrentDirectory: nil}},
		})
	})

	t.Run("write_command2", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{"testdata/terminal/transition2-1.json", processing.TerminalEffect{SeqNo: 0, TerminalName: "default", Command: "mkdir abc", Output: nil, CurrentDirectory: nil}},
			{"testdata/terminal/transition2-2.json", processing.TerminalEffect{SeqNo: 0, TerminalName: "default", Command: "mkdir def", Output: nil, CurrentDirectory: nil}},
		})
	})

	t.Run("write_output", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{"testdata/terminal/transition3-1.json", processing.TerminalEffect{SeqNo: 0, TerminalName: "default", Command: "echo abc", Output: internal.Address("abc"), CurrentDirectory: nil}},
		})
	})

	t.Run("change_current_directory1", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{"testdata/terminal/transition4-1.json", processing.TerminalEffect{SeqNo: 0, TerminalName: "default", Command: "cd hello", Output: nil, CurrentDirectory: internal.Address("hello")}},
		})
	})

	t.Run("change_current_directory2", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{"testdata/terminal/transition5-1.json", processing.TerminalEffect{SeqNo: 0, TerminalName: "default", Command: "cd hello/world/thunder", Output: nil, CurrentDirectory: internal.Address("hello/world/thunder")}},
		})
	})
}

// Test mutation after terminal.ToGraphQLModel()
// Once a GraphQL model is materialized with terminal.ToGraphQLModel(), mutation to the terminal should have no effect on the GraphQL model
func Test_TerminalMutation1(t *testing.T) {
	terminal := processing.NewTerminalProcessor("default")
	terminal.Transition("001", processing.TerminalEffect{Command: "echo abc", Output: internal.Address("abc")})
	terminal.Transition("002", processing.TerminalEffect{Command: "cd hello/world/thunder", CurrentDirectory: internal.Address("hello/world/thunder")})

	// once GraphQL model is materialized...
	materialized := terminal.ToGraphQLTerminal()
	if *update { // if `-update` flag is passed from command line, update the golden file
		internal.WriteGoldenFile(t, "testdata/terminal/mutation1-1.json", materialized)
	}
	internal.CompareAfterMarshal(t, "testdata/terminal/mutation1-1.json", materialized)

	// ...mutation to Terminal...
	terminal.Transition("003", processing.TerminalEffect{Command: "mutation extra command", Output: internal.Address("mutation extra output"), CurrentDirectory: internal.Address("mutated/current/dir")})

	// ...should of course have effect on re-materialized GraphQL model
	if *update { // if `-update` flag is passed from command line, update the golden file
		internal.WriteGoldenFile(t, "testdata/terminal/mutation1-2.json", terminal.ToGraphQLTerminal())
	}
	internal.CompareAfterMarshal(t, "testdata/terminal/mutation1-2.json", terminal.ToGraphQLTerminal())

	// ...but should have no effect on materialized GraphQL model
	internal.CompareAfterMarshal(t, "testdata/terminal/mutation1-1.json", materialized)
}

// Test mutation after terminal.ToGraphQLModel()
// Once a GraphQL model is materialized with terminal.ToGraphQLModel(), mutation to the GraphQL model should have no effect on the terminal
func TestTerminal_Mutation2(t *testing.T) {
	terminal := processing.NewTerminalProcessor("default")
	terminal.Transition("001", processing.TerminalEffect{Command: "echo abc", Output: internal.Address("abc")})
	terminal.Transition("002", processing.TerminalEffect{Command: "cd hello/world/thunder", CurrentDirectory: internal.Address("hello/world/thunder")})

	// once GraphQL model is materialized...
	materialized := terminal.ToGraphQLTerminal()
	if *update { // if `-update` flag is passed from command line, update the golden file
		internal.WriteGoldenFile(t, "testdata/terminal/mutation2-1.json", materialized)
	}
	internal.CompareAfterMarshal(t, "testdata/terminal/mutation2-1.json", materialized)

	// ...mutation to materialized GraphQL model...
	ptr1 := materialized.Nodes[0].Content.(*model.TerminalCommand).Command
	*ptr1 = "mutation extra command"
	ptr2 := materialized.Nodes[1].Content.(*model.TerminalOutput).Output
	*ptr2 = "mutation extra output"

	// ...should of course have effect on materialized GraphQL model
	if *update { // if `-update` flag is passed from command line, update the golden file
		internal.WriteGoldenFile(t, "testdata/terminal/mutation2-2.json", materialized)
	}
	internal.CompareAfterMarshal(t, "testdata/terminal/mutation2-2.json", materialized)

	// ...but should have no effect on Terminal
	internal.CompareAfterMarshal(t, "testdata/terminal/mutation2-1.json", terminal.ToGraphQLTerminal())
}

// test terminalProcessor.Clone() and terminalElementProcessor.Clone() as the former calls latter
func TestTerminal_Clone(t *testing.T) {
	terminal := processing.NewTerminalProcessor("default")
	terminal.Transition("001", processing.TerminalEffect{Command: "echo abc", Output: internal.Address("abc")})

	// once cloned...
	terminalCloned := terminal.Clone()
	if *update { // if `-update` flag is passed from command line, update the golden file
		internal.WriteGoldenFile(t, "testdata/terminal/clone1-1.json", terminal.ToGraphQLTerminal())
	}
	internal.CompareAfterMarshal(t, "testdata/terminal/clone1-1.json", terminalCloned.ToGraphQLTerminal())

	// ...mutation to terminal...
	terminal.Transition("002", processing.TerminalEffect{Command: "echo def", Output: internal.Address("def")})
	terminal.Transition("003", processing.TerminalEffect{Command: "cd hello/world/thunder", CurrentDirectory: internal.Address("hello/world/thunder")})

	// ...should of course have effect on terminal itself
	if *update { // if `-update` flag is passed from command line, update the golden file
		internal.WriteGoldenFile(t, "testdata/terminal/clone1-2.json", terminal.ToGraphQLTerminal())
	}
	internal.CompareAfterMarshal(t, "testdata/terminal/clone1-2.json", terminal.ToGraphQLTerminal())

	// ...but should have no effect on cloned terminal
	internal.CompareAfterMarshal(t, "testdata/terminal/clone1-1.json", terminalCloned.ToGraphQLTerminal())
}
