package input

type ColumnType = string

const (
	// Lower cases since they are from manual entries
	SourceType   ColumnType = "source"
	TerminalType ColumnType = "terminal"
	BrowserType  ColumnType = "browser"
)

type SubType = string

const (
	// Lower cases since they are from manual entries
	CommandSubType SubType = "command"
	OutputSubType  SubType = "output"
)

type Row struct {
	StepId        string `json:"stepId"`
	Phase         string `json:"phase"`
	Comment       string `json:"comment"`
	Column        string `json:"column"` //not Column but string, because it's input from manual entry, not sanitized
	Type          string `json:"type"`
	Trivial       string `json:"trivial"`
	Instruction   string `json:"instruction"`
	Instruction2  string `json:"instruction2"`
	Instruction3  string `json:"instruction3"`
	ModalText     string `json:"modalText"`
	Tooltip       string `json:"tooltip"`
	TooltipTiming string `json:"tooltipTiming"`
	TooltipLine   int    `json:"tooltipLine"`
	// TooltipPosition string `json:"tooltipPosition"`
}

type ResultStep struct {
	// Uppercase fields to allow json dump for testing

	// rough-step related
	IsFromRow  bool   `json:"isFromRow"`
	SubID      string `json:"subId"`
	ParentStep string `json:"parentStep"`

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
	DurationSeconds int  `json:"durationSeconds"`
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
