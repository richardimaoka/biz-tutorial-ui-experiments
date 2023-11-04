package state_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

// state.File is effectively immutable, so no need to test mutation to the state
/*
func TestFileNodeMutation1(t *testing.T) {
	s, err := fileStateFromCommit(
		"https://github.com/richardimaoka/next-sandbox.git",
		"8adac375628219e020d4b5957ff24f45954cbd3f", // commit = 'npx create-next-app@latest'
		"next/package.json",
	)
	if err != nil {
		t.Fatalf("failed in TestFileNodeMutation1, %s", err)
	}

	// once GraphQL model is materialized...
	gqlModel := s.ToGraphQLFileNode()
	goldenFile1 := "testdata/file_node_mutation_golden1-1.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile1, gqlModel)

	// ... mutation to the state ...
  // state.File is effectively immutable, so no need to test mutation to the state

	// ... has NO effect on the materialized GraphQL model
	internal.CompareAfterMarshal(t, goldenFile1, gqlModel)

	// ... has effect on a RE-materialized GraphQL model
	goldenFile2 := "testdata/file_node_mutation_golden1-2.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile2, s.ToGraphQLFileNode())

	// and golden files are indeed different
	internal.FilesMustUnmatch(t, goldenFile1, goldenFile2)

}
*/

func TestFileNodeMutation2(t *testing.T) {
	s, err := fileStateFromCommit(
		"https://github.com/richardimaoka/next-sandbox.git",
		"8adac375628219e020d4b5957ff24f45954cbd3f", // commit = 'npx create-next-app@latest'
		"next/package.json",
	)
	if err != nil {
		t.Fatalf("failed in TestFileNodeMutation2, %s", err)
	}

	// once GraphQL model is materialized...
	gqlModel := s.ToGraphQLFileNode()
	goldenFile1 := "testdata/file_node_mutation_golden2-1.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile1, gqlModel)

	// ... mutation to the materialized model ...
	*gqlModel.Name = "package-mudated.json"
	gqlModel.FilePath = "next/package-mudated.json"
	*gqlModel.Offset = 5
	*gqlModel.IsUpdated = !*gqlModel.IsUpdated
	gqlModel.NodeType = model.FileNodeTypeDirectory

	// ... has NO effect on a RE-materialized GraphQL model
	internal.CompareAfterMarshal(t, goldenFile1, s.ToGraphQLFileNode())

	// ... has effect on the materialized GraphQL model
	goldenFile2 := "testdata/file_node_mutation_golden2-2.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile2, gqlModel)

	// and golden files are indeed different
	internal.FilesMustUnmatch(t, goldenFile1, goldenFile2)
}
