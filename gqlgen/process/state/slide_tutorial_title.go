package state

import (
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type TutorialTitleSlide struct {
	title  string
	images []Image
}

func NewTutorialTitleSlide(fields TutorialTitleFields) *TutorialTitleSlide {
	imagePaths := strings.Split(fields.TutorialTitleImagePaths, "\n")
	var images []Image
	for i := 0; i < len(imagePaths); i++ {
		images = append(images, Image{src: imagePaths[i]})
	}

	return &TutorialTitleSlide{
		title:  fields.TutorialTitle,
		images: images,
	}
}

func (s *TutorialTitleSlide) ToGraphQLSlideWrapper() *model.SlideWrapper {
	var images []*model.Image
	for _, v := range s.images {
		images = append(images, v.ToGraphQL())
	}

	slide := model.TutorialTitleSlide{
		Title:  s.title,
		Images: images,
	}

	return &model.SlideWrapper{
		Slide: slide,
	}
}
