package state

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

type MarkdownAlignment string

const (
	MarkdownAlignmentLeft   MarkdownAlignment = "LEFT"
	MarkdownAlignmentCenter MarkdownAlignment = "CENTER"
)

func (this MarkdownAlignment) Convert() *model.MarkdownAlignment {
	a := model.MarkdownAlignment(this)
	if a.IsValid() {
		return &a
	} else {
		return nil
	}
}

type Markdown struct {
	Contents  string
	Alignment MarkdownAlignment
}

func (p *Markdown) ToGraphQLMarkdown() *model.Markdown {
	// copy to avoid mutation effect afterwards
	contents := internal.StringRef(p.Contents)
	alignment := p.Alignment.Convert()

	return &model.Markdown{
		Contents:  contents,
		Alignment: alignment,
	}
}
