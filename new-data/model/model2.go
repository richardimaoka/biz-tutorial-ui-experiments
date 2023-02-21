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

func (p *PageState) executeActionTerminal(action *ActionTerminal) error {
	// 1.1 pre-conditions for next step
	nextNextStep, err := calcNextStep(*p.NextStep)
	if err != nil {
		return fmt.Errorf("failed to run command, %s", err)
	}

	// 1.2 pre-conditions for terminal
	terminal := p.getTerminal(action.TerminalName)
	if terminal == nil {
		return fmt.Errorf("failed run command, terminal with name = %s not found", action.TerminalName)
	}
	if err := terminal.isLastCommandExecutable(); err != nil {
		return fmt.Errorf("failed run command, %s", err)
	}

	//execute command!
	if err := terminal.Execute(*action); err != nil {
		return err
	}

	// update step
	p.gotoNextStep(nextNextStep)

	return nil
}
