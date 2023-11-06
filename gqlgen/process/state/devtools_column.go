package state

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type DevToolsColumn struct {
	Width  int
	Height int
	Path   string
}

func NewDevToolsColumn() *DevToolsColumn {
	return &DevToolsColumn{}
}

func (p *DevToolsColumn) Process(tutorial, imageName string, width, height int) error {
	if imageName == "" {
		return nil //keep the current state
	}
	if width <= 0 {
		return fmt.Errorf("Process() failed as width = %d is less than 1", width)
	}
	if height <= 0 {
		return fmt.Errorf("Process() failed as height = %d is less than 1", height)
	}

	// *Next.js <Image> requires a leading slash in path
	imagePath := "/images/" + tutorial + "/" + imageName

	// stateless, always new state
	p.Width = width
	p.Height = height
	p.Path = imagePath

	return nil
}

func (p *DevToolsColumn) ToGraphQLDevToolsCol() *model.DevToolsColumn {
	// copy to avoid mutation effect afterwards
	width := p.Width
	height := p.Height
	path := stringRef(p.Path)

	return &model.DevToolsColumn{
		Width:  &width,
		Height: &height,
		Path:   path,
	}
}
