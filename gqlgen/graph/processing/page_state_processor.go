package processing

import (
	"errors"
)

type pageStateInternals struct {
	step        *stepProcessor
	terminalMap map[string]*TerminalProcessor
	sourceCode  *SourceCodeProcessor
	nextAction  string
}

type PageStateProcessor struct {
	pageStateInternals
	preserved *pageStateInternals
}

func InitPageStateProcessor(firstAction string) {
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

func (p *PageStateProcessor) StateTransition(nextNextAction string) {
	p.beginMutation()
	defer p.endMutation()

	err := errors.New("some error")
	if err != nil {
		p.rollbackMutation()
	}
}
