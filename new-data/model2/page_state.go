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

func (p *PageState) canCalcNextStep() (string, error) {
	return calcNextStep(*p.NextStep)
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

func (p *PageState) canTypeInTerminalCommand(terminalName string) (*Terminal, string, error) {
	nextNextStep, err := p.canCalcNextStep()
	if err != nil {
		return nil, "", err
	}

	terminal := p.getTerminal(terminalName)
	if terminal == nil {
		return nil, "", fmt.Errorf("failed to type in command, terminal with name = %s not found", terminalName)
	}

	if err := terminal.canTypeInCommand(); err != nil {
		return nil, "", fmt.Errorf("failed to type in command, %s", err)
	}

	return terminal, nextNextStep, nil
}

// public methods

func (p *PageState) TypeInTerminalCommand(command, terminalName string) error {
	terminal, nextNextStep, err := p.canTypeInTerminalCommand(terminalName)
	if err != nil {
		return fmt.Errorf("failed to type in command, %s", err)
	}

	terminal.typeInCommand(command)
	p.gotoNextStep(nextNextStep)

	return nil
}

func (p *PageState) RunTerminalCommand(command, terminalName string) error {
	// 1.1 pre-conditions for next step
	// nextNextStep, err := calcNextStep(*p.NextStep)
	// if err != nil {
	// 	return fmt.Errorf("failed to run command, %s", err)
	// }

	// pre-condition - find command's target terminal
	terminal := p.getTerminal(terminalName)
	if terminal == nil {
		return fmt.Errorf("failed run command, terminal with name = %s not found", terminalName)
	}

	// terminal.MarkLastCommandExecuted()

	// terminal.ApplyEffect()
	// SourceCode.ApplyEffect

	// update step
	// p.gotoNextStep()

	// return fmt.Errorf("runTerminalCommand() failed, terminal with name = %s not found", command.TerminalName)
	return nil
}
