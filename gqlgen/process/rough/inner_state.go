package rough

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

type InnerState struct {
	repo          *git.Repository
	currentSeqNo  int
	currentColumn string
	existingCols  [5]string //fixed size, according to DetailedStep
	uuidFinder    *UUIDFinder
	prevCommit    string
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
		// step conversion
		dSteps, err := state.Conversion(&s, state.repo)
		if err != nil {
			return nil, fmt.Errorf("GenerateTarget error - failed to convert rough step: %v", err)
		}
		detailedSteps = append(detailedSteps, dSteps...)

		// inner state update
		if s.Commit != "" {
			state.prevCommit = s.Commit
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
		return state.browserConvert(s)
	default:
		return nil, fmt.Errorf("unknown type = '%s', phase = '%s', comment = '%s'", s.Type, s.Phase, s.Comment)
	}
}

func (state *InnerState) commitConvert(s *RoughStep, repo *git.Repository) ([]DetailedStep, error) {
	detailedSteps, usedColumns, err := commitConvertInternal(s, repo, state.uuidFinder, state.currentColumn, state.prevCommit)
	if err != nil {
		return nil, err
	}
	if len(usedColumns) == 0 {
		return nil, fmt.Errorf("usedColumns is empty")
	}

	// - udpate the state
	state.currentColumn = usedColumns[len(usedColumns)-1]
	state.appendColumnIfNotExist(state.currentColumn)

	return detailedSteps, nil
}

func (state *InnerState) terminalConvert(s *RoughStep, repo *git.Repository) ([]DetailedStep, error) {
	detailedSteps, usedColumns, err := terminalConvertInternal(s, repo, state.uuidFinder, state.currentColumn, state.prevCommit, state.currentSeqNo)
	if err != nil {
		return nil, err
	}
	if len(usedColumns) == 0 {
		return nil, fmt.Errorf("usedColumns is empty")
	}

	// - udpate the state
	state.currentColumn = usedColumns[len(usedColumns)-1]
	state.appendColumnIfNotExist(state.currentColumn)

	return detailedSteps, nil
}

func (state *InnerState) sourceErrorConvert(s *RoughStep, repo *git.Repository) ([]DetailedStep, error) {
	detailedSteps, err := sourceErrorConvertInternal(s, repo, state.uuidFinder)
	if err != nil {
		return nil, err
	}

	// - udpate the state
	state.currentColumn = "Source Code"
	state.appendColumnIfNotExist(state.currentColumn)

	return detailedSteps, nil
}

func (state *InnerState) browserConvert(s *RoughStep) ([]DetailedStep, error) {
	detailedSteps, err := browserConvertInternal(s, state.uuidFinder)
	if err != nil {
		return nil, err
	}

	// - udpate the state
	state.currentColumn = "Browser"
	state.appendColumnIfNotExist(state.currentColumn)

	return detailedSteps, nil
}

func (state *InnerState) markdownConvert(s *RoughStep) ([]DetailedStep, error) {
	detailedSteps, err := markdownConvertInternal(s, state.uuidFinder)
	if err != nil {
		return nil, err
	}

	// - udpate the state
	state.currentColumn = "Markdown"
	state.appendColumnIfNotExist(state.currentColumn)

	return detailedSteps, nil
}

/////////////////////////////////////////////////////
// RoughStep to DetailedStep internal methods
//////////////////////////////////////////////////////

func commitConvertInternal(s *RoughStep, repo *git.Repository, uuidFinder *UUIDFinder, prevColumn, prevCommit string) ([]DetailedStep, []string, error) {
	var detailedSteps []DetailedStep

	// - precondition for RoughStep

	// get info from git
	if s.Commit == "" {
		return nil, nil, fmt.Errorf("commit is missing for step = '%s'", s.Step)
	}

	// find files from commit
	files, err := CommitFiles(repo, s.Commit, prevCommit)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get files for commit = %s, prevCommit = %s, %s", s.Commit, prevCommit, err)
	}

	// - step creation

	// insert file-tree step if prev column != "Source Code"
	if prevColumn != "Source Code" {
		fileTreeStep := fileTreeStep(s, uuidFinder, files[0])
		detailedSteps = append(detailedSteps, fileTreeStep)
	}

	// file steps
	for i, file := range files {
		// if prev step is "Source Code", then fileTreeStep is skipped, so the commit should be included in the 0-th openFileStep
		includeCommit := prevColumn == "Source Code" && i == 0

		openFileStep := openFileStep(s, uuidFinder, i, file, includeCommit)
		detailedSteps = append(detailedSteps, openFileStep)
		if i == 5 {
			break
		}
	}

	usedColumns := []string{"Source Code"}
	return detailedSteps, usedColumns, nil
}

