package process

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/read"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func ToStateMarkdownColumn(e read.MarkdownColumn) *state.MarkdownColumn {
	return &state.MarkdownColumn{
		Description: state.Markdown{
			Contents:  e.DescriptionContents,
			Alignment: state.MarkdownAlignment(e.DescriptionAlignment),
		},
	}
}

func ToGraphQLMarkdownColumn(e read.MarkdownColumn) *model.MarkdownColumn {
	return ToStateMarkdownColumn(e).ToGraphQLMarkdownColumn()
}
