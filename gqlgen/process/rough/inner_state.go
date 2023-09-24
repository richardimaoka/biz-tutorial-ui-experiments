package rough

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

type InnerState struct {
	repo          *git.Repository
	currentSeqNo  int
	currentColumn string
	existingCols  [5]string //fixed size, according to DetailedStep
	uuidFinder    *UUIDFinder
	prevCommit    plumbing.Hash
}

func NewInnerState(targetFile string, repo *git.Repository) (*InnerState, error) {
	finder, err := NewUUIDFinder(targetFile)
	if err != nil {
		return nil, fmt.Errorf("failed to create UUIDFinder: %s", err)
	}

	return &InnerState{
		repo:       repo,
		uuidFinder: finder,
	}, nil
}

//////////////////////////////////////////////////////
// Overall conversion methods
//////////////////////////////////////////////////////

func Process(dir, repoUrl string) error {
	roughFile := fmt.Sprintf("%s/rough-steps.json", dir)
	targetFile := fmt.Sprintf("%s/detailed-steps.json", dir)

	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: repoUrl})
	if err != nil {
		return fmt.Errorf("cannot clone repo %s, %s", repoUrl, err)
	}

	state, err := NewInnerState(targetFile, repo)
	if err != nil {
		return fmt.Errorf("failed to create InnerState: %s", err)
	}

	detailedSteps, err := state.generateTarget(roughFile)
	if err != nil {
		return fmt.Errorf("failed to generate detailed steps, %s", err)
	}

	err = internal.WriteJsonToFile(detailedSteps, targetFile)
	if err != nil {
		return fmt.Errorf("failed to write detailed steps to file, %s", err)
	}

	return nil
}

func (state *InnerState) generateTarget(roughStepsFile string) ([]DetailedStep, error) {
	var roughSteps []RoughStep
	err := internal.JsonRead2(roughStepsFile, &roughSteps)
	if err != nil {
		return nil, fmt.Errorf("GenerateTarget error - failed to read from json: %v", err)
	}

	var detailedSteps []DetailedStep
	for _, s := range roughSteps {
		dSteps, err := state.Conversion(&s, state.repo)
		if err != nil {
			return nil, fmt.Errorf("GenerateTarget error - failed to convert rough step: %v", err)
		}
		detailedSteps = append(detailedSteps, dSteps...)

		if s.Commit != "" {
			state.prevCommit = plumbing.NewHash(s.Commit)
		}

		state.currentSeqNo++
	}

	return detailedSteps, nil
}

//////////////////////////////////////////////////////
// RoughStep to DetailedStep conversion methods
//////////////////////////////////////////////////////

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

func (state *InnerState) commitConvert(s *RoughStep, repo *git.Repository) ([]DetailedStep, error) {
	var detailedSteps []DetailedStep
	prevColumn := state.currentColumn

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
	if prevColumn != "Source Code" {
		fileTreeStep := state.fileTreeStep(s, files[0])
		detailedSteps = append(detailedSteps, fileTreeStep)
	}

	// file steps
	for i, file := range files {
		// if prev step is "Source Code", then fileTreeStep is skipped, so the commit should be included in the 0-th openFileStep
		includeCommit := prevColumn == "Source Code" && i == 0

		openFileStep := state.openFileStep(s, i, file, includeCommit)
		detailedSteps = append(detailedSteps, openFileStep)
		if i == 5 {
			break
		}
	}
	state.currentColumn = "Source Code"

	return detailedSteps, nil
}

func (state *InnerState) terminalConvert(s *RoughStep, repo *git.Repository) ([]DetailedStep, error) {
	var detailedSteps []DetailedStep

	// check if it's a valid terminal step
	if s.Instruction == "" && s.Instruction2 == "" {
		return nil, fmt.Errorf("step is missing both 'instruction' and 'instruction2', phase = '%s', type = '%s', comment = '%s'", s.Phase, s.Type, s.Comment)
	}

	// insert move-to-terminal step if current column != "Terminal"
	if state.currentColumn != "Terminal" && state.currentSeqNo != 0 {
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

//////////////////////////////////////////////////////
// DetailedStep generation methods
//////////////////////////////////////////////////////

func (state *InnerState) fileTreeStep(s *RoughStep, file string) DetailedStep {
	subId := "fileTreeStep"
	stepId := state.uuidFinder.FindOrGenerateUUID(s, subId)
	step := DetailedStep{
		// fields to make the step searchable for re-generation
		FromRoughStep: true,
		ParentStep:    s.Step,
		SubID:         subId,
		// other fields
		Step:                stepId,
		FocusColumn:         "Source Code",
		Commit:              s.Commit,
		IsFoldFileTree:      false,
		DefaultOpenFilePath: file,
	}

	return step
}

func (state *InnerState) openFileStep(s *RoughStep, index int, file string, includeCommit bool) DetailedStep {
	subId := fmt.Sprintf("openFileStep-%d", index)
	stepId := state.uuidFinder.FindOrGenerateUUID(s, subId)

	var commit string
	if includeCommit {
		commit = s.Commit
	}

	step := DetailedStep{
		// fields to make the step searchable for re-generation
		FromRoughStep: true,
		ParentStep:    s.Step,
		SubID:         subId,
		// other fields
		Step:                stepId,
		FocusColumn:         "Source Code",
		DefaultOpenFilePath: file,
		IsFoldFileTree:      true,
		Commit:              commit,
	}

	return step
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
	}

	return step
}

//////////////////////////////////////////////////////
// Other utils
//////////////////////////////////////////////////////

func (state *InnerState) isColumnExist(colName string) bool {
	for _, col := range state.existingCols {
		if col == colName {
			return true
		}
	}
	return false
}

func (state *InnerState) appendColumnIfNotExist(colName string) {
	for _, col := range state.existingCols {
		if col == colName {
			// if already exists, do nothing
			return
		}
	}

	// here we didn't find the column, so append it
	for i, col := range state.existingCols {
		if col == "" {
			state.existingCols[i] = colName
			break
		}
	}
}

func (ds *DetailedStep) setColumns(cols [5]string) bool {
	ds.Column1 = cols[0]
	ds.Column2 = cols[1]
	ds.Column3 = cols[2]
	ds.Column4 = cols[3]
	ds.Column5 = cols[4]
	return false
}
