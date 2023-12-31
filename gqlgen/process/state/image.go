package state

import (
	"fmt"
	"image"
	"os"
	"strings"

	_ "image/jpeg"
	_ "image/png"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type Image struct {
	tutorial string
	src      string
	width    int
	height   int
	caption  string
}

func sourceImagePath(tutorial, src string) string {
	if strings.Contains(src, "testdata") {
		return src
	} else {
		return fmt.Sprintf("data/%s/images/%s", tutorial, src)
	}
}

func destinationImagePath(tutorial, src string) string {
	return fmt.Sprintf("../next/public/images/%s/%s", tutorial, src)
}

func NewImage(
	tutorial string,
	src string,
	width int,
	height int,
	caption string) (*Image, error) {
	errorPrefix := "NewImage() failed"

	if width == 0 || height == 0 {
		// Either width or height is the zero value, then figure it out from the image file
		sourcePath := sourceImagePath(tutorial, src)
		reader, err := os.Open(sourcePath)
		defer reader.Close()
		if err != nil {
			return nil, fmt.Errorf("%s, failed to open image file = %s, %s", errorPrefix, src, err)
		}

		imgConfig, _, err := image.DecodeConfig(reader)
		if err != nil {
			return nil, fmt.Errorf("%s, failed to get width/height of image file = %s, %s", errorPrefix, src, err)
		}

		return &Image{
			tutorial: tutorial,
			src:      src,
			width:    imgConfig.Width,
			height:   imgConfig.Height,
			caption:  caption,
		}, nil

	} else {
		// If both width and height are given, then use them
		return &Image{
			tutorial: tutorial,
			src:      src,
			width:    width,
			height:   height,
			caption:  caption,
		}, nil
	}
}

func (i *Image) copyFile() error {
	sourcePath := sourceImagePath(i.tutorial, i.src)
	destPath := destinationImagePath(i.tutorial, i.src)

	bytes, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to copy file %s, %s", i.src, err)
	}

	if err := os.WriteFile(destPath, bytes, 0666); err != nil {
		return fmt.Errorf("failed to copy file %s, %s", i.src, err)
	}

	return nil
}

func (i *Image) ToGraphQL() *model.Image {
	src := fmt.Sprintf("/images/%s/%s", i.tutorial, i.src)
	width := i.width
	height := i.height
	caption := stringRef(i.caption)

	return &model.Image{
		Src:     src,
		Width:   width,
		Height:  height,
		Caption: caption,
	}
}
