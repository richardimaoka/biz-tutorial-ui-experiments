package processing

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type MarkdownProcessor struct {
	step     string
	contents string
}

func NewMarkdownProcessor() *MarkdownProcessor {
	return &MarkdownProcessor{
		step:     "",
		contents: "",
	}
}

func (p *MarkdownProcessor) Transition(nextStep string, operation MarkdownOperation) error {
	p.step = nextStep
	p.contents = operation.Contents
	return nil
}

func (p *MarkdownProcessor) Clone() *MarkdownProcessor {
	return &MarkdownProcessor{p.step, p.contents}
}

func (p *MarkdownProcessor) ToGraphQLMarkdown() *model.Markdown {
	step := p.step         // copy to avoid effect from returned model.Markdown's mutation
	contents := p.contents // copy to avoid effect from returned model.Markdown's mutation

	return &model.Markdown{
		Step:     &step,
		Contents: &contents,
	}
}
