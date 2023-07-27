package state_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

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
