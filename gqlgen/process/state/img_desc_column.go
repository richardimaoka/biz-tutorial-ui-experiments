package state

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type ImageDescriptionColumn struct {
	Image       ImageCentered
	Description Markdown
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
