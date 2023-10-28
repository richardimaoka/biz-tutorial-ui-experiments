package input

type Abstract struct {
	StepId          string `json:"stepId"`
	Phase           string `json:"phase"`
	Column          string `json:"column"`
	Type            string `json:"type"`
	Trivial         string `json:"trivial"`
	Instruction     string `json:"instruction"`
	Instruction2    string `json:"instruction2"`
	Instruction3    string `json:"instruction3"`
	ModalText       string `json:"modalText"`
	Tooltip         string `json:"tooltip"`
	TooltipTiming   string `json:"tooltipTiming"`
	TooltipPosition string `json:"tooltipPosition"`
	Comment         string `json:"comment"`
}
