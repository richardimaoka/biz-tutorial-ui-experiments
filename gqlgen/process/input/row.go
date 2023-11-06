package input

type ColumnType = string

const (
	// Lower cases since they are from manual entries
	SourceType   ColumnType = "source"
	TerminalType ColumnType = "terminal"
	BrowserType  ColumnType = "browser"
)

type SubType = string

const (
	// Lower cases since they are from manual entries
	CommandSubType SubType = "command"
	OutputSubType  SubType = "output"
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
	ModalContents string `json:"modalContents"`
	Tooltip       string `json:"tooltip"`
	TooltipTiming string `json:"tooltipTiming"`
	TooltipLine   int    `json:"tooltipLine"`
	// TooltipPosition string `json:"tooltipPosition"`
}
