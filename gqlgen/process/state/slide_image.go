package state

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type ImageSlide struct {
	image *Image
}

func NewImageSlide(fields ImageFields, tutorial string) (*ImageSlide, error) {
	image, err := NewImage(
		tutorial,
		fields.ImagePath,
		fields.ImageWidth,
		fields.ImageHeight,
		fields.ImageCaption,
	)

	if err != nil {
		return nil, fmt.Errorf("NewImageSlide() failed, %s", err)
	}

	return &ImageSlide{
		image: image,
	}, nil
}

func (s *ImageSlide) ToGraphQLSlideWrapper() *model.SlideWrapper {
	slide := model.ImageSlide{
		Image: s.image.ToGraphQL(),
	}

	return &model.SlideWrapper{
		Slide: slide,
	}
}
