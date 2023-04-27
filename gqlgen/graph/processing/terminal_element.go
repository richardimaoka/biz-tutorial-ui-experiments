package processing

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type terminalElementProcessor interface {
	String() string
	ToGraphQLModel() model.TerminalElement
	Clone() terminalElementProcessor
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
	command := t.command // copy to avoid receiver's mutation effect afterwards
	return &model.TerminalCommand{
		BeforeExecution: &falseValue,
		Command:         &command,
	}
}

func (t *terminalOutputProcessor) ToGraphQLModel() model.TerminalElement {
	output := t.output // copy to avoid receiver's mutation effect afterwards
	return &model.TerminalOutput{
		Output: &output,
	}
}

func (t *terminalCommandProcessor) Clone() terminalElementProcessor {
	return &terminalCommandProcessor{
		promptExpression: t.promptExpression, // copy to avoid receiver's mutation effect afterwards
		promptSymbol:     t.promptSymbol,     // copy to avoid receiver's mutation effect afterwards
		command:          t.command,          // copy to avoid receiver's mutation effect afterwards
	}

}

func (t *terminalOutputProcessor) Clone() terminalElementProcessor {
	return &terminalOutputProcessor{
		output: t.output, // copy to avoid receiver's mutation effect afterwards`
	}
}
