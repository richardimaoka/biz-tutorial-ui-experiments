package processing

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type PageStateOperation struct {
	SourceCodeOperation SourceCodeOperation
	TerminalOperation   TerminalOperation
	MarkdownOperation   *MarkdownOperation //currently a concrete struct, so making it a pointer to allow nil
}

// TODO: implement this
func (p *PageStateOperation) ToGraphQLNextAction() *model.NextAction {
	return nil
}
