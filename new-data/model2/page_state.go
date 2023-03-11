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

func (p *PageState) gotoNextStep() error {
	nextNextStep, err := calcNextStep(*p.NextStep)
	if err != nil {
		return err
	}

	p.PrevStep = p.Step
	p.Step = p.NextStep
	p.NextStep = &nextNextStep

	return nil
}

func (p *PageState) canTypeInTerminalCommand(terminalName string) error {
	_, err := calcNextStep(*p.NextStep)
	if err != nil {
		return err
	}

	terminal := p.getTerminal(terminalName)
	if terminal == nil {
		return fmt.Errorf("failed to type in command, terminal with name = %s not found", terminalName)
	}

	return nil
}

// public methods

func (p *PageState) TypeInTerminalCommand(command, terminalName string) error {
	if err := p.canTypeInTerminalCommand(terminalName); err != nil {
		return fmt.Errorf("failed to type in command, %s", err)
	}

	terminal := p.getTerminal(terminalName)
	if terminal == nil {
		return fmt.Errorf("failed to type in command, terminal with name = %s not found", terminalName)
	}

	// type in command
	if err := terminal.TypeInCommand(command); err != nil {
		return fmt.Errorf("failed to type in command, %s", err)
	}

	// update step
	if err := p.gotoNextStep(); err != nil {
		return fmt.Errorf("failed to type in command, calc next step failed %s", err)
	}

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

	// 1.3 pre-conditions for TerminalCommand.UpdateSourceCode

	// pre-condition AddFiles does not have a matching node in fileTree, and the parent dir
	// pre-condition UpdateFiles has matching node in fileTree
	// pre-condition DeleteFiles has matching node in fileTree

	// 2.1 Terminal update
	// p.updateTerminal(command.UpdateTerminal)

	//execute command!
	// if err := terminal.executeEffect(); err != nil {

	// Process UpdateTerminal.Output

	// 2.2 SourceCode update

	// if len(command.UpdateTerminal.CurrentDirectory) > 0 {
	// 	terminal.CurrentDirectory = []*string{}
	// 	for _, d := range command.UpdateTerminal.CurrentDirectory {
	// 		terminal.CurrentDirectory = append(terminal.CurrentDirectory, &d)
	// 	}
	// }

	// Process UpdateSourceCode.AddDirectories

	//TODO: sort FileTree

	// update step
	p.gotoNextStep()

	// return fmt.Errorf("runTerminalCommand() failed, terminal with name = %s not found", command.TerminalName)
	return nil
}
