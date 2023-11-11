package result

type ColumnType string

const (
	NoColumnType       ColumnType = ""
	SourceColumnType   ColumnType = "SourceCode"
	TerminalColumnType ColumnType = "Terminal"
	BrowserColumnType  ColumnType = "Browser"
)

type TerminalStepType string

const (
	TerminalCommand TerminalStepType = "command"
	TerminalOutput  TerminalStepType = "output"
	TerminalCd      TerminalStepType = "cd"
	TerminalMove    TerminalStepType = "move"
	TerminalOpen    TerminalStepType = "open"
)

type SourceStepType string

const (
	FileTree         SourceStepType = "fileTree"
	SourceOpen       SourceStepType = "openFile"
	SourceOpenCommit SourceStepType = "openFileCommit"
	SourceError      SourceStepType = "error"
	SourceMove       SourceStepType = "move"
)

type BrowserType string

const (
	BrowserOpen BrowserType = "open"
	BrowserMove BrowserType = "move"
)

type Step struct {
	FromRowFields // Fields to make the step searchable for re-generation

	IntrinsicFields

	AnimationFields

	ModalFields

	ColumnFields

	TerminalFields

	SourceFields

	BrowserFields

	// dev tools
	DevToolsImageName   string `json:"devtoolsImageName"`
	DevToolsImageWidth  int    `json:"devtoolsImageWidth"`
	DevToolsImageHeight int    `json:"devtoolsImageHeight"`

	// markdown
	MarkdownContents            string `json:"markdownContents"`
	MarkdownVerticalAlignment   string `json:"markdownVerticalAlignment"`
	MarkdownHorizontalAlignment string `json:"markdownHorizontalAlignment"`

	// youtube
	YouTubeVideoId string `json:"youtubeVideoId"`
	YouTubeWidth   int    `json:"youtubeWidth"`
	YouTubeHeight  int    `json:"youtubeHeight"`
}

type FromRowFields struct {
	IsFromRow  bool   `json:"isFromRow"`
	SubID      string `json:"subId"`
	ParentStep string `json:"parentStep"`
}

type IntrinsicFields struct {
	StepId  string `json:"stepId"`
	Comment string `json:"comment"`
}

type AnimationFields struct {
	DurationSeconds int  `json:"durationSeconds"`
	IsTrivial       bool `json:"isTrivial"`
}

type ModalFields struct {
	ModalContents string `json:"modalContents"`
}

type ColumnFields struct {
	FocusColumn ColumnType `json:"focusColumn"`
	Column1     ColumnType `json:"column1"`
	Column2     ColumnType `json:"column2"`
	Column3     ColumnType `json:"column3"`
	Column4     ColumnType `json:"column4"`
	Column5     ColumnType `json:"column5"`
	Column6     ColumnType `json:"column6"`
	Column7     ColumnType `json:"column7"`
	Column8     ColumnType `json:"column8"`
	Column9     ColumnType `json:"column9"`
	Column10    ColumnType `json:"column10"`
}

type TerminalTooltip struct {
	TerminalTooltipContents string `json:"terminalTooltipContents"`
	TerminalTooltipTiming   string `json:"terminalTooltipTiming"`
}

type TerminalFields struct {
	CurrentDir       string           `json:"currentDir"`
	TerminalStepType TerminalStepType `json:"terminalType"`
	TerminalText     string           `json:"terminalText"`
	TerminalName     string           `json:"terminalName"`
	TerminalTooltip
}

type SourceTooltip struct {
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
	SourceTooltip
}

type BrowserFields struct {
	BrowserStepType  BrowserType
	BrowserImagePath string `json:"browserImagePath"`
}
