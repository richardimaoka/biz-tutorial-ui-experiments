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

func (p *Page) recursivelyFindPrevStep(currentStepIndex int) *Step {
	prevStepIndex := currentStepIndex - 1

	if prevStepIndex < 0 {
		return nil
	} else {
		prevStep := p.steps[prevStepIndex]

		if prevStep.IsTrivial {
			return p.recursivelyFindPrevStep(prevStepIndex)
		}

		return &prevStep
	}
}

// The previous non-trivial step ID
func (p *Page) EffectivePrevStep() *Step {
	return p.recursivelyFindPrevStep(p.currentStepIndex)
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

func (p *Page) cleanUpPrevStep() error {
	// If this is the very first step, no need to clean up
	if !p.hasPrev() {
		return nil
	}

	prevStep := p.steps[p.currentStepIndex-1]
	currentStep := p.steps[p.currentStepIndex]

	// If prev step column is same as current step column, no need to clean up
	if prevStep.FocusColumn == currentStep.FocusColumn {
		return nil
	}

	// Here, prev step column != current step column, so clean up the prev column
	switch prevStep.FocusColumn {
	case SourceColumnType:
		if p.sourceCodeColumn == nil {
			return fmt.Errorf("failed to clean up as prev source column = nil")
		}
		return p.sourceCodeColumn.CleanUpPrevStep()
	case TerminalColumnType:
		// if p.terminalColumn == nil {
		// 	return fmt.Errorf("failed to clean up as prev source column = nil")
		// }
		// return p.terminalColumn.CleanUp()
		return nil
	case BrowserColumnType:
		// if p.browserColumn == nil {
		// 	return fmt.Errorf("failed to clean up as prev source column = nil")
		// }
		// return p.browserColumn.CleanUp()
		return nil
	default:
		return nil
	}
}

func (p *Page) processStep(step *Step) error {
	// Using switch, instead of interface, because stat changes pile up in the member fields
	// of the Page struct. So it is awkward to switch the implementation of the page itself
	// or the members of the page upon every step
	switch step.FocusColumn {
	case SourceColumnType:
		if p.sourceCodeColumn == nil {
			p.sourceCodeColumn = NewSourceColumn(p.repo, p.projectDir, p.tutorial)
			p.columns = append(p.columns, p.sourceCodeColumn)
		}
		return p.sourceCodeColumn.Update(&step.SourceFields)

	case TerminalColumnType:
		if p.terminalColumn == nil {
			p.terminalColumn = NewTerminalColumn()
			p.columns = append(p.columns, p.terminalColumn)
		}
		return p.terminalColumn.Update(step.StepId, &step.TerminalFields)

	case BrowserColumnType:
		if p.browserColumn == nil {
			p.browserColumn = NewBrowserColumn()
			p.columns = append(p.columns, p.browserColumn)
		}
		return p.browserColumn.Update(&step.BrowserFields)

	default:
		return fmt.Errorf("Failed to process step = '%s', column type = '%s' is not implemented", step.StepId, step.FocusColumn)
	}
}

func (p *Page) ProcessCurrentStep() error {
	currentStep := p.steps[p.currentStepIndex]

	// Clean up prev step if necessary
	if err := p.cleanUpPrevStep(); err != nil {
		return fmt.Errorf("Failed to clean up before step = '%s', %s", currentStep.StepId, err)
	}

	// Process current step
	if err := p.processStep(&currentStep); err != nil {
		return fmt.Errorf("Failed to process step = '%s', %s", currentStep.StepId, err)
	}

	// if everything is ok, then exit
	return nil
}

func (p *Page) hasPrev() bool {
	return 0 < p.currentStepIndex
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

	if effectivePrevStep := p.EffectivePrevStep(); effectivePrevStep != nil {
		prevStepId = stringRef(effectivePrevStep.StepId)
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
