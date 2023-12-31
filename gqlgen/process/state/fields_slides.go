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
	TutorialTitle              string            `json:"tutorialTitle,omitempty"`
	TutorialTitleImagePaths    string            `json:"tutorialTitleImagePaths,omitempty"`
	TutorialTitleImageWidths   csvfield.MultiInt `json:"tutorialTitleImageWidths,omitempty"`
	TutorialTitleImageHeights  csvfield.MultiInt `json:"tutorialTitleImageHeights,omitempty"`
	TutorialTitleImageCaptions string            `json:"tutorialTitleImageCaptions,omitempty"`
}

/**
 * SectionTitle slide fields
 */
type SectionTitleFields struct {
	SectionTitle string `json:"sectionTitle,omitempty"`
}

/**
 * Markdown slide fields
 */
type MarkdownFields struct {
	MarkdownContents string `json:"markdownContents,omitempty"`
	// MarkdownVerticalAlignment   string `json:"markdownVerticalAlignment,omitempty"`
	// MarkdownHorizontalAlignment string `json:"markdownHorizontalAlignment,omitempty"`
}

/**
 * Image slide fields
 */
type ImageFields struct {
	ImagePath    string `json:"ImagePath,omitempty"`
	ImageWidth   int    `json:"ImageWidth,omitempty"`
	ImageHeight  int    `json:"ImageHeight,omitempty"`
	ImageCaption string `json:"ImageCaption,omitempty"`
}
