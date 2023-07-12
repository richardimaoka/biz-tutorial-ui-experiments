package state

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type MarkdownColumn struct {
	Image       ImageCentered
	Description Markdown
}

func (p *MarkdownColumn) ToGraphQLMarkdownColumn() *model.MarkdownColumn {
	// copy to avoid mutation effect afterwards
	description := p.Description.ToGraphQLMarkdown()

	return &model.MarkdownColumn{
		Description: description,
	}
}
