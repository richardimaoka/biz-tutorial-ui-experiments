package state_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func TestTerminalActions(t *testing.T) {
	cases := []struct {
		goldenFile  string
		elementType state.TerminalElementType
		text        string
	}{
		{"testdata/terminal_action_golden1.json", state.TerminalTypeCommand, "ls"},
		{"testdata/terminal_action_golden2.json", state.TerminalTypeOutput, "aa.txt bb.txt cc.txt"},
		{"testdata/terminal_action_golden3.json", state.TerminalTypeCommand, "cat aa.txt"},
		{"testdata/terminal_action_golden4.json", state.TerminalTypeOutput, "aaaaaaaaaaa"},
	}

	terminal := state.NewTerminal()
	for _, c := range cases {
		switch c.elementType {
		case state.TerminalTypeCommand:
			terminal.WriteCommand(c.text)
		case state.TerminalTypeOutput:
			terminal.WriteOutput(c.text)
		}

		internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, terminal.ToGraphQLTerminal())
	}

}

func TestTerminalMutation1(t *testing.T) {
	s := state.NewTerminal()
	s.WriteCommand("ls")
	s.WriteOutput("aa.txt bb.txt")

	// once GraphQL model is materialized...
	gqlModel := s.ToGraphQLTerminal()
	goldenFile1 := "testdata/terminal_golden1-1.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile1, gqlModel)

	// ... mutation to the state ...
	s.WriteCommand("cat aa.txt")
	s.WriteOutput("aaaaaaaaaa")

	// ... has NO effect on the materialized GraphQL model
	internal.CompareAfterMarshal(t, goldenFile1, gqlModel)

	// ... has effect on a RE-materialized GraphQL model
	goldenFile2 := "testdata/terminal_golden1-2.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile2, s.ToGraphQLTerminal())
}

func TestTerminalMutation2(t *testing.T) {
	s := state.NewTerminal()
	s.WriteCommand("ls")
	s.WriteOutput("aa.txt bb.txt")

	// once GraphQL model is materialized...
	gqlModel := s.ToGraphQLTerminal()
	goldenFile1 := "testdata/terminal_golden2-1.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile1, gqlModel)

	// ... mutation to the materialized model ...
	*gqlModel.Nodes[0].Content.(*model.TerminalCommand).Command = "updated command in terminal"
	*gqlModel.Nodes[1].Content.(*model.TerminalOutput).Output = "updated output in terminal"

	// ... has NO effect on a RE-materialized GraphQL model
	internal.CompareAfterMarshal(t, goldenFile1, s.ToGraphQLTerminal())

	// ... has effect on the materialized GraphQL model
	goldenFile2 := "testdata/terminal_golden2-2.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile2, gqlModel)
}
