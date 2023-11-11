package state

/**
 * Column fields
 */

type ColumnType string

const (
	NoColumnType       ColumnType = ""
	SourceColumnType   ColumnType = "SourceCode"
	TerminalColumnType ColumnType = "Terminal"
	BrowserColumnType  ColumnType = "Browser"
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

/**
 * Source column fields
 */

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

/**
 * Termianl column fields
 */
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

/**
 * Browser column fields
 */
type BrowserStepType string

const (
	BrowserOpen BrowserStepType = "open"
	BrowserMove BrowserStepType = "move"
)

type BrowserFields struct {
	BrowserStepType  BrowserStepType
	BrowserImagePath string `json:"browserImagePath"`
}

/**
 * Browser DevTools column fields
 */
type BrowserDevToolsFields struct {
	DevToolsImageName   string `json:"devtoolsImageName"`
	DevToolsImageWidth  int    `json:"devtoolsImageWidth"`
	DevToolsImageHeight int    `json:"devtoolsImageHeight"`
}

/**
 * Markdown column fields
 */
type MarkdownFields struct {
	MarkdownContents            string `json:"markdownContents"`
	MarkdownVerticalAlignment   string `json:"markdownVerticalAlignment"`
	MarkdownHorizontalAlignment string `json:"markdownHorizontalAlignment"`
}

/**
 * YouTube column fields
 */
type YoutubeFields struct {
	YouTubeVideoId string `json:"youtubeVideoId"`
	YouTubeWidth   int    `json:"youtubeWidth"`
	YouTubeHeight  int    `json:"youtubeHeight"`
}
