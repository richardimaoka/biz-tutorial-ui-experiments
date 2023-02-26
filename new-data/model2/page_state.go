package model2

import (
	"fmt"
	"strconv"
)

func calcNextStep(stepNumString string) (string, error) {
	stepNum, err := strconv.Atoi(stepNumString)
	if err != nil {
		return "", fmt.Errorf("next step calc failed, as step %s is not number format", stepNumString)
	}

	formatted := fmt.Sprintf("%03d", stepNum)
	if stepNumString != formatted {
		return "", fmt.Errorf("next step calc failed, as step %s is expected 3-digit number format %s", stepNumString, formatted)
	}

	return fmt.Sprintf("%03d", stepNum+1), nil
}

func (p *PageState) getTerminal(terminalName string) *Terminal {
	var terminal *Terminal // nil as zero value
	for _, t := range p.Terminals {
		if *t.Name == terminalName {
			terminal = t
		}
	}
	return terminal
}

func (p *PageState) gotoNextStep(nextNextStep string) {
	p.PrevStep = p.Step
	p.Step = p.NextStep
	p.NextStep = &nextNextStep
}

// public methods

func (p *PageState) TypeInTerminalCommand(command, terminalName string) error {
	// pre-condition - next step calculation
	nextNextStep, err := calcNextStep(*p.NextStep)
	if err != nil {
		return fmt.Errorf("failed to type in command, %s", err)
	}

	// pre-condition - find command's target terminal
	terminal := p.getTerminal(terminalName)
	if terminal == nil {
		return fmt.Errorf("failed to type in command, terminal with name = %s not found", command.TerminalName)
	}

	// type in command
	terminal.TypeInCommand(command)

	// update step
	p.gotoNextStep(nextNextStep)

	return nil

}
