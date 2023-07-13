package state

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

type TerminalElementType string

const (
	TerminalCommand TerminalElementType = "COMMAND"
	TerminalOutput  TerminalElementType = "OUTPUT"
)

type TerminalElement struct {
	Type TerminalElementType
	Text string `json:"text"`
}

func (p *TerminalElement) ToGraphQLTerminalElement() model.TerminalElement {
	// copy to avoid mutation effect afterwards
	text := internal.StringRef(p.Text)
	falseValue := false

	switch p.Type {
	case TerminalCommand:
		return &model.TerminalCommand{
			BeforeExecution: &falseValue,
			Command:         text,
		}
	case TerminalOutput:
		return &model.TerminalOutput{
			Output: text,
		}
	default:
		return nil
	}
}
