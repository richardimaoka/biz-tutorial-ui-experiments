package state

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

// Step        *string          `json:"step"`
// NextStep    *string          `json:"nextStep"`
// PrevStep    *string          `json:"prevStep"`
// IsTrivial   *bool            `json:"isTrivial"`
// Columns     []*ColumnWrapper `json:"columns"`
// FocusColumn *string          `json:"focusColumn"`
// Modal       *Modal           `json:"modal"`}
type Page struct {
	repo       *git.Repository
	tutorial   string
	projectDir string

	terminalColumn   *TerminalColumn
	sourceCodeColumn *SourceColumn
	browserColumn    *BrowserColumn

	columns []Column

	steps            []Step
	currentStepIndex int
}

func NewPage(repo *git.Repository, tutorial string, steps []Step) *Page {
	return &Page{
		repo:             repo,
		tutorial:         tutorial,
		steps:            steps,
		currentStepIndex: 0,
	}
}

func (p *Page) CurrentStepId() string {
	currentStep := p.steps[p.currentStepIndex]
	return currentStep.StepId
}

func (p *Page) HasNext() bool {
	return p.currentStepIndex < len(p.steps)-1
}

func (p *Page) IncrementStep() error {
	if p.HasNext() {
		p.currentStepIndex++
		return nil
	} else {
		return fmt.Errorf("Cannot increment step since there is no next step")
	}
}

func (p *Page) ProcessCurrentStep() error {
	currentStep := p.steps[p.currentStepIndex]

	var err error
	switch currentStep.FocusColumn {
	case SourceColumnType:
		if p.sourceCodeColumn == nil {
			p.sourceCodeColumn = NewSourceColumn(p.repo, p.projectDir, p.tutorial)
			p.columns = append(p.columns, p.sourceCodeColumn)
		}
		err = p.sourceCodeColumn.Update(&currentStep.SourceFields)

	case TerminalColumnType:
		if p.terminalColumn == nil {
			p.terminalColumn = NewTerminalColumn()
			p.columns = append(p.columns, p.terminalColumn)
		}
		err = p.terminalColumn.Update(currentStep.StepId, &currentStep.TerminalFields)

	case BrowserColumnType:
		if p.browserColumn == nil {
			p.browserColumn = NewBrowserColumn()
			p.columns = append(p.columns, p.browserColumn)
		}
		err = p.browserColumn.Update(&currentStep.BrowserFields)

	default:
		err = fmt.Errorf("Failed to process step = '%s', column type = '%s' is not implemented", currentStep.StepId, currentStep.FocusColumn)
	}

	// checi if error happend
	if err != nil {
		return fmt.Errorf("Failed to process step = '%s', %s", currentStep.StepId, err)
	}

	// if everything is ok, then exit
	return nil
}

func (p *Page) hasPrev() bool {
	return 1 < p.currentStepIndex
}

func (p *Page) ToGraphQL() *model.Page {
	// Handle columns
	var modelColumns []*model.ColumnWrapper
	for _, c := range p.columns {
		modelColumns = append(modelColumns, c.ToGraphQLColumnWrapper())
	}

	// Handle step IDs
	currentStep := p.steps[p.currentStepIndex]
	var currentStepId, nextStepId, prevStepId *string
	currentStepId = stringRef(currentStep.StepId)
	if p.HasNext() {
		nextStepId = stringRef(p.steps[p.currentStepIndex+1].StepId)
	}
	if p.hasPrev() {
		prevStepId = stringRef(p.steps[p.currentStepIndex-1].StepId)
	}

	// Handle isTrivial
	var isTrivial *bool
	if p.steps[p.currentStepIndex].IsTrivial {
		trueValue := true
		isTrivial = &trueValue
	}

	// Handle FocusColumn
	focusColumn := stringRef(string(currentStep.FocusColumn))

	return &model.Page{
		Columns:     modelColumns,
		FocusColumn: focusColumn,
		Step:        currentStepId,
		NextStep:    nextStepId,
		PrevStep:    prevStepId,
		IsTrivial:   isTrivial,
	}
}
