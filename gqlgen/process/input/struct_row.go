package input

type Row struct {
	StepId  string `json:"stepId"`
	Phase   string `json:"phase"`
	Comment string `json:"comment"`
	Trivial string `json:"trivial"`

	// Mode and type fileds
	Mode    string `json:"mode"`
	RowType string `json:"rowType"`
	SubType string `json:"subtype"`

	// Modal fileds
	ModalContents string `json:"modalContents"`
	ModalPosition string `json:"modalPosition"`

	// Instruction fileds
	Instruction  string `json:"instruction"`
	Instruction2 string `json:"instruction2"`
	Instruction3 string `json:"instruction3"`

	// Tooltip fileds
	Tooltip       string      `json:"tooltip"`
	TooltipTiming string      `json:"tooltipTiming"`
	TooltipLine   IntOrString `json:"tooltipLine"`   // IntOrString because an empty cell from CSV becomes empty string = ""
	TooltipAppend string      `json:"tooltipAppend"` // string (i.e. 'TRUE', 'FALSE', or empty string), not bool, as it is a value from CSV
}
