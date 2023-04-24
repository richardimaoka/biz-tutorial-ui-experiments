package processing

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type terminalElement interface {
	String() string
	ToTerminalElement() model.TerminalElement
}

type terminalCommand struct {
	promptExpression string
	promptSymbol     rune
	command          string
}

type terminalOutput struct {
	output string
}

func (t *terminalCommand) String() string {
	//TODO: reflect promptExpression and promptSymbol
	return t.command
}

func (t *terminalOutput) String() string {
	return t.output
}

func (t *terminalCommand) ToTerminalElement() model.TerminalElement {
	falseValue := false
	command := t.command // copy to avoid effect from receiver's mutation afterwards
	return &model.TerminalCommand{
		BeforeExecution: &falseValue,
		Command:         &command,
	}
}

func (t *terminalOutput) ToTerminalElement() model.TerminalElement {
	output := t.output // copy to avoid effect from receiver's mutation afterwards
	return &model.TerminalOutput{
		Output: &output,
	}
}
