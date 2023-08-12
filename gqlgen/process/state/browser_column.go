package state

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

type BrowserColumn struct {
	Width  int
	Height int
	Path   string
}

func NewBrowserColumn(width, height int, path string) *BrowserColumn {
	return &BrowserColumn{
		Width:  width,
		Height: height,
		Path:   path,
	}
}

func (p *BrowserColumn) ToGraphQLBrowserCol() *model.BrowserColumn {
	// copy to avoid mutation effect afterwards
	width := p.Width
	height := p.Height
	path := internal.StringRef(p.Path)

	return &model.BrowserColumn{
		Width:  &width,
		Height: &height,
		Path:   path,
	}
}
