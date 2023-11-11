package state2

/**
 * Intrinsic step fields
 */

type IntrinsicFields struct {
	StepId  string `json:"stepId"`
	Comment string `json:"comment"`
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
	DurationSeconds int  `json:"durationSeconds"`
	IsTrivial       bool `json:"isTrivial"`
}
