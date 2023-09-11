package rough

import "github.com/google/uuid"

type InnerState struct {
	currentSeqNo int
	currentCol   string
	existingCols []string
}

type RoughStep struct {
	Phase       string `json:"phase"`
	Type        string `json:"type"`
	Instruction string `json:"instruction"`
	Commit      string `json:"commit"`
	Comment     string `json:"comment"`
}

type DetailedStep struct {
	// Uppercase fields to allow json dump for testing

	// steps
	Step            string `json:"step"`
	AutoNextSeconds int    `json:"autoNextSeconds,omitempty"`
	DurationSeconds int    `json:"duration,omitempty"`
	IsTrivialStep   bool   `json:"isTrivialStep,omitempty"`

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
	CurrentDir   string `json:"currentDir,omitempty"`

	// git
	Commit              string `json:"commit,omitempty"`
	CommitMessage       string `json:"commitMessage,omitempty"`
	PrevCommit          string `json:"prevCommit,omitempty"`
	RepoUrl             string `json:"repoUrl,omitempty"`
	DefaultOpenFilePath string `json:"defaultOpenFilePath,omitempty"`
	IsFoldFileTree      string `json:"isFoldFileTree,omitempty"` // string, as CSV from Google Spreadsheet has TRUE as upper-case 'TRUE'

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

func simpleCommand(command string) DetailedStep {
	uuid := uuid.NewString()
	return DetailedStep{
		Step:         uuid,
		TerminalText: command,
		TerminalType: "command",
	}
}

func GenDetailedSteps(filename string) []DetailedStep {
	var detailedSteps []DetailedStep

	detailedSteps = append(detailedSteps, DetailedStep{
		FocusColumn:  "Terminal",
		Column1:      "Terminal",
		Step:         "c8238063-dd5a-4cd4-9d62-5c9c8ebd35ec",
		TerminalText: "mkdir gqlgen-todos",
	})

	return detailedSteps
}
