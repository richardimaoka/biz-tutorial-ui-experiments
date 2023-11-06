package processing_test

import (
	"fmt"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/testio"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/preprocess/processing"
)

func Test_NewTerminalProcessor(t *testing.T) {
	terminal := processing.NewTerminalProcessor("default")
	result := terminal.ToGraphQLTerminal()
	testio.CompareAfterMarshal(t, "testdata/terminal/new-terminal.json", result)
}

func Test_TerminalTransition(t *testing.T) {
	type TestCase struct {
		ExpectedFile string
		Operation    processing.TerminalOperation
	}

	runTestCases := func(t *testing.T, testCases []TestCase) {
		terminal := processing.NewTerminalProcessor("default")
		for i, c := range testCases {
			step := fmt.Sprintf("%03d", i)
			terminal.Transition(step, c.Operation)

			testio.CompareWithGoldenFile(t, *updateFlag, c.ExpectedFile, terminal.ToGraphQLTerminal())
		}
	}

	t.Run("write_command1", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{"testdata/terminal/transition1-1.json", processing.TerminalCommand{TerminalName: "default", Command: "mkdir abc"}},
		})
	})

	t.Run("write_command2", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{"testdata/terminal/transition2-1.json", processing.TerminalCommand{TerminalName: "default", Command: "mkdir abc"}},
			{"testdata/terminal/transition2-2.json", processing.TerminalCommand{TerminalName: "default", Command: "mkdir def"}},
		})
	})

	t.Run("write_output", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{"testdata/terminal/transition3-1.json", processing.TerminalCommandWithOutput{TerminalName: "default", Command: "echo abc", Output: "abc"}},
		})
	})

	t.Run("change_current_directory1", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{"testdata/terminal/transition4-1.json", processing.TerminalCommandWithCd{TerminalName: "default", Command: "cd hello", CurrentDirectory: "hello"}},
		})
	})

	t.Run("change_current_directory2", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{"testdata/terminal/transition5-1.json", processing.TerminalCommandWithCd{TerminalName: "default", Command: "cd hello/world/thunder", CurrentDirectory: "hello/world/thunder"}},
		})
	})
}

// Test mutation after terminal.ToGraphQLModel()
// Once a GraphQL model is materialized with terminal.ToGraphQLModel(), mutation to the terminal should have no effect on the GraphQL model
func Test_TerminalMutation1(t *testing.T) {
	terminal := processing.NewTerminalProcessor("default")
	terminal.Transition("001", processing.TerminalCommandWithOutput{Command: "echo abc", Output: "abc"})
	terminal.Transition("002", processing.TerminalCommandWithCd{Command: "cd hello/world/thunder", CurrentDirectory: "hello/world/thunder"})

	// once GraphQL model is materialized...
	materialized := terminal.ToGraphQLTerminal()
	testio.CompareWithGoldenFile(t, *updateFlag, "testdata/terminal/mutation1-1.json", materialized)

	// ...mutation to Terminal...
	terminal.Transition("003", processing.TerminalCommandWithOutputCd{Command: "mutation extra command", Output: "mutation extra output", CurrentDirectory: "mutated/current/dir"})

	// ...should of course have effect on re-materialized GraphQL model
	testio.CompareWithGoldenFile(t, *updateFlag, "testdata/terminal/mutation1-2.json", terminal.ToGraphQLTerminal())

	// ...but should have no effect on materialized GraphQL model
	testio.CompareAfterMarshal(t, "testdata/terminal/mutation1-1.json", materialized)
}

// Test mutation after terminal.ToGraphQLModel()
// Once a GraphQL model is materialized with terminal.ToGraphQLModel(), mutation to the GraphQL model should have no effect on the terminal
func TestTerminal_Mutation2(t *testing.T) {
	terminal := processing.NewTerminalProcessor("default")
	terminal.Transition("001", processing.TerminalCommandWithOutput{Command: "echo abc", Output: "abc"})
	terminal.Transition("002", processing.TerminalCommandWithCd{Command: "cd hello/world/thunder", CurrentDirectory: "hello/world/thunder"})

	// once GraphQL model is materialized...
	materialized := terminal.ToGraphQLTerminal()

	testio.CompareWithGoldenFile(t, *updateFlag, "testdata/terminal/mutation2-1.json", materialized)

	// ...mutation to materialized GraphQL model...
	ptr1 := materialized.Nodes[0].Content.(*model.TerminalCommand).Command
	*ptr1 = "mutation extra command"
	ptr2 := materialized.Nodes[1].Content.(*model.TerminalOutput).Output
	*ptr2 = "mutation extra output"

	// ...should of course have effect on materialized GraphQL model
	testio.CompareWithGoldenFile(t, *updateFlag, "testdata/terminal/mutation2-2.json", materialized)

	// ...but should have no effect on Terminal
	testio.CompareAfterMarshal(t, "testdata/terminal/mutation2-1.json", terminal.ToGraphQLTerminal())
}

// Clone() method testing is needed as TerminalProcessor is a state**ful** structure
func TestTerminal_Clone(t *testing.T) {
	terminal := processing.NewTerminalProcessor("default")
	terminal.Transition("001", processing.TerminalCommandWithOutput{Command: "echo abc", Output: "abc"})

	// once cloned...
	terminalCloned := terminal.Clone()
	testio.CompareWithGoldenFile(t, *updateFlag, "testdata/terminal/clone1-1.json", terminalCloned.ToGraphQLTerminal())

	// ...mutation to terminal...
	terminal.Transition("002", processing.TerminalCommandWithOutput{Command: "echo def", Output: "def"})
	terminal.Transition("003", processing.TerminalCommandWithCd{Command: "cd hello/world/thunder", CurrentDirectory: "hello/world/thunder"})

	// ...should of course have effect on terminal itself
	testio.CompareWithGoldenFile(t, *updateFlag, "testdata/terminal/clone1-2.json", terminal.ToGraphQLTerminal())

	// ...but should have no effect on cloned terminal
	testio.CompareAfterMarshal(t, "testdata/terminal/clone1-1.json", terminalCloned.ToGraphQLTerminal())
}
