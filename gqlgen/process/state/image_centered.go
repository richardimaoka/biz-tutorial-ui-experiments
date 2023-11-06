package state

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type ImageCentered struct {
	Width          int
	Height         int
	OriginalWidth  int
	OriginalHeight int
	Path           string
	URL            string
}

func (p *ImageCentered) ToGraphQLImageCentered() *model.ImageCentered {
	// copy to avoid mutation effect afterwards
	width := p.Width
	height := p.Height
	path := stringRef(p.Path)
	url := stringRef(p.URL)

	return &model.ImageCentered{
		Width:  &width,
		Height: &height,
		Path:   path,
		URL:    url,
	}
}
