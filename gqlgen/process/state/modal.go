package state

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type ModalPosition string

const (
	MODAL_TOP    ModalPosition = "TOP"
	MODAL_CENTER ModalPosition = "CENTER"
	MODAL_BOTTOM ModalPosition = "BOTTOM"
)

func (t ModalPosition) toGraphQL() model.ModalPosition {
	switch t {
	case MODAL_TOP:
		return model.ModalPositionTop
	case MODAL_CENTER:
		return model.ModalPositionCenter
	case MODAL_BOTTOM:
		return model.ModalPositionBottom
	default:
		panic(fmt.Sprintf("ModalPosition has an invalid value = '%s'", t))
	}
}

type Modal struct {
	markdownBody string
	position     ModalPosition
}

func (m *Modal) ToGraphQL() *model.Modal {
	position := m.position.toGraphQL()

	return &model.Modal{
		MarkdownBody: m.markdownBody,
		Position:     &position,
	}
}
