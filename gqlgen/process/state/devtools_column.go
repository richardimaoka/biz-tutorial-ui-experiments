package state

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

type DevToolsColumn struct {
	Width  int
	Height int
	Path   string
}

func NewDevToolsColumn(width, height int, path string) *DevToolsColumn {
	return &DevToolsColumn{
		Width:  width,
		Height: height,
		Path:   path,
	}
}

func (p *DevToolsColumn) ToGraphQLDevToolsCol() *model.DevToolsColumn {
	// copy to avoid mutation effect afterwards
	width := p.Width
	height := p.Height
	path := internal.StringRef(p.Path)

	return &model.DevToolsColumn{
		Width:  &width,
		Height: &height,
		Path:   path,
	}
}
