package state_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/testio"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func TestMarkdownMutation1(t *testing.T) {
	md := state.Markdown{
		Contents:  "abcde",
		Alignment: state.MarkdownAlignmentLeft,
	}

	// once GraphQL model is materialized...
	gqlModel := md.ToGraphQLMarkdown()
	goldenFile1 := "testdata/markdown_golden1-1.json"
	testio.CompareWithGoldenFile(t, *updateFlag, goldenFile1, gqlModel)

	// ... mutation to the state ...
	md.Contents = "fghi"
	md.Alignment = state.MarkdownAlignmentCenter

	// ... has NO effect on the materialized GraphQL model
	testio.CompareAfterMarshal(t, goldenFile1, gqlModel)

	// ... has effect on a RE-materialized GraphQL model
	goldenFile2 := "testdata/markdown_golden1-2.json"
	testio.CompareWithGoldenFile(t, *updateFlag, goldenFile2, md.ToGraphQLMarkdown())
}

func TestMarkdownMutation2(t *testing.T) {
	md := state.Markdown{
		Contents:  "abcde",
		Alignment: state.MarkdownAlignmentLeft,
	}

	// once GraphQL model is materialized...
	gqlModel := md.ToGraphQLMarkdown()
	goldenFile1 := "testdata/markdown_golden2-1.json"
	testio.CompareWithGoldenFile(t, *updateFlag, goldenFile1, gqlModel)

	// ... mutation to the materialized model ...
	*gqlModel.Contents = "fghi"
	*gqlModel.Alignment = model.MarkdownAlignmentCenter

	// ... has NO effect on a RE-materialized GraphQL model
	testio.CompareAfterMarshal(t, goldenFile1, md.ToGraphQLMarkdown())

	// ... has effect on the materialized GraphQL model
	goldenFile2 := "testdata/markdown_golden2-2.json"
	testio.CompareWithGoldenFile(t, *updateFlag, goldenFile2, gqlModel)
}
