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
	Tooltip         string
	BeforeExecution bool
}

type TerminalOutput struct {
	Output  string
	Tooltip string
}

func (p *TerminalCommand) ToGraphQLTerminalElement() model.TerminalElement {
	// copy to avoid mutation effect afterwards
	command := internal.StringRef(p.Command)
	tooltip := internal.StringRef(p.Tooltip)
	beforeExecution := p.BeforeExecution

	return &model.TerminalCommand{
		BeforeExecution: &beforeExecution,
		Command:         command,
		Tooltip:         tooltip,
	}
}

func (p *TerminalOutput) ToGraphQLTerminalElement() model.TerminalElement {
	// copy to avoid mutation effect afterwards
	output := internal.StringRef(p.Output)
	tooltip := internal.StringRef(p.Tooltip)

	return &model.TerminalOutput{
		Output:  output,
		Tooltip: tooltip,
	}

}
