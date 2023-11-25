package state

type SourceCodeTooltipTiming string

const (
	SOURCE_TOOLTIP_START SourceCodeTooltipTiming = "START"
	SOURCE_TOOLTIP_END   SourceCodeTooltipTiming = "END"
)

type SourceCodeTooltip struct {
	filePath     string
	markdownBody string
	timing       SourceCodeTooltipTiming
	lineNumber   int
}
