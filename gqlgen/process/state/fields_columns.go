package state

/**
 * Column fields
 */

type ColumnType string

const (
	// `Type` suffix is needed to avoid conflict with structs
	NoColumnType       ColumnType = ""
	SourceColumnType   ColumnType = "SourceCode"
	TerminalColumnType ColumnType = "Terminal"
	BrowserColumnType  ColumnType = "Browser"
)

type ColumnFields struct {
	FocusColumn ColumnType `json:"focusColumn"`
}

/**
 * Source column fields
 */

type SourceStepType string

const (
	FileTree     SourceStepType = "fileTree"
	SourceOpen   SourceStepType = "openFile"
	SourceCommit SourceStepType = "sourceCommit"
	SourceError  SourceStepType = "error"
	SourceMove   SourceStepType = "move"
)

type SourceTooltipFields struct {
	SourceTooltipContents   string                  `json:"sourceTooltipContents"`
	SourceTooltipTiming     SourceCodeTooltipTiming `json:"sourceTooltipTiming"`
	SourceTooltipLineNumber int                     `json:"sourceTooltipLineNumber"`
	SourceTooltipIsAppend   bool                    `json:"SourceTooltipIsAppend"`
}

type SourceFields struct {
	SourceStepType      SourceStepType `json:"sourceStepType"`
	DefaultOpenFilePath string         `json:"defaultOpenFilePath"`
	Commit              string         `json:"commit"`
	TypingAnimation     bool           `json:"typingAnimation"`
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
	TerminalTooltipContents string                `json:"terminalTooltipContents"`
	TerminalTooltipTiming   TerminalTooltipTiming `json:"terminalTooltipTiming"`
}

type TerminalFields struct {
	CurrentDir       string           `json:"currentDir"`
	TerminalStepType TerminalStepType `json:"terminalStepType"`
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
 * YouTube column fields
 */
type YoutubeFields struct {
	YouTubeVideoId string `json:"youtubeVideoId"`
	YouTubeWidth   int    `json:"youtubeWidth"`
	YouTubeHeight  int    `json:"youtubeHeight"`
}
