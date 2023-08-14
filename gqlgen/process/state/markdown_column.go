package state

import (
	"fmt"
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type MarkdownVerticalAlignment string

const (
	MarkdownVerticalTop    MarkdownVerticalAlignment = "TOP"
	MarkdownVerticalCenter MarkdownVerticalAlignment = "CENTER"
	MarkdownVerticalBottom MarkdownVerticalAlignment = "BOTTOM"
)

func ToMarkdownVerticalAlignment(a string) (MarkdownVerticalAlignment, error) {
	switch strings.ToUpper(a) {
	case "TOP":
		return MarkdownVerticalTop, nil
	case "CENTER":
		return MarkdownVerticalCenter, nil
	case "BOTTOM":
		return MarkdownVerticalBottom, nil
	default:
		return "", fmt.Errorf("'%s' is unknown MarkdownVerticalAlignment", a)
	}
}

func (this MarkdownVerticalAlignment) Convert() *model.ColumnVerticalPosition {
	a := model.ColumnVerticalPosition(this)
	if a.IsValid() {
		return &a
	} else {
		return nil
	}
}

type MarkdownColumn struct {
	Description       Markdown
	VerticalAlignment MarkdownVerticalAlignment
}

func NewMarkdownColumn() *MarkdownColumn {
	return &MarkdownColumn{}
}

func (p *MarkdownColumn) Process(markdownContents, verticalAlignment, horizontalAlignment string) error {
	va, err := ToMarkdownVerticalAlignment(verticalAlignment)
	if err != nil {
		return fmt.Errorf("Process() failed to convert vertical alignment, %s", err)
	}

	ha, err := ToMarkdownAlignment(horizontalAlignment)
	if err != nil {
		return fmt.Errorf("Process() failed to convert horizontal alignment, %s", err)
	}

	p.Description = Markdown{
		Contents:  markdownContents,
		Alignment: ha,
	}
	p.VerticalAlignment = va

	return nil
}

func (p *MarkdownColumn) ToGraphQLMarkdownColumn() *model.MarkdownColumn {
	// copy to avoid mutation effect afterwards
	description := p.Description.ToGraphQLMarkdown()
	verticalAlignment := p.VerticalAlignment.Convert()

	return &model.MarkdownColumn{
		Description:      description,
		ContentsPosition: verticalAlignment,
	}
}
