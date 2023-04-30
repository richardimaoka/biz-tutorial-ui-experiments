package processing

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type PageStateProcessor struct {
	step        *stepProcessor
	terminalMap map[string]*TerminalProcessor
	sourceCode  *SourceCodeProcessor
	nextAction  Action
	nextState   *PageStateProcessor
}

func (p *PageStateProcessor) cloneForNextAction() *PageStateProcessor {
	clonedTerminalMap := make(map[string]*TerminalProcessor)
	for k, t := range p.terminalMap {
		clonedTerminalMap[k] = t.Clone()
	}

	return &PageStateProcessor{
		step:        p.step.Clone(),
		terminalMap: clonedTerminalMap,
		sourceCode:  p.sourceCode.Clone(),
		// not cloning nextAction & nextState, as they will be set after applyNextAction()
	}
}

func (p *PageStateProcessor) applyAction(nextAction Action) error {
	errorPreceding := "failed to apply action"

	// not p.nextAction but passed-in nextAction, so that this method can also verify nextNextAction
	switch action := nextAction.(type) {
	case *ActionCommand:
		// 1.1. source code mutation
		if err := p.sourceCode.ApplyDiff(action.Diff); err != nil {
			return fmt.Errorf("%s, %s", errorPreceding, err)
		}

		// 1.3. terminal mutation
		terminal, ok := p.terminalMap[action.TerminalName]
		if !ok {
			return fmt.Errorf("%s, terminal [%s] does not exist", errorPreceding, action.TerminalName)
		}
		terminal.WriteCommand(action.Command)
		if action.Output != nil {
			terminal.WriteOutput(*action.Output)
		}
		if action.CurrentDirectory != nil {
			terminal.ChangeCurrentDirectory(*action.CurrentDirectory)
		}
		p.terminalMap[action.TerminalName] = terminal

		return nil

	case *ManualUpdate:
		// 2. source code mutation
		if err := p.sourceCode.ApplyDiff(action.Diff); err != nil {
			return fmt.Errorf("%s, %s", errorPreceding, err)
		}

		return nil

	default:
		// this should never happen
		return fmt.Errorf("%s, unknown action type %T", errorPreceding, action)
	}
}

//------------------------------------------------------------
// public methods
//------------------------------------------------------------

func InitPageStateProcessor(firstAction Action) (*PageStateProcessor, error) {
	terminalMap := make(map[string]*TerminalProcessor)
	terminalMap["default"] = NewTerminalProcessor("default")

	init := PageStateProcessor{
		step:        NewStepProcessor(),
		terminalMap: terminalMap,
		sourceCode:  NewSourceCodeProcessor(),
	}

	cloned := init.cloneForNextAction()
	if err := cloned.applyAction(firstAction); err != nil {
		return nil, fmt.Errorf("init page state failed, %s", err)
	}
	init.nextAction = firstAction
	init.nextState = cloned

	return &init, nil
}

func (p *PageStateProcessor) StateTransition(nextNextAction Action) error {
	// 1. verify nextNextAction
	if p.nextState == nil {
		return fmt.Errorf("state transition failed at step = %s, next transition is nil", p.step.currentStep)
	}

	cloned := p.nextState.cloneForNextAction()
	if err := cloned.applyAction(nextNextAction); err != nil {
		return fmt.Errorf("state transition failed at step = %s, nextNextAction is invalid, %s", p.step.currentStep, err)
	}
	cloned.step.AutoIncrementStep()

	// 2. transition to nextState
	p.sourceCode = p.nextState.sourceCode
	p.sourceCode.SetStep(p.step.nextStep)
	p.terminalMap = p.nextState.terminalMap
	p.step.AutoIncrementStep()

	// 3. update step, nextAction & nextState
	p.nextAction = nextNextAction
	p.nextState = cloned

	return nil
}

func (p *PageStateProcessor) LastTransition() {
	// 1. transition to nextState
	p.sourceCode = p.nextState.sourceCode
	p.sourceCode.SetStep(p.step.nextStep)
	p.terminalMap = p.nextState.terminalMap
	p.step.IncrementStep("")

	// 2. update step, nextAction & nextState
	p.nextAction = nil
	p.nextState = nil
}

func (p *PageStateProcessor) ToGraphQLPageState() *model.PageState {
	terminals := []*model.Terminal{}
	for _, t := range p.terminalMap {
		terminals = append(terminals, t.ToGraphQLTerminal())
	}

	var nextAction model.NextAction
	if p.nextAction != nil {
		nextAction = p.nextAction.ToGraphQLNextAction()
	} else {
		nextAction = model.NextAction{
			Content: nil,
		}
	}

	return &model.PageState{
		Step:       &p.step.currentStep,
		NextStep:   &p.step.nextStep,
		PrevStep:   &p.step.prevStep,
		SourceCode: p.sourceCode.ToGraphQLModel(),
		Terminals:  terminals,
		NextAction: &nextAction,
	}
}
