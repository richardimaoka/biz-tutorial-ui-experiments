package rough

type DetailedStepTest struct {
	// Uppercase fields to allow json dump for testing

	// internal fields
	FromRoughStep bool   `json:"fromRoughStep,omitempty"`
	SubID         string `json:"subId,omitempty"`

	// steps
	ParentStep      string `json:"parentStep,omitempty"`
	Step            string `json:"step"`
	AutoNextSeconds int    `json:"autoNextSeconds,omitempty"`
	DurationSeconds int    `json:"duration,omitempty"`
	IsTrivialStep   bool   `json:"isTrivialStep,omitempty"`
	Comment         string `json:"comment,omitempty"`

	// columns
	FocusColumn string `json:"focusColumn,omitempty"`
	Column1     string `json:"column1,omitempty"`
	Column2     string `json:"column2,omitempty"`
	Column3     string `json:"column3,omitempty"`
	Column4     string `json:"column4,omitempty"`
	Column5     string `json:"column5,omitempty"`

	// modal
	ModalText     string `json:"modalText,omitempty"`
	ModalPosition string `json:"modalPosition,omitempty"`

	// terminal
	TerminalText string `json:"terminalText,omitempty"`
	TerminalType string `json:"terminalType,omitempty"`
	TerminalName string `json:"terminalName,omitempty"`
	CurrentDir   string `json:"currentDir,omitempty"`

	// git
	Commit              string `json:"commit,omitempty"`
	CommitMessage       string `json:"commitMessage,omitempty"`
	PrevCommit          string `json:"prevCommit,omitempty"`
	RepoUrl             string `json:"repoUrl,omitempty"`
	DefaultOpenFilePath string `json:"defaultOpenFilePath,omitempty"`
	IsFoldFileTree      bool   `json:"isFoldFileTree,omitempty"`

	// browser
	BrowserImageName   string `json:"browserImageName,omitempty"`
	BrowserImageWidth  int    `json:"browserImageWidth,omitempty"`
	BrowserImageHeight int    `json:"browserImageHeight,omitempty"`

	// dev tools
	DevToolsImageName   string `json:"devtoolsImageName,omitempty"`
	DevToolsImageWidth  int    `json:"devtoolsImageWidth,omitempty"`
	DevToolsImageHeight int    `json:"devtoolsImageHeight,omitempty"`

	// markdown
	MarkdownContents            string `json:"markdownContents,omitempty"`
	MarkdownVerticalAlignment   string `json:"markdownVerticalAlignment,omitempty"`
	MarkdownHorizontalAlignment string `json:"markdownHorizontalAlignment,omitempty"`

	// youtube
	YouTubeVideoId string `json:"youtubeVideoId,omitempty"`
	YouTubeWidth   int    `json:"youtubeWidth,omitempty"`
	YouTubeHeight  int    `json:"youtubeHeight,omitempty"`
}
