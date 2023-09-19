package rough

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/google/uuid"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

type InnerState struct {
	CurrentSeqNo int
	CurrentCol   string
	ExistingCols []string
}

type RoughStep struct {
	Phase        string `json:"phase"`
	Type         string `json:"type"`
	Instruction  string `json:"instruction"`
	Instruction2 string `json:"instruction2"`
	Instruction3 string `json:"instruction3"`
	Commit       string `json:"commit"`
	Comment      string `json:"comment"`
}

type DetailedStep struct {
	// Uppercase fields to allow json dump for testing

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

func (s *RoughStep) Conversion(state *InnerState, repo *git.Repository) ([]DetailedStep, error) {
	switch s.Type {
	case "terminal":
		return s.TerminalConvert(state, repo)
	case "commit":
		return s.CommitConvert(state, repo)
	case "source error":
		return s.SourceErrorConvert(state, repo)
	case "browser":
		return s.BrowserConvert(state, repo)
	default:
		return nil, fmt.Errorf("unknown type = '%s', phase = '%s', comment = '%s'", s.Type, s.Phase, s.Comment)
	}
}

func fileExists(filename string) (bool, error) {
	_, err := os.Stat(filename)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func findUUID(r *RoughStep, d *DetailedStep, targetFile string) (string, error) {
	// check if targetFile exists
	_, err := os.Stat(targetFile)
	if errors.Is(err, os.ErrNotExist) {
		return uuid.NewString(), nil // if not exists, then new UUID
	}

	// filter by rough step uuid
	var allDetailedSteps []DetailedStep
	err = internal.JsonRead2("rough.json", &allDetailedSteps)
	if err != nil {
		return "", fmt.Errorf("failed to read rough.json, %s", err)
	}

	var filtered []DetailedStep
	for _, s := range allDetailedSteps {
		// if r.Phase == r.Phase && s.Type == r.Type && s.Comment == r.Comment {
		filtered = append(filtered, s)
		// }
	}

	// for each detailed steps
	for _, s := range filtered {
		// if s.DefaultOpenFilePath == targetFile {
		return s.Step, nil
		// }
	}

	// if not found, then new UUID
	return "", nil
}

func (s *RoughStep) CommitConvert(state *InnerState, repo *git.Repository) ([]DetailedStep, error) {
	var detailedSteps []DetailedStep

	// Get info from git
	if s.Commit == "" {
		return nil, fmt.Errorf("commit is missing for manual commit, phase = '%s', type = '%s', comment = '%s'", s.Phase, s.Type, s.Comment)
	}

	files, err := gitFilesForCommit(repo, s.Commit)
	if err != nil {
		return nil, fmt.Errorf("failed to get files for commit = %s, %s", s.Commit, err)
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("failed to get files for commit = %s, no files found", s.Commit)
	}

	// Insert file-tree step if current column != "Source Code"
	if state.CurrentCol != "Source Code" {
		fileTreeStep := DetailedStep{
			FocusColumn:         "Source Code",
			IsFoldFileTree:      false,
			DefaultOpenFilePath: files[0],
			Commit:              s.Commit,
		}
		// uuid = findUUID(s, &fileTreeStep, files[0])
		detailedSteps = append(detailedSteps, fileTreeStep)
	}

	// file steps
	for i, file := range files {
		commitStep := DetailedStep{
			FocusColumn:         "Source Code",
			DefaultOpenFilePath: file,
			IsFoldFileTree:      true,
		}
		// uuid = findUUID(s, &fileTreeStep, files[0])
		detailedSteps = append(detailedSteps, commitStep)

		if i == 5 {
			break
		}
	}

	return detailedSteps, nil
}

func (s *RoughStep) TerminalConvert(state *InnerState, repo *git.Repository) ([]DetailedStep, error) {
	var detailedSteps []DetailedStep

	// check if it's a valid terminal step
	if s.Instruction == "" && s.Instruction2 == "" {
		return nil, fmt.Errorf("step is missing both 'instruction' and 'instruction2', phase = '%s', type = '%s', comment = '%s'", s.Phase, s.Type, s.Comment)
	}

	// insert move-to-terminal step if current column != "Terminal"
	if state.CurrentCol != "Terminal" {
		fileTreeStep := DetailedStep{
			FocusColumn: "Terminal",
			Comment:     "(move)",
		}
		detailedSteps = append(detailedSteps, fileTreeStep)
	}

	// command step
	// * check if it's a 'cd' command
	var currentDir string
	if strings.HasPrefix(s.Instruction, "cd ") {
		currentDir = strings.TrimPrefix(s.Instruction, "cd ")
	}
	// * create command step
	cmdStep := DetailedStep{
		FocusColumn:  "Terminal",
		TerminalType: "command",
		TerminalText: s.Instruction,
		TerminalName: s.Instruction3, // Go zero value is ""
		CurrentDir:   currentDir,     // Go zero value is ""
		Commit:       s.Commit,       // Go zero value is ""
	}
	detailedSteps = append(detailedSteps, cmdStep)

	// output step
	if s.Instruction2 != "" {
		outputStep := DetailedStep{
			FocusColumn:  "Terminal",
			TerminalType: "output",
			TerminalText: s.Instruction2,
		}
		detailedSteps = append(detailedSteps, outputStep)
	}

	// Udpate the state
	state.CurrentCol = "Terminal"

	// source code steps
	if s.Commit != "" {
		commitSteps, err := s.CommitConvert(state, repo)
		if err != nil {
			return nil, fmt.Errorf("failed to convert commit steps, %s", err)
		}
		detailedSteps = append(detailedSteps, commitSteps...)
	}

	return detailedSteps, nil
}

func (s *RoughStep) SourceErrorConvert(state *InnerState, repo *git.Repository) ([]DetailedStep, error) {
	var detailedSteps []DetailedStep

	// 1. source code step
	sourceErrorStep := DetailedStep{
		FocusColumn:         "Source Code",
		DefaultOpenFilePath: s.Instruction, // Go zero value is ""
	}

	detailedSteps = append(detailedSteps, sourceErrorStep)

	// 2. udpate the state
	state.CurrentCol = "Source Code"

	return detailedSteps, nil
}

func (s *RoughStep) BrowserConvert(state *InnerState, repo *git.Repository) ([]DetailedStep, error) {
	var detailedSteps []DetailedStep

	// Browser step
	if s.Instruction == "" {
		// no instruction - single browser step
		browserStep := DetailedStep{
			FocusColumn: "Browser",
		}
		detailedSteps = append(detailedSteps, browserStep)
	} else {
		// no instruction - multiple browser steps
		split := strings.Split(s.Instruction, ",")
		for _, s := range split {
			browserImageName := strings.ReplaceAll(s, " ", "")
			browserStep := DetailedStep{
				FocusColumn:      "Browser",
				BrowserImageName: browserImageName,
			}
			detailedSteps = append(detailedSteps, browserStep)
		}
	}

	// 2. udpate the state
	state.CurrentCol = "Browser"

	return detailedSteps, nil
}
