package state

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

type ModalPosition string

const (
	ModalPositionTop    ModalPosition = "TOP"
	ModalPositionCenter ModalPosition = "CENTER"
	ModalPositionBottom ModalPosition = "BOTTOM"
)

type Modal struct {
	Text     string
	Position ModalPosition
}

func convertModalPosition(pos ModalPosition) *model.ModalPosition {
	p := model.ModalPosition(pos)
	if p.IsValid() {
		return &p
	} else {
		return nil
	}
}

func (p *Modal) ToGraphQLModal() *model.Modal {
	// if no text, then no need to show modal
	if p.Text == "" {
		return nil
	}

	// copy to avoid mutation effect afterwards
	text := internal.StringRef(p.Text)
	position := convertModalPosition(p.Position) //p.Position is passed-by-copy, to avoid mutation effect afterwards

	return &model.Modal{
		Text:     text,
		Position: position,
	}
}