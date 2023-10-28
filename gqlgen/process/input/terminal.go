package input

type TooltipTiming = string

const (
	BEGINNING TooltipTiming = "beginning"
	END       TooltipTiming = "end"
)

type TerminalTooltip struct {
	Contents string
	Timing   TooltipTiming
}

type TerminalCommand struct {
	StepId  string
	Comment string
	Command string
	Tooltip TerminalTooltip
}

type TerminalOutput struct {
	StepId  string
	Comment string
	Output  string
	Tooltip TerminalTooltip
}
