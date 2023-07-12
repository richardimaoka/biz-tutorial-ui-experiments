package process

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/read"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func ToStateImgDescColumn(e read.ImageDescriptionColumn) *state.ImageDescriptionColumn {
	return &state.ImageDescriptionColumn{
		Image: state.ImageCentered{
			Width:          e.Width,
			Height:         e.Height,
			OriginalWidth:  e.OriginalWidth,
			OriginalHeight: e.OriginalHeight,
			Path:           e.Path,
		},
		Description: state.Markdown{
			Contents:  e.DescriptionContents,
			Alignment: state.MarkdownAlignment(e.DescriptionAlignment),
		},
	}
}

func ToGraphQLImgDescCol(e read.ImageDescriptionColumn) *model.ImageDescriptionColumn {
	return ToStateImgDescColumn(e).ToGraphQLImgDescCol()
}
