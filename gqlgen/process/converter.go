package process

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/read"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func ToStateBgImgColumn(e read.BackgroundImageColumn) *state.BackgroundImageColumn {
	return &state.BackgroundImageColumn{
		Width:  e.Width,
		Height: e.Height,
		Path:   e.Path,
		Modal: state.Modal{
			Text:     e.ModalText,
			Position: state.ModalPosition(e.ModalPosition),
		},
	}
}

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

func ToStateMarkdownColumn(e read.MarkdownColumn) *state.MarkdownColumn {
	return &state.MarkdownColumn{
		Description: state.Markdown{
			Contents:  e.DescriptionContents,
			Alignment: state.MarkdownAlignment(e.DescriptionAlignment),
		},
	}
}
