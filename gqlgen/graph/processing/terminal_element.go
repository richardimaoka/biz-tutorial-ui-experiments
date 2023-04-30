package processing

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type terminalElementProcessor interface {
	String() string
	ToGraphQLModel() model.TerminalElement
}

type terminalCommandProcessor struct {
	promptExpression string
	promptSymbol     rune
	command          string
}

type terminalOutputProcessor struct {
	output string
}

func (t *terminalCommandProcessor) String() string {
	//TODO: reflect promptExpression and promptSymbol
	return t.command
}

func (t *terminalOutputProcessor) String() string {
	return t.output
}

func (t *terminalCommandProcessor) ToGraphQLModel() model.TerminalElement {
	falseValue := false
	return &model.TerminalCommand{
		BeforeExecution: &falseValue,
		Command:         &t.command, // t is effectively immutable, so no need to copy
	}
}

func (t *terminalOutputProcessor) ToGraphQLModel() model.TerminalElement {
	return &model.TerminalOutput{
		Output: &t.output, // t is effectively immutable, so no need to copy
	}
}
