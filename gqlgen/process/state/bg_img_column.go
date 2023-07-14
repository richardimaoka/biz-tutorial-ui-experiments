package state

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

type BackgroundImageColumn struct {
	Width  int
	Height int
	Path   string
	URL    string
	Modal  Modal
}

func NewBackgroundImageColumn(
	width int,
	height int,
	path string,
	modalText string,
	modalPosition ModalPosition,
) *BackgroundImageColumn {

	return &BackgroundImageColumn{
		Width:  width,
		Height: height,
		Path:   path,
		Modal: Modal{
			Text:     modalText,
			Position: modalPosition,
		},
	}
}

func (p *BackgroundImageColumn) ToGraphQLBgImgCol() *model.BackgroundImageColumn {
	// copy to avoid mutation effect afterwards
	width := p.Width
	height := p.Height
	path := internal.StringRef(p.Path)
	url := internal.StringRef(p.URL)
	modal := p.Modal.ToGraphQLModal() //ToGraphQLModal() performs copy internally, to avoid mutation effect afterwards

	return &model.BackgroundImageColumn{
		Width:  &width,
		Height: &height,
		Path:   path,
		URL:    url,
		Modal:  modal,
	}
}
