package processing

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type PageStateProcessor struct {
	step          *stepProcessor
	terminalMap   map[string]*TerminalProcessor
	sourceCode    *SourceCodeProcessor
	markdown      *MarkdownProcessor
	nextOperation *PageStateOperation
	nextState     *PageStateProcessor
}

func (p *PageStateProcessor) cloneCurrentState() *PageStateProcessor {
	clonedTerminalMap := make(map[string]*TerminalProcessor)
	for k, t := range p.terminalMap {
		clonedTerminalMap[k] = t.Clone()
	}

	return &PageStateProcessor{
		step:        p.step.Clone(),
		terminalMap: clonedTerminalMap,
		sourceCode:  p.sourceCode.Clone(),
		markdown:    p.markdown.Clone(),
		// not cloning nextXxx as they are updated in actual next step
	}
}

func (p *PageStateProcessor) transition(nextStep string, nextOperation *PageStateOperation) error {
	errorPreceding := "failed to apply operation"

	sourceCodeOp := nextOperation.SourceCodeOperation
	if sourceCodeOp != nil {
		if err := p.sourceCode.Transition(nextStep, sourceCodeOp); err != nil {
			return fmt.Errorf("%s, %s", errorPreceding, err)
		}
	}

	terminalOp := nextOperation.TerminalOperation
	if terminalOp != nil {
		terminal, ok := p.terminalMap[terminalOp.GetTerminalName()]
		if !ok {
			return fmt.Errorf("%s, terminal [%s] does not exist", errorPreceding, terminalOp.GetTerminalName())
		}
		terminal.Transition(nextStep, terminalOp)
	}

	markdownOp := nextOperation.MarkdownOperation
	if markdownOp != nil {
		if err := p.markdown.Transition(nextStep, *markdownOp); err != nil {
			return fmt.Errorf("%s, markdown transition failed, %s", errorPreceding, err)
		}
	}

	return nil
}

//------------------------------------------------------------
// public methods
//------------------------------------------------------------

func NewPageStateProcessor() *PageStateProcessor {
	terminalMap := make(map[string]*TerminalProcessor)
	terminalMap["default"] = NewTerminalProcessor("default")

	init := PageStateProcessor{
		step:        NewStepProcessor(),
		terminalMap: terminalMap,
		sourceCode:  NewSourceCodeProcessor(),
		markdown:    NewMarkdownProcessor(),
	}

	return &init
}

func NewPageStateGitProcessorFromGit(repoUrl string) (*PageStateProcessor, error) {
	sourceCode, err := SourceCodeProcessorFromGit(repoUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create PageStateGitProcessorFromGit, %s", err)
	}

	terminalMap := make(map[string]*TerminalProcessor)
	terminalMap["default"] = NewTerminalProcessor("default")

	init := PageStateProcessor{
		step:        NewStepProcessor(),
		terminalMap: terminalMap,
		sourceCode:  sourceCode,
		markdown:    NewMarkdownProcessor(),
	}

	return &init, nil
}

func (p *PageStateProcessor) RegisterNext(nextStep string, op *PageStateOperation) error {
	cloned := p.cloneCurrentState()
	if err := cloned.transition(nextStep, op); err != nil {
		return fmt.Errorf("init page state failed, %s", err)
	}
	p.nextOperation = op
	p.nextState = cloned
	p.step.nextStep = nextStep

	return nil
}

func (p *PageStateProcessor) TransitionToNext() error {
	// 1. verify nextNextAction
	if p.nextState == nil {
		return fmt.Errorf("TransitionToNext() in PageStateProcessor failed at step = %s, next state is nil", p.step.currentStep)
	}

	// 2. transition to nextState
	p.sourceCode = p.nextState.sourceCode
	p.terminalMap = p.nextState.terminalMap
	p.markdown = p.nextState.markdown

	// 3. update step, nextAction & nextState
	p.nextOperation = nil
	p.nextState = nil
	p.step.prevStep = p.step.currentStep
	p.step.currentStep = p.step.nextStep
	p.step.nextStep = ""

	return nil
}

func (p *PageStateProcessor) ToGraphQLPageState() *model.PageState {
	terminals := []*model.Terminal{}
	for _, t := range p.terminalMap {
		terminals = append(terminals, t.ToGraphQLTerminal())
	}

	// var nextAction model.NextAction
	// if p.nextAction != nil {
	// 	nextAction = p.nextAction.ToGraphQLNextAction()
	// } else {
	// 	nextAction = model.NextAction{
	// 		Content: nil,
	// 	}
	// }

	return &model.PageState{
		Step:       &p.step.currentStep,
		NextStep:   &p.step.nextStep,
		PrevStep:   &p.step.prevStep,
		SourceCode: p.sourceCode.ToGraphQLModel(),
		Terminals:  terminals,
		Markdown:   p.markdown.ToGraphQLMarkdown(),
	}
}
