package state_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/testio"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func TestMarkdownColumnMutation1(t *testing.T) {
	s := state.MarkdownColumn{
		Description: state.Markdown{
			Contents:  "markdown default contents",
			Alignment: state.MarkdownAlignmentLeft,
		},
	}

	// once GraphQL model is materialized...
	gqlModel := s.ToGraphQLMarkdownColumn()
	goldenFile1 := "testdata/md_column_golden1-1.json"
	testio.CompareWithGoldenFile(t, *updateFlag, goldenFile1, gqlModel)

	// ... mutation to the state ...
	s.Description.Contents = "markdown updated contents"
	s.Description.Alignment = state.MarkdownAlignmentCenter

	// ... has NO effect on the materialized GraphQL model
	testio.CompareAfterMarshal(t, goldenFile1, gqlModel)

	// ... has effect on a RE-materialized GraphQL model
	goldenFile2 := "testdata/md_column_golden1-2.json"
	testio.CompareWithGoldenFile(t, *updateFlag, goldenFile2, s.ToGraphQLMarkdownColumn())
}

func TestMarkdownColumnMutation2(t *testing.T) {
	s := state.MarkdownColumn{
		Description: state.Markdown{
			Contents:  "markdown default contents",
			Alignment: state.MarkdownAlignmentLeft,
		},
	}

	// once GraphQL model is materialized...
	gqlModel := s.ToGraphQLMarkdownColumn()
	goldenFile1 := "testdata/md_column_golden2-1.json"
	testio.CompareWithGoldenFile(t, *updateFlag, goldenFile1, gqlModel)

	// ... mutation to the state ...
	*gqlModel.Description.Contents = "markdown updated contents"
	*gqlModel.Description.Alignment = model.MarkdownAlignmentCenter

	// ... has NO effect on a RE-materialized GraphQL model
	testio.CompareAfterMarshal(t, goldenFile1, s.ToGraphQLMarkdownColumn())

	// ... has effect on the materialized GraphQL model
	goldenFile2 := "testdata/md_column_golden2-2.json"
	testio.CompareWithGoldenFile(t, *updateFlag, goldenFile2, gqlModel)
}
