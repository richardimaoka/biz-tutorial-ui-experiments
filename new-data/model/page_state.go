package model

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

func (p *PageState) canTypeInTerminalCommand(command ActionCommand) (*Terminal, error) {
	terminal := p.getTerminal(command.TerminalName)
	if terminal == nil {
		return nil, fmt.Errorf("cannot type in command, terminal with name = %s not found", command.TerminalName)
	}

	if err := terminal.canTypeInCommand(); err != nil {
		return nil, err
	}

	return terminal, nil
}

func (p *PageState) canExecuteLastCommand(command ActionCommand) (*Terminal, error) {
	terminal := p.getTerminal(command.TerminalName)
	if terminal == nil {
		return nil, fmt.Errorf("terminal with name = %s not found", command.TerminalName)
	}

	if err := terminal.canMarkLastCommandExecuted(command.Command); err != nil {
		return nil, err
	}

	if command.Output != nil {
		if err := terminal.canWriteOutput(); err != nil {
			return nil, err
		}
	}

	// terminal.ChangeCurrentDirectory() always succeed

	// if command.Effect != nil && p.SourceCode.canApplyDiff(command.Effect) != nil {
	// 	return nil, nil, err
	// }

	return terminal, nil
}

// public methods

func (p *PageState) TypeInTerminalCommand(command ActionCommand) error {
	// precondition
	nextNextStep, err := p.canCalcNextStep()
	if err != nil {
		return err
	}
	terminal, err := p.canTypeInTerminalCommand(command)
	if err != nil {
		return fmt.Errorf("failed to type in command, %s", err)
	}

	// mutation
	terminal.typeInCommand(command.Command)
	p.gotoNextStep(nextNextStep)

	return nil
}

func (p *PageState) ExecuteLastCommand(command ActionCommand) error {
	// precondition
	nextNextStep, err := p.canCalcNextStep()
	if err != nil {
		return err
	}
	terminal, err := p.canExecuteLastCommand(command)
	if err != nil {
		return fmt.Errorf("failed to type in command, %s", err)
	}

	// mutation
	terminal.markCommandExecuted(command.Command)
	if command.Output != nil {
		terminal.writeOutput(*command.Output)
	}
	if command.CurrentDirectory != nil {
		terminal.ChangeCurrentDirectory(*command.CurrentDirectory)
	}
	p.gotoNextStep(nextNextStep)

	return nil
}
