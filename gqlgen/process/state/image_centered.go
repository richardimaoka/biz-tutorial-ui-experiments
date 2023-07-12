package state

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
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
	path := internal.StringRef(p.Path)
	url := internal.StringRef(p.URL)

	return &model.ImageCentered{
		Width:  &width,
		Height: &height,
		Path:   path,
		URL:    url,
	}
}
