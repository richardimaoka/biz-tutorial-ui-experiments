package state

import (
	"fmt"
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
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

func ToModalPosition(p string) (ModalPosition, error) {
	switch strings.ToUpper(p) {
	case "TOP":
		return ModalPositionTop, nil
	case "CENTER":
		return ModalPositionCenter, nil
	case "BOTTOM":
		return ModalPositionBottom, nil
	default:
		return "", fmt.Errorf("'%s' is unknown ModalPosition", p)
	}
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
	text := stringRef(p.Text)
	position := convertModalPosition(p.Position) //p.Position is passed-by-copy, to avoid mutation effect afterwards

	return &model.Modal{
		Text:     text,
		Position: position,
	}
}
