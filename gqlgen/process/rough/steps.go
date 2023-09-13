package rough

import (
	"fmt"
	"strings"
)

type InnerState struct {
	currentSeqNo int
	currentCol   string
	existingCols []string
}

type RoughStep struct {
	// Phase       string `json:"phase"`
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

func (step *DetailedStep) setColumns(existingColumns []string, focusColumn string) {
	var focusColumnExists bool

	if len(existingColumns) > 0 {
		step.Column1 = existingColumns[0]
		focusColumnExists = existingColumns[0] == focusColumn
	} else {
		step.Column1 = focusColumn
		return
	}

	if len(existingColumns) > 1 {
		step.Column2 = existingColumns[1]
		focusColumnExists = existingColumns[1] == focusColumn
	} else {
		if !focusColumnExists {
			step.Column2 = focusColumn
		}
		return
	}

	if len(existingColumns) > 2 {
		step.Column3 = existingColumns[2]
		focusColumnExists = existingColumns[2] == focusColumn
	} else {
		if !focusColumnExists {
			step.Column3 = focusColumn
		}
		return
	}

	if len(existingColumns) > 3 {
		step.Column4 = existingColumns[3]
		focusColumnExists = existingColumns[3] == focusColumn
	} else {
		if !focusColumnExists {
			step.Column4 = focusColumn
		}
		return
	}

	if len(existingColumns) > 4 {
		step.Column5 = existingColumns[4]
		focusColumnExists = existingColumns[4] == focusColumn
	} else {
		if !focusColumnExists {
			step.Column5 = focusColumn
		}
		return
	}
}

func command(uuid, command, commit string, existingColumns []string) DetailedStep {
	step := DetailedStep{
		Step:         uuid,
		FocusColumn:  "Terminal",
		TerminalText: command,
		TerminalType: "command",
		Commit:       commit,
	}
	step.setColumns(existingColumns, "Terminal")
	return step
}

func isCommand(instruction string) bool {
	return strings.HasPrefix(instruction, "mkdir ") ||
		strings.HasPrefix(instruction, "cd ") ||
		strings.HasPrefix(instruction, "go ")
}

func isManualCommit(instruction string) bool {
	return instruction == "# manual commit"
}

func isAutoCommit(instruction string) bool {
	return instruction == "# auto commit"
}

func isSourceError(instruction string) bool {
	return instruction == "# source error"
}

func isBrowser(instruction string) bool {
	return instruction == "# browser"
}

func (s *RoughStep) Convert(uuid string, columns []string) ([]DetailedStep, error) {
	if isCommand(s.Instruction) {
		ds := command(uuid, s.Instruction, s.Commit, columns)
		return []DetailedStep{ds}, nil
	} else if isManualCommit(s.Instruction) {
	} else if isAutoCommit(s.Instruction) {
	} else if isSourceError(s.Instruction) {
	} else if isBrowser(s.Instruction) {
	}

	return nil, fmt.Errorf("unhandled case")
}
