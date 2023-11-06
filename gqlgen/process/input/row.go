package input

type Column = string

const (
	// Lower cases since they are from manual entries
	Source   Column = "source"
	Terminal Column = "terminal"
	Browser  Column = "browser"
)

type SubType = string

const (
	// Lower cases since they are from manual entries
	Command SubType = "command"
	Output  SubType = "output"
)

type Row struct {
	StepId        string `json:"stepId"`
	Phase         string `json:"phase"`
	Comment       string `json:"comment"`
	Column        string `json:"column"` //not Column but string, because it's input from manual entry, not sanitized
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