func terminalConvertInternal(s *RoughStep, repo *git.Repository, uuidFinder *UUIDFinder, prevColumn string, prevCommit string, currentSeqNo int) ([]DetailedStep, []string, error) {
	var steps []DetailedStep
	var usedColumns []string

	// - precondition for RoughStep

	// check if it's a valid terminal step
	if s.Instruction == "" && s.Instruction2 == "" {
		return nil, nil, fmt.Errorf("step is missing both 'instruction' and 'instruction2', phase = '%s', type = '%s', comment = '%s'", s.Phase, s.Type, s.Comment)
	}

	// - step creation

	// insert move-to-terminal step if current column != "Terminal"
	if prevColumn != "Terminal" && currentSeqNo != 0 {
		moveToTerminalStep := moveToTerminalStep(s, uuidFinder)
		steps = append(steps, moveToTerminalStep)
	}

	// command step
	cmdStep := terminalCommandStep(s, uuidFinder)
	steps = append(steps, cmdStep)

	// cd step
	if strings.HasPrefix(s.Instruction, "cd ") {
		cmdStep := terminalCdStep(s, uuidFinder)
		steps = append(steps, cmdStep)
	}

	// output step
	if s.Instruction2 != "" {
		outputStep := terminalOutputStep(s, uuidFinder)
		steps = append(steps, outputStep)
	}

	// update used columns
	usedColumns = append(usedColumns, "Terminal")

	// source code steps
	if s.Commit != "" {
		commitSteps, commitColumns, err := commitConvertInternal(s, repo, uuidFinder, "Terminal", prevCommit)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to convert commit steps, %s", err)
		}
		steps = append(steps, commitSteps...)
		usedColumns = append(usedColumns, commitColumns...)
	}

	return steps, usedColumns, nil
}

func sourceErrorConvertInternal(s *RoughStep, repo *git.Repository, uuidFinder *UUIDFinder) ([]DetailedStep, error) {
	var detailedSteps []DetailedStep

	sourceErrorStep := sourceErrorStep(s, uuidFinder)
	detailedSteps = append(detailedSteps, sourceErrorStep)

	return detailedSteps, nil
}

func browserConvertInternal(s *RoughStep, uuidFinder *UUIDFinder) ([]DetailedStep, error) {
	var detailedSteps []DetailedStep

	// precondition for RoughStep
	if s.Instruction == "" {
		return nil, fmt.Errorf("instruction is missing for browser step = '%s'", s.Step)
	}

	// browser steps
	split := strings.Split(s.Instruction, ",")
	for i, each := range split {
		browserImageName := strings.ReplaceAll(each, " ", "")
		browserStep := browserStep(s, uuidFinder, i, browserImageName)
		detailedSteps = append(detailedSteps, browserStep)
	}

	return detailedSteps, nil
}

func markdownConvertInternal(s *RoughStep, uuidFinder *UUIDFinder) ([]DetailedStep, error) {
	var detailedSteps []DetailedStep

	// precondition for RoughStep
	if s.Instruction == "" {
		return nil, fmt.Errorf("instruction is missing for markdown step = '%s'", s.Step)
	}

	// browser steps
	markdownStep := markdownStep(s, uuidFinder)
	detailedSteps = append(detailedSteps, markdownStep)

	return detailedSteps, nil
}

//////////////////////////////////////////////////////
// DetailedStep generation methods
//////////////////////////////////////////////////////

func fileTreeStep(s *RoughStep, uuidFinder *UUIDFinder, file string) DetailedStep {
	subId := "fileTreeStep"
	stepId := uuidFinder.FindOrGenerateUUID(s, subId)
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

func openFileStep(s *RoughStep, uuidFinder *UUIDFinder, index int, file string, includeCommit bool) DetailedStep {
	subId := fmt.Sprintf("openFileStep-%d", index)
	stepId := uuidFinder.FindOrGenerateUUID(s, subId)

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

func moveToTerminalStep(s *RoughStep, uuidFinder *UUIDFinder) DetailedStep {
	subId := "moveToTerminalStep"
	stepId := uuidFinder.FindOrGenerateUUID(s, subId)
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

func terminalOutputStep(s *RoughStep, uuidFinder *UUIDFinder) DetailedStep {
	subId := "terminalOutputStep"
	stepId := uuidFinder.FindOrGenerateUUID(s, subId)
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

func sourceErrorStep(s *RoughStep, uuidFinder *UUIDFinder) DetailedStep {
	subId := "sourceErrorStep"
	stepId := uuidFinder.FindOrGenerateUUID(s, subId)
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

func browserStep(s *RoughStep, uuidFinder *UUIDFinder, index int, browserImageName string) DetailedStep {
	subId := fmt.Sprintf("browserStep-%d", index)
	stepId := uuidFinder.FindOrGenerateUUID(s, subId)
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

func terminalCommandStep(s *RoughStep, uuidFinder *UUIDFinder) DetailedStep {
	subId := "terminalCommandStep"
	stepId := uuidFinder.FindOrGenerateUUID(s, subId)

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
	}

	return step
}

func terminalCdStep(s *RoughStep, uuidFinder *UUIDFinder) DetailedStep {
	currentDir := strings.TrimPrefix(s.Instruction, "cd ")

	subId := "terminalCdStep"
	stepId := uuidFinder.FindOrGenerateUUID(s, subId)

	step := DetailedStep{
		// fields to make the step searchable for re-generation
		FromRoughStep: true,
		ParentStep:    s.Step,
		SubID:         subId,
		// other fields
		Step:         stepId,
		FocusColumn:  "Terminal",
		TerminalType: "cd",
		TerminalName: s.Instruction3, // Go zero value is ""
		CurrentDir:   currentDir,     // Go zero value is ""
	}

	return step
}

func markdownStep(s *RoughStep, uuidFinder *UUIDFinder) DetailedStep {
	subId := "markdownStep"
	stepId := uuidFinder.FindOrGenerateUUID(s, subId)

	step := DetailedStep{
		// fields to make the step searchable for re-generation
		FromRoughStep: true,
		ParentStep:    s.Step,
		SubID:         subId,
		// other fields
		Step:             stepId,
		MarkdownContents: s.Instruction,
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
