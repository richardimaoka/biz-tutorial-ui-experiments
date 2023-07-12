package process

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/read"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func ToStateBgImgColumn(e *read.BackgroundImageColumn) *state.BackgroundImageColumn {
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

func ToGraphQLBgImgCol(e *read.BackgroundImageColumn) *model.BackgroundImageColumn {
	return ToStateBgImgColumn(e).ToGraphQLBgImgCol()
}
