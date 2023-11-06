package input

type Row struct {
	StepId        string `json:"stepId"`
	Phase         string `json:"phase"`
	Comment       string `json:"comment"`
	Column        string `json:"column"`
	Type          string `json:"type"`
	Trivial       string `json:"trivial"`
	Instruction   string `json:"instruction"`
	Instruction2  string `json:"instruction2"`
	Instruction3  string `json:"instruction3"`
	ModalText     string `json:"modalText"`
	Tooltip       string `json:"tooltip"`
	TooltipTiming string `json:"tooltipTiming"`
	TooltipLine   int    `json:"tooltipLine"`
	// TooltipPosition string `json:"tooltipPosition"`
}
