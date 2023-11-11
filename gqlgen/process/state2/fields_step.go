package state2

type IntrinsicFields struct {
	StepId  string `json:"stepId"`
	Comment string `json:"comment"`
}

type Step struct {
	IntrinsicFields
	ColumnFields
	SourceFields
	TerminalFields
}
