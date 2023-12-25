package state

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type Image struct {
	Src    string
	Width  int
	Height int
	Captin string
}

type SectionTitleSlide struct {
	SectionNum int
	Title      string
}

func (s *SectionTitleSlide) ToGraphQLSlideWrapper() *model.SlideWrapper {
	slide := model.SectionTitleSlide{
		SectionNum: s.SectionNum,
		Title:      s.Title,
	}

	return &model.SlideWrapper{
		Slide: slide,
	}
}
