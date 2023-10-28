package input

type SourceCodeTooltip struct {
	Contents string
}

type SourceCodeCommit struct {
	Comment         string
	Commit          string
	Tooltip         string
	TypingAnimation bool
	ShowDiff        bool
}
