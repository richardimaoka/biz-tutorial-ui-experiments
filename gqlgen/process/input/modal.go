package input

import (
	"fmt"
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

type ModalPosition string

const (
	MODAL_TOP    ModalPosition = "TOP"
	MODAL_CENTER ModalPosition = "CENTER"
	MODAL_BOTTOM ModalPosition = "BOTTOM"
)

func toModalPosition(s string) (ModalPosition, error) {
	switch strings.ToUpper(s) {
	case string(MODAL_TOP):
		return MODAL_TOP, nil
	case string(MODAL_CENTER):
		return MODAL_CENTER, nil
	case string(MODAL_BOTTOM):
		return MODAL_BOTTOM, nil
	case "": // default value
		return MODAL_TOP, nil
	default:
		return "", fmt.Errorf("ModalPosition value = '%s' is invalid", s)
	}
}

func (t ModalPosition) toState() state.ModalPosition {
	switch t {
	case MODAL_TOP:
		return state.MODAL_TOP
	case MODAL_CENTER:
		return state.MODAL_CENTER
	case MODAL_BOTTOM:
		return state.MODAL_BOTTOM
	default:
		panic(fmt.Sprintf("ModalPosition has an invalid value = '%s'", t))
	}
}
