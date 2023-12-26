package state

import (
	"fmt"
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type TutorialTitleSlide struct {
	title  string
	images []*Image
}

func NewTutorialTitleSlide(fields TutorialTitleFields, tutorial string) (*TutorialTitleSlide, error) {
	imagePaths := strings.Split(fields.TutorialTitleImagePaths, "\n")

	var images []*Image
	for i := 0; i < len(imagePaths); i++ {
		image := NewImage(tutorial, imagePaths[i], 0, 0, "")
		if err := image.copyFile(); err != nil {
			return nil, fmt.Errorf("NewTutorialTitleSlide faield, %s", err)
		}

		images = append(images, image)
	}

	return &TutorialTitleSlide{
		title:  fields.TutorialTitle,
		images: images,
	}, nil
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
