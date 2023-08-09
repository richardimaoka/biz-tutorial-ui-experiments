package state_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

func TestOpenFileMutation1(t *testing.T) {
	s, err := stateFileFromCommit(
		"https://github.com/richardimaoka/next-sandbox.git",
		"8adac375628219e020d4b5957ff24f45954cbd3f", // commit = 'npx create-next-app@latest'
		"next/package.json",
	)
	if err != nil {
		t.Fatalf("failed in TestOpenFileMutation1, %s", err)
	}

	// once GraphQL model is materialized...
	gqlModel := s.ToGraphQLFileNode()
	goldenFile1 := "testdata/openfile_mutation_golden1-1.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile1, gqlModel)

	// ... mutation to the state ...
	s.ToFileAdded2()

	// ... has NO effect on the materialized GraphQL model
	internal.CompareAfterMarshal(t, goldenFile1, gqlModel)

	// ... has effect on a RE-materialized GraphQL model
	goldenFile2 := "testdata/openfile_mutation_golden1-2.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile2, s.ToGraphQLOpenFile())

	// and golden files are indeed different
	internal.FilesMustUnmatch(t, goldenFile1, goldenFile2)
}

func TestOpenFileMutation2(t *testing.T) {
	s, err := stateFileFromCommit(
		"https://github.com/richardimaoka/next-sandbox.git",
		"8adac375628219e020d4b5957ff24f45954cbd3f", // commit = 'npx create-next-app@latest'
		"next/package.json",
	)
	if err != nil {
		t.Fatalf("failed in TestOpenFileMutation1, %s", err)
	}

	// once GraphQL model is materialized...
	gqlModel := s.ToGraphQLOpenFile()
	goldenFile1 := "testdata/file_node_mutation_golden2-1.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile1, gqlModel)

	// ... mutation to the materialized model ...
	*gqlModel.FilePath = "next/package-mutated.json"
	*gqlModel.FileName = "package-mudated.json"
	*gqlModel.IsFullContent = false
	*gqlModel.Content = "mutated contents - " + *gqlModel.Content
	line100 := 100
	line200 := 200
	highlight := model.FileHighlight{
		FromLine: &line100,
		ToLine:   &line200,
	}
	gqlModel.Highlight = append(gqlModel.Highlight, &highlight)

	// ... has NO effect on a RE-materialized GraphQL model
	internal.CompareAfterMarshal(t, goldenFile1, s.ToGraphQLOpenFile())

	// ... has effect on the materialized GraphQL model
	goldenFile2 := "testdata/openfile_mutation_golden2-2.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile2, gqlModel)

	// and golden files are indeed different
	internal.FilesMustUnmatch(t, goldenFile1, goldenFile2)
}
