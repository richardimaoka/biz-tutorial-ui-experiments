package state

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type ModalPosition string

const (
	MODAL_TOP    ModalPosition = "TOP"
	MODAL_CENTER ModalPosition = "CENTER"
	MODAL_BOTTOM ModalPosition = "BOTTOM"
)

func (p *ModalPosition) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return fmt.Errorf("failed in ModalPosition UnmarshalJSON(), %w", err)
	}

	switch strings.ToUpper(s) {
	case string(MODAL_TOP):
		*p = MODAL_TOP
	case string(MODAL_CENTER):
		*p = MODAL_CENTER
	case string(MODAL_BOTTOM):
		*p = MODAL_BOTTOM
	case "": // default value - this is necessary as both modalContents and modalPosition can be zero values
		*p = MODAL_TOP
	default:
		panic(fmt.Sprintf("failed in ModalPosition UnmarshalJSON(), '%s' is invalid", s))
	}

	return nil
}

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
