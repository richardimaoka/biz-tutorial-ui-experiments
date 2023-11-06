package input

type UsedColumns = [10]string
type CurrentColumn = string

type ResultStep struct {
	// Uppercase fields to allow json dump for testing

	// Fields to make the step searchable for re-generation
	IsFromRow  bool   `json:"isFromRow"`
	SubID      string `json:"subId"`
	ParentStep string `json:"parentStep"`

	// steps
	Step    string `json:"step"`
	Comment string `json:"comment"`

	// columns
	FocusColumn string `json:"focusColumn"`
	Column1     string `json:"column1"`
	Column2     string `json:"column2"`
	Column3     string `json:"column3"`
	Column4     string `json:"column4"`
	Column5     string `json:"column5"`
	Column6     string `json:"column6"`
	Column7     string `json:"column7"`
	Column8     string `json:"column8"`
	Column9     string `json:"column9"`
	Column10    string `json:"column10"`

	// animation
	DurationSeconds int  `json:"durationSeconds"`
	IsTrivial       bool `json:"isTrivial"`

	// modal
	ModalContents string `json:"modalContents"`
	ModalPosition string `json:"modalPosition"`

	// terminal
	CurrentDir              string `json:"currentDir"`
	TerminalType            string `json:"terminalType"`
	TerminalText            string `json:"terminalText"`
	TerminalName            string `json:"terminalName"`
	TerminalTooltipContents string `json:"terminalTooltipContents"`
	TerminalTooltipTiming   string `json:"terminalTooltipTiming"`

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

func (step *ResultStep) setColumns(cols UsedColumns) {
	step.Column1 = cols[0]
	step.Column2 = cols[1]
	step.Column3 = cols[2]
	step.Column4 = cols[3]
	step.Column5 = cols[4]
	step.Column6 = cols[5]
	step.Column7 = cols[6]
	step.Column8 = cols[7]
	step.Column9 = cols[8]
	step.Column10 = cols[9]
}
