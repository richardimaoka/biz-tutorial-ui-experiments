package result

type ColumnType string

const (
	// Lower cases since they are from manual entries
	NoColumn       ColumnType = ""
	SourceColumn   ColumnType = "SourceCode"
	TerminalColumn ColumnType = "Terminal"
	BrowserColumn  ColumnType = "Browser"
)

type TerminalType string

const (
	// Lower cases since they are from manual entries
	TerminalCommand TerminalType = "command"
	TerminalOutput  TerminalType = "output"
	TerminalCd      TerminalType = "cd"
	TerminalMove    TerminalType = "move"
	TerminalOpen    TerminalType = "open"
)

type SourceCodeType string

const (
	// Lower cases since they are from manual entries
	FileTree    SourceCodeType = "fileTree"
	SourceOpen  SourceCodeType = "openFile"
	SourceError SourceCodeType = "error"
	SourceMove  SourceCodeType = "move"
)

type BrowserType string

const (
	// Lower cases since they are from manual entries
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

	SourceCodeFields

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

type TerminalTooltipFields struct {
	TerminalTooltipContents string `json:"terminalTooltipContents"`
	TerminalTooltipTiming   string `json:"terminalTooltipTiming"`
}

type TerminalFields struct {
	CurrentDir       string       `json:"currentDir"`
	TerminalStepType TerminalType `json:"terminalType"`
	TerminalText     string       `json:"terminalText"`
	TerminalName     string       `json:"terminalName"`
	TerminalTooltipFields
}

type SourceCodeTooltip struct {
	SourceCodeTooltipContents   string `json:"sourceCodeTooltipContents"`
	SourceCodeTooltipTiming     string `json:"sourceCodeTooltipTiming"`
	SourceCodeTooltipLineNumber int    `json:"sourceCodeTooltipLineNumber"`
	SourceCodeTooltipIsAppend   bool   `json:"SourceCodeTooltipIsAppend"`
}

type SourceCodeFields struct {
	Commit              string `json:"commit"`
	DefaultOpenFilePath string `json:"defaultOpenFilePath"`
	ShowFileTree        bool   `json:"showFileTree"`
	TypingAnimation     bool   `json:"typingAnimation"`
	SourceCodeTooltip
}

type BrowserFields struct {
	BrowserStepType  BrowserType
	BrowserImagePath string `json:"browserImagePath"`
}
