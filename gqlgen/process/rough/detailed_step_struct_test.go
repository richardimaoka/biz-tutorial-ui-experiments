package rough

type DetailedStepTest struct {
	// Uppercase fields to allow json dump for testing

	// rough-step related
	FromRoughStep bool   `json:"fromRoughStep,omitempty"`
	SubID         string `json:"subId,omitempty"`
	ParentStep    string `json:"parentStep,omitempty"`

	// steps
	Step string `json:"step"`

	// columns
	FocusColumn string `json:"focusColumn,omitempty"`
	Column1     string `json:"column1,omitempty"`
	Column2     string `json:"column2,omitempty"`
	Column3     string `json:"column3,omitempty"`
	Column4     string `json:"column4,omitempty"`
	Column5     string `json:"column5,omitempty"`

	Comment string `json:"comment,omitempty"`

	// animation
	DurationSeconds int  `json:"duration,omitempty"`
	IsTrivialStep   bool `json:"isTrivialStep,omitempty"`

	// modal
	ModalText     string `json:"modalText,omitempty"`
	ModalPosition string `json:"modalPosition,omitempty"`

	// terminal
	TerminalType string `json:"terminalType,omitempty"`
	TerminalText string `json:"terminalText,omitempty"`
	CurrentDir   string `json:"currentDir,omitempty"`
	TerminalName string `json:"terminalName,omitempty"`

	// git
	Commit              string `json:"commit,omitempty"`
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
