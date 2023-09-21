package rough

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
)

type RoughStep struct {
	Step         string `json:"step"`
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

func (state *InnerState) fileTreeStep(s *RoughStep, file string) DetailedStep {
	subId := "fileTreeStep"
	stepId := state.uuidFinder.FindOrGenerateUUID(s, subId)
	fileTreeStep := DetailedStep{
		// fields to make the step searchable for re-generation
		FromRoughStep: true,
		ParentStep:    s.Step,
		SubID:         subId,
		// other fields
		Step:                stepId,
		FocusColumn:         "Source Code",
		IsFoldFileTree:      false,
		DefaultOpenFilePath: file,
		Commit:              s.Commit,
	}

	return fileTreeStep
}

func (state *InnerState) openFileStep(s *RoughStep, index int, file string) DetailedStep {
	subId := fmt.Sprintf("openFileStep-%d", index)
	stepId := state.uuidFinder.FindOrGenerateUUID(s, subId)
	fileTreeStep := DetailedStep{
		// fields to make the step searchable for re-generation
		FromRoughStep: true,
		ParentStep:    s.Step,
		SubID:         subId,
		// other fields
		Step:                stepId,
		FocusColumn:         "Source Code",
		DefaultOpenFilePath: file,
		IsFoldFileTree:      true,
	}

	return fileTreeStep
}

func (state *InnerState) moveToTerminalStep(s *RoughStep) DetailedStep {
	subId := "moveToTerminalStep"
	stepId := state.uuidFinder.FindOrGenerateUUID(s, subId)
	step := DetailedStep{
		// fields to make the step searchable for re-generation
		FromRoughStep: true,
		ParentStep:    s.Step,
		SubID:         subId,
		// other fields
		Step:        stepId,
		FocusColumn: "Terminal",
		Comment:     "(move)",
	}
	return step
}

func (state *InnerState) terminalOutputStep(s *RoughStep) DetailedStep {
	subId := "terminalOutputStep"
	stepId := state.uuidFinder.FindOrGenerateUUID(s, subId)
	step := DetailedStep{
		// fields to make the step searchable for re-generation
		FromRoughStep: true,
		ParentStep:    s.Step,
		SubID:         subId,
		// other fields
		Step:         stepId,
		FocusColumn:  "Terminal",
		TerminalType: "output",
		TerminalText: s.Instruction2,
	}

	return step
}

func (state *InnerState) sourceErrorStep(s *RoughStep) DetailedStep {
	subId := "sourceErrorStep"
	stepId := state.uuidFinder.FindOrGenerateUUID(s, subId)
	step := DetailedStep{
		// fields to make the step searchable for re-generation
		FromRoughStep: true,
		ParentStep:    s.Step,
		SubID:         subId,
		// other fields
		Step:                stepId,
		FocusColumn:         "Source Code",
		DefaultOpenFilePath: s.Instruction, // Go zero value is ""
	}

	return step
}

func (state *InnerState) browserStep(s *RoughStep, index int, browserImageName string) DetailedStep {
	subId := fmt.Sprintf("browserStep-%d", index)
	stepId := state.uuidFinder.FindOrGenerateUUID(s, subId)
	step := DetailedStep{
		// fields to make the step searchable for re-generation
		FromRoughStep: true,
		ParentStep:    s.Step,
		SubID:         subId,
		// other fields
		Step:             stepId,
		FocusColumn:      "Browser",
		BrowserImageName: browserImageName,
	}

	return step
}

func (state *InnerState) terminalCommandStep(s *RoughStep) DetailedStep {
	// * check if it's a 'cd' command
	var currentDir string
	if strings.HasPrefix(s.Instruction, "cd ") {
		currentDir = strings.TrimPrefix(s.Instruction, "cd ")
	}

	subId := "terminalCommandStep"
	stepId := state.uuidFinder.FindOrGenerateUUID(s, subId)

	step := DetailedStep{
		// fields to make the step searchable for re-generation
		FromRoughStep: true,
		ParentStep:    s.Step,
		SubID:         subId,
		// other fields
		Step:         stepId,
		FocusColumn:  "Terminal",
		TerminalType: "command",
		TerminalText: s.Instruction,
		TerminalName: s.Instruction3, // Go zero value is ""
		CurrentDir:   currentDir,     // Go zero value is ""
		Commit:       s.Commit,       // Go zero value is ""
	}

	return step
}

func (state *InnerState) commitConvert(s *RoughStep, repo *git.Repository) ([]DetailedStep, error) {
	var detailedSteps []DetailedStep

	// Get info from git
	if s.Commit == "" {
		return nil, fmt.Errorf("commit is missing for manual commit, phase = '%s', type = '%s', comment = '%s'", s.Phase, s.Type, s.Comment)
	}

	// find files from commit
	files, err := gitFilesForCommit(repo, s.Commit)
	if err != nil {
		return nil, fmt.Errorf("failed to get files for commit = %s, %s", s.Commit, err)
	}

	// Insert file-tree step if current column != "Source Code"
	if state.currentColumn != "Source Code" {
		fileTreeStep := state.fileTreeStep(s, files[0])
		detailedSteps = append(detailedSteps, fileTreeStep)
	}

	// file steps
	for i, file := range files {
		openFileStep := state.openFileStep(s, i, file)
		detailedSteps = append(detailedSteps, openFileStep)
		if i == 5 {
			break
		}
	}

	return detailedSteps, nil
}

func (state *InnerState) terminalConvert(s *RoughStep, repo *git.Repository) ([]DetailedStep, error) {
	var detailedSteps []DetailedStep

	// check if it's a valid terminal step
	if s.Instruction == "" && s.Instruction2 == "" {
		return nil, fmt.Errorf("step is missing both 'instruction' and 'instruction2', phase = '%s', type = '%s', comment = '%s'", s.Phase, s.Type, s.Comment)
	}

	// insert move-to-terminal step if current column != "Terminal"
	if state.currentColumn != "Terminal" {
		moveToTerminalStep := state.moveToTerminalStep(s)
		detailedSteps = append(detailedSteps, moveToTerminalStep)
	}

	// command step
	cmdStep := state.terminalCommandStep(s)
	detailedSteps = append(detailedSteps, cmdStep)

	// output step
	if s.Instruction2 != "" {
		outputStep := state.terminalOutputStep(s)
		detailedSteps = append(detailedSteps, outputStep)
	}

	// Udpate the state
	state.currentColumn = "Terminal"

	// source code steps
	if s.Commit != "" {
		commitSteps, err := state.commitConvert(s, repo)
		if err != nil {
			return nil, fmt.Errorf("failed to convert commit steps, %s", err)
		}
		detailedSteps = append(detailedSteps, commitSteps...)
	}

	return detailedSteps, nil
}

func (state *InnerState) sourceErrorConvert(s *RoughStep, repo *git.Repository) ([]DetailedStep, error) {
	var detailedSteps []DetailedStep

	// source code step
	sourceErrorStep := state.sourceErrorStep(s)
	detailedSteps = append(detailedSteps, sourceErrorStep)

	// udpate the state
	state.currentColumn = "Source Code"

	return detailedSteps, nil
}

func (state *InnerState) browserConvert(s *RoughStep, repo *git.Repository) ([]DetailedStep, error) {
	var detailedSteps []DetailedStep

	if s.Instruction == "" {
		return nil, fmt.Errorf("instruction is missing for browser step, phase = '%s', type = '%s', comment = '%s'", s.Phase, s.Type, s.Comment)
	}

	// browser steps
	split := strings.Split(s.Instruction, ",")
	for i, each := range split {
		browserImageName := strings.ReplaceAll(each, " ", "")
		browserStep := state.browserStep(s, i, browserImageName)
		detailedSteps = append(detailedSteps, browserStep)
	}

	// 2. udpate the state
	state.currentColumn = "Browser"

	return detailedSteps, nil
}

func (state *InnerState) Conversion(s *RoughStep, repo *git.Repository) ([]DetailedStep, error) {
	switch s.Type {
	case "terminal":
		return state.terminalConvert(s, repo)
	case "commit":
		return state.commitConvert(s, repo)
	case "source error":
		return state.sourceErrorConvert(s, repo)
	case "browser":
		return state.browserConvert(s, repo)
	default:
		return nil, fmt.Errorf("unknown type = '%s', phase = '%s', comment = '%s'", s.Type, s.Phase, s.Comment)
	}
}
