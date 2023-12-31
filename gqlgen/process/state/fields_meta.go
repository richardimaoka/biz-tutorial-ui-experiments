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
	StepId      string     `json:"stepId,omitempty"`
	Comment     string     `json:"comment,omitempty"`
	Mode        Mode       `json:"mode,omitempty"`
	FocusColumn ColumnType `json:"focusColumn,omitempty"`
	SlideType   SlideType  `json:"slideType,omitempty"`
}

/**
 * From row fields.  to make the step searchable for re-generation
 */
type FromRowFields struct {
	IsFromRow   bool   `json:"isFromRow,omitempty"`
	SubID       string `json:"subId,omitempty"`
	ParentRowId string `json:"parentRowId,omitempty"`
}

/**
 * Modal fields
 */
type ModalFields struct {
	ModalContents string        `json:"modalContents,omitempty"`
	ModalPosition ModalPosition `json:"modalPosition,omitempty"`
}

/**
 * Animation fields
 */
type AnimationFields struct {
	// DurationSeconds int  `json:"durationSeconds,omitempty"`
	IsTrivial bool `json:"isTrivial,omitempty"`
}
