package state

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type MarkdownSlide struct {
	markdownBody string
}

func NewMarkdownSlide(fields MarkdownFields) *MarkdownSlide {
	return &MarkdownSlide{markdownBody: fields.MarkdownContents}
}

func (s *MarkdownSlide) ToGraphQLSlideWrapper() *model.SlideWrapper {
	slide := model.MarkdownSlide{
		MarkdownBody: s.markdownBody,
	}

	return &model.SlideWrapper{
		Slide: slide,
	}
}
