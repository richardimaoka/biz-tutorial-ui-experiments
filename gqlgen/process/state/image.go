package state

import (
	"fmt"
	"os"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type Image struct {
	tutorial string
	src      string
	width    int
	height   int
	caption  string
}

func NewImage(tutorial string,
	src string,
	width int,
	height int,
	caption string) *Image {

	return &Image{tutorial: tutorial,
		src:     src,
		width:   width,
		height:  height,
		caption: caption,
	}
}

func (i *Image) copyFile() error {
	sourcePath := fmt.Sprintf("data/%s/images/%s", i.tutorial, i.src)

	bytes, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to copy file %s, %s", i.src, err)
	}

	destPath := fmt.Sprintf("../next/public/images/%s/%s", i.tutorial, i.src)
	if err := os.WriteFile(destPath, bytes, 666); err != nil {
		return fmt.Errorf("failed to copy file %s, %s", i.src, err)
	}

	return nil
}

func (i *Image) ToGraphQL() *model.Image {
	src := fmt.Sprintf("/images/%s/%s", i.tutorial, i.src)
	width := i.width
	height := i.width
	caption := stringRef(i.caption)

	return &model.Image{
		Src:     src,
		Width:   width,
		Height:  height,
		Caption: caption,
	}
}
