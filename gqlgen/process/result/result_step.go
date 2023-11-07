package result

type ColumnType string

const (
	// Lower cases since they are from manual entries
	NoColumn       ColumnType = ""
	SourceColumn   ColumnType = "source"
	TerminalColumn ColumnType = "terminal"
	BrowserColumn  ColumnType = "browser"
)

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

type TerminalType string

const (
	// Lower cases since they are from manual entries
	TerminalCommand TerminalType = "command"
	TerminalOutput  TerminalType = "output"
	TerminalCd      TerminalType = "cd"
	TerminalMove    TerminalType = "move"
)

type TerminalTooltip struct {
	TerminalTooltipContents string `json:"terminalTooltipContents"`
	TerminalTooltipTiming   string `json:"terminalTooltipTiming"`
}

type TerminalFields struct {
	CurrentDir   string       `json:"currentDir"`
	TerminalType TerminalType `json:"terminalType"`
	TerminalText string       `json:"terminalText"`
	TerminalName string       `json:"terminalName"`
	TerminalTooltip
}

type Step struct {
	// Uppercase fields to allow json dump for testing

	// Fields to make the step searchable for re-generation
	IsFromRow  bool   `json:"isFromRow"`
	SubID      string `json:"subId"`
	ParentStep string `json:"parentStep"`

	// steps
	StepId  string `json:"stepId"`
	Comment string `json:"comment"`

	// animation
	DurationSeconds int  `json:"durationSeconds"`
	IsTrivial       bool `json:"isTrivial"`

	// modal
	ModalContents string `json:"modalContents"`

	// columns
	ColumnFields

	// terminal
	TerminalFields

	// git
	Commit              string `json:"commit"`
	DefaultOpenFilePath string `json:"defaultOpenFilePath"`
	IsFoldFileTree      bool   `json:"isFoldFileTree"`

	// browser
	BrowserImageName   string `json:"browserImageName"`
	BrowserImageWidth  int    `json:"browserImageWidth"`
	BrowserImageHeight int    `json:"browserImageHeight"`

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
