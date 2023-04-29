package processing

import (
	"fmt"
)

type pageStateInternals struct {
	step        *stepProcessor
	terminalMap map[string]*TerminalProcessor
	sourceCode  *SourceCodeProcessor
	nextAction  Action
}

type PageStateProcessor struct {
	pageStateInternals
	preserved *pageStateInternals
}

func (p *PageStateProcessor) beginMutation() {
	terminalMap := make(map[string]*TerminalProcessor)
	for k, t := range p.terminalMap {
		terminalMap[k] = t.Clone()
	}

	p.preserved = &pageStateInternals{
		step:        p.step,
		terminalMap: terminalMap,
		sourceCode:  p.sourceCode.Clone(),
		nextAction:  p.nextAction,
	}
}

func (p *PageStateProcessor) endMutation() {
	p.preserved = nil
}

func (p *PageStateProcessor) rollbackMutation() {
	p.step = p.preserved.step
	p.terminalMap = p.preserved.terminalMap
	p.sourceCode = p.preserved.sourceCode
	p.nextAction = p.preserved.nextAction
}

func (p *PageStateProcessor) applyNextAction() error {
	// if err := p.sourceCode.ApplyDiff(p.nextAction.Effect); err != nil {
	// }
	// // this needs to be applied after diff, as the default file might be created in apply diff
	// if err := p.sourceCode.setDefault(p.nextAction.DefaultOpenFilePath); err != nil {
	// }

	return nil
}

//------------------------------------------------------------
// public methods
//------------------------------------------------------------

func InitPageStateProcessor(firstAction Action) *PageStateProcessor {
	//p.canApplyAction(firstAction)

	return &PageStateProcessor{
		pageStateInternals: pageStateInternals{
			step:        NewStepProcessor(),
			terminalMap: make(map[string]*TerminalProcessor),
			sourceCode:  NewSourceCodeProcessor(),
			nextAction:  firstAction,
		},
		preserved: nil,
	}
}

func (p *PageStateProcessor) StateTransition(nextNextAction Action) error {
	// p.canApplyAction(nextNextAction)

	p.beginMutation()
	defer p.endMutation()

	if err := p.applyNextAction(); err != nil {
		p.rollbackMutation()
		return fmt.Errorf("cannot apply next action %s, %s", p.nextAction, err)
	}

	p.nextAction = nextNextAction

	return nil
}
