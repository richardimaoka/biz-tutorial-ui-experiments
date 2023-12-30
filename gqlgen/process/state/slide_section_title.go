package state

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type SectionTitleSlide struct {
	sectionNum int
	title      string
}

func NewSectionTitleSlide() *SectionTitleSlide {
	return &SectionTitleSlide{sectionNum: 0, title: ""}
}

func (s *SectionTitleSlide) Update(fields SectionTitleFields) {
	s.title = fields.SectionTitle
	s.sectionNum++
}

func (s *SectionTitleSlide) ToGraphQLSlideWrapper() *model.SlideWrapper {
	slide := model.SectionTitleSlide{
		SectionNum: s.sectionNum,
		Title:      s.title,
	}

	return &model.SlideWrapper{
		Slide: slide,
	}
}
