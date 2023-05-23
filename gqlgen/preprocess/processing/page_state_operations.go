package processing

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type PageStateOperation struct {
	SourceCodeOperation SourceCodeOperation
	TerminalOperation   TerminalOperation
	MarkdownOperation   *MarkdownOperation //currently a concrete struct, so making it a pointer to allow nil
}

// TODO: implement this
func (p *PageStateOperation) ToGraphQLNextAction() *model.NextAction {
	var terminalName *string
	var terminalCommand *model.TerminalCommand
	if p.TerminalOperation != nil {
		t := p.TerminalOperation.GetTerminalName()
		terminalName = &t

		cmd := p.TerminalOperation.GetCommand()
		trueValue := true
		terminalCommand = &model.TerminalCommand{
			BeforeExecution: &trueValue,
			Command:         &cmd,
		}
	}

	var markdown *model.Markdown
	if p.MarkdownOperation != nil {
		contents := p.MarkdownOperation.Contents // copy to avoid
		markdown = &model.Markdown{
			Contents: &contents,
		}
	}

	return &model.NextAction{
		TerminalName:    terminalName,
		TerminalCommand: terminalCommand,
		Markdown:        markdown,
	}
}
