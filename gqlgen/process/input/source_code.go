package input

type SourceCodeTooltip struct {
	Contents           string             `json:"contents"`
	LineNumber         int                `json:"lineNumber"`
	PositionPreference PositionPreference `json:"positionPreference"`
	Timing             TooltipTiming      `json:"timing"`
}

type SourceCodeCommit struct {
	Comment         string
	Commit          string
	Tooltip         *SourceCodeTooltip
	TypingAnimation bool
	ShowDiff        bool
}

type SourceCodeOpen struct {
	Comment  string
	FilePath string
	Tooltip  *SourceCodeTooltip
}

type SourceCodeError struct {
	Comment  string
	FilePath string
	Tooltip  *SourceCodeTooltip
}
