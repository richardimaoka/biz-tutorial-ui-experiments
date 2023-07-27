package state

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

type TerminalElement interface {
	ToGraphQLTerminalElement() model.TerminalElement
}

type TerminalCommand struct {
	Command         string
	BeforeExecution bool
}

type TerminalOutput struct {
	Output string
}

func (p *TerminalCommand) ToGraphQLTerminalElement() model.TerminalElement {
	// copy to avoid mutation effect afterwards
	command := internal.StringRef(p.Command)
	beforeExecution := p.BeforeExecution

	return &model.TerminalCommand{
		BeforeExecution: &beforeExecution,
		Command:         command,
	}
}

func (p *TerminalOutput) ToGraphQLTerminalElement() model.TerminalElement {
	// copy to avoid mutation effect afterwards
	output := internal.StringRef(p.Output)

	return &model.TerminalOutput{
		Output: output,
	}

}
