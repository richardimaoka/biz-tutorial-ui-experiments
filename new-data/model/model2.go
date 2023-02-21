package model

import "fmt"

func (p *PageState) typeIn(action *ActionTerminal) error {
	// pre-condition - next step calculation
	nextNextStep, err := calcNextStep(*p.NextStep)
	if err != nil {
		return fmt.Errorf("failed to type in command, %s", err)
	}

	// pre-condition - find command's target terminal
	terminal := p.getTerminal(action.TerminalName)
	if terminal == nil {
		return fmt.Errorf("failed to type in command, terminal with name = %s not found", action.TerminalName)
	}

	// type in command
	terminal.typeIn(*action)

	// update step
	p.gotoNextStep(nextNextStep)

	return nil
}
