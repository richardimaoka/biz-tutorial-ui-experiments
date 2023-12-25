package state

/**
 * Slide fields
 */

type SlideType string

const (
	// `Type` suffix is needed to avoid conflict with structs
	NoSlideType            SlideType = ""
	TutorialTitleSlideType SlideType = "TutorialTitle"
	SectionTitleSlideType  SlideType = "SectionTitle"
	TocSlideType           SlideType = "Toc"
	MarkdownSlideType      SlideType = "Markdown"
	ImageSlideType         SlideType = "Image"
)

type SlideFields struct {
	SlideType SlideType `json:"slideType"`
}

/**
 * TutorialTitle slide fields
 */
type TutorialTitleFields struct {
	TutorialTitle              string `json:"tutorialTitle"`
	TutorialTitleImagePaths    string `json:"tutorialTitleImagePaths"`
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

/**
 * Image slide fields
 */
type ImageFields struct {
	ImagePath    string `json:"ImagePath"`
	ImageSize    string `json:"ImageSize"`
	ImageCaption string `json:"ImageCaption"`
}
