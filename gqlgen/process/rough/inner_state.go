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
	existingCols  UsedColumns //fixed size, according to DetailedStep
	uuidFinder    *UUIDFinder
	prevCommit    string
}

type UsedColumns = [5]string
type CurrentColumn = string

var EmptyColumns = UsedColumns{}

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
		dSteps, err := state.Conversion(&s)
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

func (state *InnerState) Conversion(s *RoughStep) ([]DetailedStep, error) {
	var steps []DetailedStep
	var usedColumns []string
	var currentColumn string
	var existingCols [5]string
	var err error

	// call internal conversion logic
	switch s.Type {
	case "terminal":
		steps, existingCols, err = terminalConvertInternal(s, state.uuidFinder, state.existingCols, state.repo, state.currentColumn, state.prevCommit, state.currentSeqNo)
	case "commit":
		steps, existingCols, err = commitConvertInternal(s, state.uuidFinder, state.existingCols, state.repo, state.currentColumn, state.prevCommit)
	case "source error":
		steps, existingCols, err = sourceErrorConvertInternal(s, state.uuidFinder, state.existingCols)
	case "browser":
		steps, existingCols, err = browserConvertInternal(s, state.uuidFinder, state.existingCols)
	case "markdown":
		steps, existingCols, err = markdownConvertInternal(s, state.uuidFinder, state.existingCols)
	default:
		return nil, fmt.Errorf("unknown type = '%s' for step = '%s'", s.Type, s.Step)
	}

	// check if results are valid
	if err != nil {
		return nil, err
	}

	// - udpate the state
	currentColumn = getCurrentColumn(existingCols)
	if currentColumn != "" {
		state.currentColumn = currentColumn
	} else if len(usedColumns) != 0 {
		state.currentColumn = usedColumns[len(usedColumns)-1]
	}

	return steps, nil
}

/////////////////////////////////////////////////////
// RoughStep to DetailedStep internal methods
//////////////////////////////////////////////////////

