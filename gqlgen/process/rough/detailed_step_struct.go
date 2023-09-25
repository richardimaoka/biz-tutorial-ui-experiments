package rough

type DetailedStep struct {
	// Uppercase fields to allow json dump for testing

	// rough-step related
	FromRoughStep bool   `json:"fromRoughStep"`
	SubID         string `json:"subId"`
	ParentStep    string `json:"parentStep"`

	// steps
	Step string `json:"step"`

	// columns
	FocusColumn string `json:"focusColumn"`
	Column1     string `json:"column1"`
	Column2     string `json:"column2"`
	Column3     string `json:"column3"`
	Column4     string `json:"column4"`
	Column5     string `json:"column5"`

	Comment string `json:"comment"`

	// animation
	DurationSeconds int  `json:"duration"`
	IsTrivialStep   bool `json:"isTrivialStep"`

	// modal
	ModalText     string `json:"modalText"`
	ModalPosition string `json:"modalPosition"`

	// terminal
	TerminalType string `json:"terminalType"`
	TerminalText string `json:"terminalText"`
	CurrentDir   string `json:"currentDir"`
	TerminalName string `json:"terminalName"`

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
