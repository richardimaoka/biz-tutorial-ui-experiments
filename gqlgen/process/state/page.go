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

func (p *Page) ToNextStep() error {
	// Check if next step exists
	if !p.HasNext() {
		return fmt.Errorf("No next step after step = '%s'", p.CurrentStepId())
	}

	// Process next step
	var err error
	nextStep := p.steps[p.currentStepIndex]
	switch nextStep.FocusColumn {
	case SourceColumnType:
		if p.sourceCodeColumn == nil {
			p.sourceCodeColumn = NewSourceColumn(p.repo, p.projectDir, p.tutorial)
			p.columns = append(p.columns, p.sourceCodeColumn)
		}
		err = p.sourceCodeColumn.Update(&nextStep.SourceFields)

	case TerminalColumnType:
		if p.terminalColumn == nil {
			p.terminalColumn = NewTerminalColumn()
			p.columns = append(p.columns, p.terminalColumn)
		}
		err = p.terminalColumn.Update(nextStep.StepId, &nextStep.TerminalFields)

	case BrowserColumnType:
		if p.browserColumn == nil {
			p.browserColumn = NewBrowserColumn()
			p.columns = append(p.columns, p.browserColumn)
		}
		err = p.browserColumn.Update(&nextStep.BrowserFields)

	default:
		err = fmt.Errorf("Update failed to process step = '%s', column type = '%s' is not implemented", nextStep.StepId, nextStep.FocusColumn)
	}

	// checi if error happend
	if err != nil {
		return fmt.Errorf("Update failed to process step = '%s', %s", nextStep.StepId, err)
	}

	// if everything is ok, then exit
	p.currentStepIndex++
	return nil
}

func (p *Page) hasPrev() bool {
	return 1 < p.currentStepIndex
}

func (p *Page) ToGraphQL() *model.Page {
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

	return &model.Page{
		Columns:   modelColumns,
		Step:      currentStepId,
		NextStep:  nextStepId,
		PrevStep:  prevStepId,
		IsTrivial: isTrivial,
	}
}