func commitConvertInternal(
	s *RoughStep,
	uuidFinder *UUIDFinder,
	existingColumns UsedColumns,
	repo *git.Repository,
	prevColumn string,
	prevCommit string,
) ([]DetailedStep, UsedColumns, error) {
	var detailedSteps []DetailedStep

	// - precondition for RoughStep

	// get info from git
	if s.Commit == "" {
		return nil, EmptyColumns, fmt.Errorf("commit is missing for step = '%s'", s.Step)
	}

	// find files from commit
	files, err := CommitFiles(repo, s.Commit, prevCommit)
	if err != nil {
		return nil, EmptyColumns, fmt.Errorf("failed to get files for commit = %s, prevCommit = %s, %s", s.Commit, prevCommit, err)
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

	usedColumns := appendIfNotExists(existingColumns, "Source Code")
	return detailedSteps, usedColumns, nil
}

func terminalConvertInternal(
	s *RoughStep,
	uuidFinder *UUIDFinder,
	existingColumns UsedColumns,
	repo *git.Repository,
	prevColumn string,
	prevCommit string,
	currentSeqNo int,
) ([]DetailedStep, UsedColumns, error) {
	var steps []DetailedStep

	// - precondition for RoughStep

	// check if it's a valid terminal step
	if s.Instruction == "" && s.Instruction2 == "" {
		return nil, EmptyColumns, fmt.Errorf("step is missing both 'instruction' and 'instruction2', phase = '%s', type = '%s', comment = '%s'", s.Phase, s.Type, s.Comment)
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
	usedColumns := appendIfNotExists(existingColumns, "Terminal")

	// source code steps
	if s.Commit != "" {
		commitSteps, commitUsedColumns, err := commitConvertInternal(s, uuidFinder, usedColumns, repo, "Terminal", prevCommit)
		if err != nil {
			return nil, EmptyColumns, fmt.Errorf("failed to convert commit steps, %s", err)
		}
		steps = append(steps, commitSteps...)
		usedColumns = commitUsedColumns
	}

	return steps, usedColumns, nil
}

func sourceErrorConvertInternal(
	s *RoughStep,
	uuidFinder *UUIDFinder,
	existingColumns UsedColumns,
) ([]DetailedStep, UsedColumns, error) {
	usedColumns := appendIfNotExists(existingColumns, "Source Code")
	sourceErrorStep := sourceErrorStep(s, uuidFinder, usedColumns)
	return []DetailedStep{sourceErrorStep}, usedColumns, nil
}

func browserConvertInternal(
	s *RoughStep,
	uuidFinder *UUIDFinder,
	existingColumns UsedColumns,
) ([]DetailedStep, UsedColumns, error) {
	// precondition for RoughStep
	if s.Instruction == "" {
		return nil, EmptyColumns, fmt.Errorf("instruction is missing for browser step = '%s'", s.Step)
	}

	usedColumns := appendIfNotExists(existingColumns, "Browser")
	split := strings.Split(s.Instruction, ",")

	// browser steps
	var detailedSteps []DetailedStep
	for i, each := range split {
		browserImageName := strings.ReplaceAll(each, " ", "")
		browserStep := browserStep(s, uuidFinder, usedColumns, i, browserImageName)
		detailedSteps = append(detailedSteps, browserStep)
	}

	return detailedSteps, usedColumns, nil
}

func markdownConvertInternal(
	s *RoughStep,
	uuidFinder *UUIDFinder,
	existingColumns UsedColumns,
) ([]DetailedStep, UsedColumns, error) {
	// precondition for RoughStep
	if s.Instruction == "" {
		return nil, EmptyColumns, fmt.Errorf("instruction is missing for markdown step = '%s'", s.Step)
	}

	// markdown step
	usedColumns := appendIfNotExists(existingColumns, "Markdown")
	markdownStep := markdownStep(s, uuidFinder, usedColumns)

	return []DetailedStep{markdownStep}, usedColumns, nil
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

func sourceErrorStep(s *RoughStep, uuidFinder *UUIDFinder, usedColumns UsedColumns) DetailedStep {
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
	step.setColumns(usedColumns)

	return step
}

func browserStep(s *RoughStep, uuidFinder *UUIDFinder, usedColumns UsedColumns, index int, browserImageName string) DetailedStep {
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
	step.setColumns(usedColumns)

	return step
}

func markdownStep(s *RoughStep, uuidFinder *UUIDFinder, usedColumns UsedColumns) DetailedStep {
	subId := "markdownStep"
	stepId := uuidFinder.FindOrGenerateUUID(s, subId)

	step := DetailedStep{
		// fields to make the step searchable for re-generation
		FromRoughStep: true,
		ParentStep:    s.Step,
		SubID:         subId,
		// other fields
		Step:             stepId,
		FocusColumn:      "Markdown",
		MarkdownContents: s.Instruction,
	}
	step.setColumns(usedColumns)

	return step
}

//////////////////////////////////////////////////////
// Other utils
//////////////////////////////////////////////////////

func getCurrentColumn(columns UsedColumns) string {
	for i, col := range columns {
		if col == "" {
			if i == 0 {
				return ""
			} else {
				return columns[i-1]
			}
		}
	}

	return ""
}

func appendIfNotExists(columns UsedColumns, colName string) UsedColumns {
	for _, col := range columns {
		if col == colName {
			// if already exists, do nothing
			return columns
		}
	}

	// here we didn't find the column, so append it
	for i, col := range columns {
		if col == "" {
			// columns is copied as an argument, so we can modify it without affecting the caller
			columns[i] = colName
			break
		}
	}

	return columns
}

// func (state *InnerState) appendColumnsIfNotExist(cols []string) {
// 	for _, col := range cols {
// 		state.appendColumnIfNotExist(col)
// 	}
// }

func (ds *DetailedStep) setColumns(cols UsedColumns) bool {
	ds.Column1 = cols[0]
	ds.Column2 = cols[1]
	ds.Column3 = cols[2]
	ds.Column4 = cols[3]
	ds.Column5 = cols[4]
	return false
}
