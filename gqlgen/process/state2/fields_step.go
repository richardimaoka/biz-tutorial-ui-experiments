package state2

type IntrinsicFields struct {
	StepId  string `json:"stepId"`
	Comment string `json:"comment"`
}

type Step struct {
	/**
	 * Step meta fields
	 */
	FromRowFields // Fields to make the step searchable for re-generation
	ColumnFields
	IntrinsicFields
	ModalFields
	AnimationFields

	/**
	 * Fields for each column type
	 */
	SourceFields
	TerminalFields
}
