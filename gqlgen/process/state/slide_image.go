package state

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type ImageSlide struct {
	image *Image
}

func NewImageSlide(fields ImageFields, tutorial string) *ImageSlide {
	return &ImageSlide{
		image: NewImage(
			tutorial,
			fields.ImagePath,
			fields.ImageWidth,
			fields.ImageHeight,
			fields.ImageCaption,
		),
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
