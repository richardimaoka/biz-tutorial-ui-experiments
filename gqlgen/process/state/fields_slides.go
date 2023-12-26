package state

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/csvfield"

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

/**
 * TutorialTitle slide fields
 */
type TutorialTitleFields struct {
	TutorialTitle              string            `json:"tutorialTitle"`
	TutorialTitleImagePaths    string            `json:"tutorialTitleImagePaths"`
	TutorialTitleImageWidths   csvfield.MultiInt `json:"tutorialTitleImageWidths"`
	TutorialTitleImageHeights  csvfield.MultiInt `json:"tutorialTitleImageHeights"`
	TutorialTitleImageCaptions string            `json:"tutorialTitleImageCaptions"`
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
	ImageWidth   int    `json:"ImageWidth"`
	ImageHeight  int    `json:"ImageHeight"`
	ImageCaption string `json:"ImageCaption"`
}
