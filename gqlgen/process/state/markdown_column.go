package state

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type MarkdownColumn struct {
	Description Markdown
}

func NewMarkdownColumn(markdownContents string, alignment MarkdownAlignment) *MarkdownColumn {
	return &MarkdownColumn{
		Description: Markdown{
			Contents:  markdownContents,
			Alignment: alignment,
		},
	}
}

func (p *MarkdownColumn) ToGraphQLMarkdownColumn() *model.MarkdownColumn {
	// copy to avoid mutation effect afterwards
	description := p.Description.ToGraphQLMarkdown()

	return &model.MarkdownColumn{
		Description: description,
	}
}
