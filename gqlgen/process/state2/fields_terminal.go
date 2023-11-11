package state2

type TerminalStepType string

const (
	TerminalCommand TerminalStepType = "command"
	TerminalOutput  TerminalStepType = "output"
	TerminalCd      TerminalStepType = "cd"
	TerminalMove    TerminalStepType = "move"
	TerminalOpen    TerminalStepType = "open"
)

type TerminalTooltipFields struct {
	TerminalTooltipContents string `json:"terminalTooltipContents"`
	TerminalTooltipTiming   string `json:"terminalTooltipTiming"`
}

type TerminalFields struct {
	CurrentDir       string           `json:"currentDir"`
	TerminalStepType TerminalStepType `json:"terminalType"`
	TerminalText     string           `json:"terminalText"`
	TerminalName     string           `json:"terminalName"`
	// embed tooltip
	TerminalTooltipFields
}
