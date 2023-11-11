package state2

type SourceStepType string

const (
	FileTree         SourceStepType = "fileTree"
	SourceOpen       SourceStepType = "openFile"
	SourceOpenCommit SourceStepType = "openFileCommit"
	SourceError      SourceStepType = "error"
	SourceMove       SourceStepType = "move"
)

type SourceTooltipFields struct {
	SourceTooltipContents   string `json:"sourceTooltipContents"`
	SourceTooltipTiming     string `json:"sourceTooltipTiming"`
	SourceTooltipLineNumber int    `json:"sourceTooltipLineNumber"`
	SourceTooltipIsAppend   bool   `json:"SourceTooltipIsAppend"`
}

type SourceFields struct {
	Commit              string `json:"commit"`
	DefaultOpenFilePath string `json:"defaultOpenFilePath"`
	ShowFileTree        bool   `json:"showFileTree"`
	TypingAnimation     bool   `json:"typingAnimation"`
	// embed tooltip
	SourceTooltipFields
}
