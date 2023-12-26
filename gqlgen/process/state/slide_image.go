package state

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type ImageSlide struct {
	image Image
}

func NewImageSlide(fields ImageFields) *ImageSlide {
	return &ImageSlide{
		image: Image{
			src:     fields.ImagePath,
			width:   fields.ImageWidth,
			height:  fields.ImageHeight,
			caption: fields.ImageCaption,
		},
	}
}

func (s *ImageSlide) ToGraphQLSlideWrapper() *model.SlideWrapper {
	slide := model.ImageSlide{
		Image: s.image.ToGraphQL(),
	}

	return &model.SlideWrapper{
		Slide: slide,
	}
}
