package state

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type MarkdownColumn struct {
	Description Markdown
}

func NewMarkdownColumn() *MarkdownColumn {
	return &MarkdownColumn{}
}

func (p *MarkdownColumn) Process(markdownContents string, alignment string) error {
	markdownAlignment, err := ToMarkdownAlignment(alignment)
	if err != nil {
		return fmt.Errorf("Process() failed to convert alignment, %s", err)
	}

	p.Description = Markdown{
		Contents:  markdownContents,
		Alignment: markdownAlignment,
	}

	return nil
}

func (p *MarkdownColumn) ToGraphQLMarkdownColumn() *model.MarkdownColumn {
	// copy to avoid mutation effect afterwards
	description := p.Description.ToGraphQLMarkdown()

	return &model.MarkdownColumn{
		Description: description,
	}
}
