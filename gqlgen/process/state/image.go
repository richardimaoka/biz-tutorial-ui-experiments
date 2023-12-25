package state

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type Image struct {
	src     string
	width   int
	height  int
	caption string
}

func (i *Image) ToGraphQL() *model.Image {
	src := i.src
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
