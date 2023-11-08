package input

type ColumnType = string

const (
	// Lower cases since they are from manual entries
	SourceType   ColumnType = "source"
	TerminalType ColumnType = "terminal"
	BrowserType  ColumnType = "browser"
)

type Row struct {
	StepId        string `json:"stepId"`
	Phase         string `json:"phase"`
	Comment       string `json:"comment"`
	Trivial       string `json:"trivial"`
	ModalContents string `json:"modalContents"`

	// Not Column but string, because it's input from manual entry, not sanitized
	Column string `json:"column"`
	Type   string `json:"type"`

	// Instruction fileds
	Instruction  string `json:"instruction"`
	Instruction2 string `json:"instruction2"`
	Instruction3 string `json:"instruction3"`

	// Tooltip fileds
	Tooltip string `json:"tooltip"`
	// TODO: those tooltip fields are specific to source code, so move them to instruction?
	TooltipTiming string `json:"tooltipTiming"`
	TooltipLine   int    `json:"tooltipLine"`
	// string (i.e. 'TRUE', 'FALSE', or empty string), not bool, as it is a value from CSV
	TooltipAppend string `json:"tooltipAppend"`
}
