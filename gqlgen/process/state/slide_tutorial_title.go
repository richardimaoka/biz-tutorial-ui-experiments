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
	errorPrefix := "NewTutorialTitleSlide failed"
	imagePaths := strings.Split(fields.TutorialTitleImagePaths, "\n")
	captions := strings.Split(fields.TutorialTitleImageCaptions, "\n")

	var images []*Image
	for i := 0; i < len(imagePaths); i++ {
		width, err := fields.TutorialTitleImageWidths.Get(i)
		if err != nil {
			return nil, fmt.Errorf("%s, invalid width, %s", errorPrefix, err)
		}

		height, err := fields.TutorialTitleImageWidths.Get(i)
		if err != nil {
			return nil, fmt.Errorf("%s, invalid height, %s", errorPrefix, err)
		}

		image := NewSvgImage(tutorial, imagePaths[i], width, height, captions[i])
		if err := image.copyFile(); err != nil {
			return nil, fmt.Errorf("%s, %s", errorPrefix, err)
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
