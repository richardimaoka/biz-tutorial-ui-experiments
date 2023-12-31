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
	SourceTooltipContents   string                  `json:"sourceTooltipContents,omitempty"`
	SourceTooltipTiming     SourceCodeTooltipTiming `json:"sourceTooltipTiming,omitempty"`
	SourceTooltipLineNumber int                     `json:"sourceTooltipLineNumber,omitempty"`
	SourceTooltipIsAppend   bool                    `json:"SourceTooltipIsAppend,omitempty"`
}

type SourceFields struct {
	SourceStepType      SourceStepType `json:"sourceStepType,omitempty"`
	DefaultOpenFilePath string         `json:"defaultOpenFilePath,omitempty"`
	Commit              string         `json:"commit,omitempty"`
	TypingAnimation     bool           `json:"typingAnimation,omitempty"`
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
	TerminalTooltipContents string                `json:"terminalTooltipContents,omitempty"`
	TerminalTooltipTiming   TerminalTooltipTiming `json:"terminalTooltipTiming,omitempty"`
}

type TerminalFields struct {
	CurrentDir       string           `json:"currentDir,omitempty"`
	TerminalStepType TerminalStepType `json:"terminalStepType,omitempty"`
	TerminalText     string           `json:"terminalText,omitempty"`
	TerminalName     string           `json:"terminalName,omitempty"`
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
	BrowserStepType    BrowserStepType
	BrowserImagePath   string `json:"browserImagePath,omitempty"`
	BrowserImageWidth  int    `json:"browserImageWidth,omitempty"`
	BrowserImageHeight int    `json:"browserImageHeight,omitempty"`
}

/**
 * Browser DevTools column fields
 */
type BrowserDevToolsFields struct {
	DevToolsImageName   string `json:"devtoolsImageName,omitempty"`
	DevToolsImageWidth  int    `json:"devtoolsImageWidth,omitempty"`
	DevToolsImageHeight int    `json:"devtoolsImageHeight,omitempty"`
}

/**
 * YouTube column fields
 */
type YoutubeFields struct {
	YouTubeVideoId string `json:"youtubeVideoId,omitempty"`
	YouTubeWidth   int    `json:"youtubeWidth,omitempty"`
	YouTubeHeight  int    `json:"youtubeHeight,omitempty"`
}
