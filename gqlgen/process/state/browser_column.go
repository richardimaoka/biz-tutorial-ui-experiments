package state

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

type BrowserColumn struct {
	Width  int
	Height int
	Path   string
}

func NewBrowserColumn() *BrowserColumn {
	return &BrowserColumn{}
}

func (p *BrowserColumn) Process(tutorial, imageName string, width, height int) error {
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
