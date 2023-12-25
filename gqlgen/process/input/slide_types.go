package input

import (
	"fmt"
	"strings"
)

type SlideType string

const (
	// Lower cases since they are from manual entries
	TutorialTitleSlide SlideType = "tutorial title"
	SectionTitleSlide  SlideType = "section title"
	MarkdownSlide      SlideType = "markdown"
	ImageSlide         SlideType = "image"
)

func toSlideType(s string) (SlideType, error) {
	lower := strings.ToLower(s)

	switch lower {
	case string(TutorialTitleSlide):
		return TutorialTitleSlide, nil
	case string(SectionTitleSlide):
		return SectionTitleSlide, nil
	case string(MarkdownSlide):
		return MarkdownSlide, nil
	case string(ImageSlide):
		return ImageSlide, nil
	default:
		return "", fmt.Errorf("'%s' is an invalid slide type", s)
	}
}
