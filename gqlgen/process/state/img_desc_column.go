package state

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type ImageDescriptionColumn struct {
	Image       ImageCentered
	Description Markdown
}

func NewImgDescColumn(
	width int,
	height int,
	originalWidth int,
	originalHeight int,
	path string,
	descriptionContents string,
	descriptionAlignment MarkdownAlignment,
) *ImageDescriptionColumn {
	// TODO: if both path and url are passed, then error (or even better to create separate functions for path and url)

	return &ImageDescriptionColumn{
		Image: ImageCentered{
			Width:          width,
			Height:         height,
			OriginalWidth:  originalWidth,
			OriginalHeight: originalHeight,
			Path:           path,
		},
		Description: Markdown{
			Contents:  descriptionContents,
			Alignment: descriptionAlignment,
		},
	}
}

func (p *ImageDescriptionColumn) ToGraphQLImgDescCol() *model.ImageDescriptionColumn {
	// copy to avoid mutation effect afterwards
	image := p.Image.ToGraphQLImageCentered()
	description := p.Description.ToGraphQLMarkdown()

	return &model.ImageDescriptionColumn{
		Image:       image,
		Description: description,
	}
}
