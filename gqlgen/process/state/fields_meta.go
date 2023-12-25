package state

/**
 * Slide fields
 */

type Mode string

const (
	// `Type` suffix is needed to avoid conflict with structs
	SlideshowMode Mode = "Slideshow"
	HandsonMode   Mode = "Handson"
)

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
	IsFromRow  bool   `json:"isFromRow"`
	SubID      string `json:"subId"`
	ParentStep string `json:"parentStep"`
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
