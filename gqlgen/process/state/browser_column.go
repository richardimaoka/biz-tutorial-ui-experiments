package state

import (
	"fmt"
	"image"
	"os"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

	_ "image/jpeg"
	_ "image/png"
)

type BrowserColumn struct {
	Width  int
	Height int
	Path   string
}

func NewBrowserColumn() *BrowserColumn {
	return &BrowserColumn{}
}

func (p *BrowserColumn) Process(tutorial, imageName string) error {
	if imageName == "" {
		return nil //keep the current state
	}

	// *Next.js <Image> requires a leading slash in path
	imagePath := "/images/" + tutorial + "/" + imageName
	nextJsImagePath := "../next/public" + imagePath

	width, height, err := p.ImageDimension(nextJsImagePath)
	if err != nil {
		return fmt.Errorf("ImageDimension() failed, %s", err)
	}

	// stateless, always new state
	p.Width = width
	p.Height = height
	p.Path = imagePath

	return nil
}

func (p *BrowserColumn) ImageDimension(imagePath string) (int, int, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return 0, 0, fmt.Errorf("ImageDimension() failed to open file, %s", err)
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}

	return image.Width, image.Height, nil
}

func (p *BrowserColumn) ToGraphQLBrowserCol() *model.BrowserColumn {
	// copy to avoid mutation effect afterwards
	width := p.Width
	height := p.Height
	path := stringRef(p.Path)

	return &model.BrowserColumn{
		Width:  &width,
		Height: &height,
		Path:   path,
	}
}
