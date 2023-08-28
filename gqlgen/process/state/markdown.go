package state

import (
	"fmt"
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

type MarkdownAlignment string

const (
	MarkdownAlignmentLeft   MarkdownAlignment = "LEFT"
	MarkdownAlignmentCenter MarkdownAlignment = "CENTER"
)

func ToMarkdownAlignment(a string) (MarkdownAlignment, error) {
	switch strings.ToUpper(a) {
	case "LEFT":
		return MarkdownAlignmentLeft, nil
	case "CENTER":
		return MarkdownAlignmentCenter, nil
	default:
		return "", fmt.Errorf("'%s' is unknown MarkdownAlignment", a)
	}
}

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
	s := "str"
	alignment := &s

	return &model.Markdown{
		Contents:  contents,
		Alignment: alignment,
	}
}
