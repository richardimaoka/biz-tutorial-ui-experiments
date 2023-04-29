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
	switch action := p.nextAction.(type) {
	case *ActionCommand:
		// 1.1 source code mutation
		if err := p.sourceCode.ApplyDiff(action.Diff); err != nil {
			return fmt.Errorf("failed to apply next action, %s", err)
		}

		// 1.2 terminal mutation
		terminal, ok := p.terminalMap[action.TerminalName]
		if !ok {
			return fmt.Errorf("failed to apply next action, terminal [%s] does not exist", action.TerminalName)
		}
		terminal.WriteCommand(action.Command)
		if action.Output != nil {
			terminal.WriteOutput(*action.Output)
		}
		if action.CurrentDirectory != nil {
			terminal.ChangeCurrentDirectory(*action.CurrentDirectory)
		}

		return nil
	case *ManualUpdate:
		// 2.1 source code mutation
		if err := p.sourceCode.ApplyDiff(action.Diff); err != nil {
			return fmt.Errorf("failed to apply next action, %s", err)
		}

		return nil
	default:
		// this should never happen
		return fmt.Errorf("unknown action type %T", action)
	}
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
