package input

import (
	"fmt"
	"strings"
)

type SlideType string

const (
	// Lower cases since they are from manual entries
	TutorialTitle SlideType = "tutorial title"
	SectionTitle  SlideType = "section title"
	Markdown      SlideType = "markdown"
	Image         SlideType = "image"
)

func toSlideType(s string) (SlideType, error) {
	lower := strings.ToLower(s)

	switch lower {
	case string(TutorialTitle):
		return TutorialTitle, nil
	case string(SectionTitle):
		return SectionTitle, nil
	case string(Markdown):
		return Markdown, nil
	case string(Image):
		return Image, nil
	default:
		return "", fmt.Errorf("'%s' is an invalid slide type", s)
	}
}
