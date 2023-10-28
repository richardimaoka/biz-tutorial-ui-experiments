package input

type PositionPreference = string

const (
	ABOVE PositionPreference = "ABOVE"
	BELOW PositionPreference = "BELOW"
	EXACT PositionPreference = "EXACT"
)

type SourceCodeTooltip struct {
	Contents           string
	LineNumber         int
	PositionPreference PositionPreference
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
