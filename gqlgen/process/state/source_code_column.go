package state

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type SourceCodeColumn struct {
	step string
	sc   SourceCode
}

func NewSourceCodeColumn(initialStep string) *SourceCodeColumn {
	return &SourceCodeColumn{
		step: initialStep,
	}
}

func (c *SourceCodeColumn) ToGraphQLSourceCodeColumn() *model.SourceCodeColumn {
	return &model.SourceCodeColumn{
		SourceCode: c.sc.ToGraphQLSourceCode(),
	}
}
