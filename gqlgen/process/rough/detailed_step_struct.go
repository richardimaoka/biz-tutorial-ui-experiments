package rough

type DetailedStep struct {
	// Uppercase fields to allow json dump for testing

	// internal fields
	FromRoughStep bool   `json:"fromRoughStep"`
	SubID         string `json:"subId"`

	// steps
	ParentStep      string `json:"parentStep"`
	Step            string `json:"step"`
	AutoNextSeconds int    `json:"autoNextSeconds"`
	DurationSeconds int    `json:"duration"`
	IsTrivialStep   bool   `json:"isTrivialStep"`
	Comment         string `json:"comment"`

	// columns
	FocusColumn string `json:"focusColumn"`
	Column1     string `json:"column1"`
	Column2     string `json:"column2"`
	Column3     string `json:"column3"`
	Column4     string `json:"column4"`
	Column5     string `json:"column5"`

	// modal
	ModalText     string `json:"modalText"`
	ModalPosition string `json:"modalPosition"`

	// terminal
	TerminalText string `json:"terminalText"`
	TerminalType string `json:"terminalType"`
	TerminalName string `json:"terminalName"`
	CurrentDir   string `json:"currentDir"`

	// git
	Commit              string `json:"commit"`
	CommitMessage       string `json:"commitMessage"`
	PrevCommit          string `json:"prevCommit"`
	RepoUrl             string `json:"repoUrl"`
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
