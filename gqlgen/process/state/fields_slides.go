package state

/**
 * Column fields
 */

type SlideType string

const (
	// `Type` suffix is needed to avoid conflict with structs
	NoSlideType            SlideType = ""
	TutorialTitleSlideType SlideType = "TutorialSlide"
)

/**
 * TutorialTitle slide fields
 */
type TutorialTitleFields struct {
	TutorialTitle              string `json:"tutorialTitle"`
	TutorialTitleImageFiles    string `json:"tutorialTitleImageFiles"`
	TutorialTitleImageSizes    string `json:"tutorialTitleImageSizes"`
	TutorialTitleImageCaptions string `json:"tutorialTitleImageCaptions"`
}

/**
 * SectionTitle slide fields
 */
type SectionTitleFields struct {
	SectionTitle string `json:"sectionTitle"`
}

/**
 * Markdown slide fields
 */
type MarkdownFields struct {
	MarkdownContents string `json:"markdownContents"`
	// MarkdownVerticalAlignment   string `json:"markdownVerticalAlignment"`
	// MarkdownHorizontalAlignment string `json:"markdownHorizontalAlignment"`
}
