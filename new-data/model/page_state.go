package model

import (
	"fmt"
	"strconv"
)

func (p *PageState) canCalcNextStep() (string, error) {
	stepNumString := *p.NextStep
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

func (p *PageState) canTypeInCommand(command ActionCommand) (*Terminal, error) {
	terminal := p.getTerminal(command.TerminalName)
	if terminal == nil {
		return nil, fmt.Errorf("cannot type in command, terminal with name = %s not found", command.TerminalName)
	}

	if err := terminal.canTypeInCommand(); err != nil {
		return nil, fmt.Errorf("cannot type in command, %s", err)
	}

	return terminal, nil
}

func (p *PageState) canExecuteLastCommand(command ActionCommand) (*Terminal, error) {
	terminal := p.getTerminal(command.TerminalName)
	if terminal == nil {
		return nil, fmt.Errorf("cannot execute last command, terminal with name = %s not found", command.TerminalName)
	}

	if err := terminal.canMarkLastCommandExecuted(command.Command); err != nil {
		return nil, fmt.Errorf("cannot execute last command, %s", err)
	}

	// TODO: bundle ChangeCurrentDirectory, WriteOutput, MarkLastCommandExecuted into one method
	// canWriteOutput() will always fail at this point
	//
	// if command.Output != nil {
	// 	if err := terminal.canWriteOutput(); err != nil {
	// 		return nil, fmt.Errorf("cannot execute last command, %s", err)
	// 	}
	// }

	if command.Effect != nil {
		if err := p.SourceCode.canApplyDiff(command.Effect); err != nil {
			return nil, fmt.Errorf("cannot execute last command, %s", err)
		}
	}

	return terminal, nil
}

// function to check if pageState can apply diff
func (p *PageState) canApplyDiff(diff DiffEffect) error {
	return p.SourceCode.canApplyDiff(diff)
}

// public methods

func NewPageState() *PageState {
	zeroString := "000"
	oneString := "001"
	return &PageState{
		Step:       &zeroString,
		PrevStep:   nil,
		NextStep:   &oneString,
		Terminals:  []*Terminal{NewTerminal("default")},
		SourceCode: NewSourceCode(),
	}
}

func (p *PageState) TypeInCommand(command ActionCommand) error {
	// precondition
	nextNextStep, err := p.canCalcNextStep()
	if err != nil {
		return fmt.Errorf("TypeInCommand failed, %s", err)
	}
	terminal, err := p.canTypeInCommand(command)
	if err != nil {
		return fmt.Errorf("TypeInCommand failed, %s", err)
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
		return fmt.Errorf("ExecuteLastCommand failed, %s", err)
	}
	terminal, err := p.canExecuteLastCommand(command)
	if err != nil {
		return fmt.Errorf("ExecuteLastCommand failed, %s", err)
	}

	// mutation
	terminal.markCommandExecuted(command.Command)
	if command.Output != nil {
		terminal.writeOutput(*command.Output)
	}
	if command.CurrentDirectory != nil {
		terminal.ChangeCurrentDirectory(*command.CurrentDirectory)
	}
	p.SourceCode.preMutation()
	if command.Effect != nil {
		p.SourceCode.applyDiff(command.Effect)
		p.SourceCode.postMutation()
	}
	p.gotoNextStep(nextNextStep)

	return nil
}
