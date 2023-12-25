package state

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

/**
 * Slide fields
 */

type Mode string

const (
	// `Type` suffix is needed to avoid conflict with structs
	SlideshowMode Mode = "Slideshow"
	HandsonMode   Mode = "Handson"
)

func (m Mode) ToGraphQL() model.PageMode {
	switch m {
	case SlideshowMode:
		return model.PageModeSlideshow
	case HandsonMode:
		return model.PageModeHandson
	default:
		panic(fmt.Sprintf("mode = '%s' is invalid", m))
	}
}

/**
 * Intrinsic step fields
 */

type IntrinsicFields struct {
	StepId      string     `json:"stepId"`
	Comment     string     `json:"comment"`
	Mode        Mode       `json:"mode"`
	FocusColumn ColumnType `json:"focusColumn"`
	SlideType   SlideType  `json:"slideType"`
}

/**
 * From row fields.  to make the step searchable for re-generation
 */
type FromRowFields struct {
	IsFromRow   bool   `json:"isFromRow"`
	SubID       string `json:"subId"`
	ParentRowId string `json:"parentRowId"`
}

/**
 * Modal fields
 */
type ModalFields struct {
	ModalContents string `json:"modalContents"`
}

/**
 * Animation fields
 */
type AnimationFields struct {
	// DurationSeconds int  `json:"durationSeconds"`
	IsTrivial bool `json:"isTrivial"`
}
