package state2

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type BrowserColumn struct {
}

func NewBrowserColumn() *BrowserColumn {
	return &BrowserColumn{}
}

func (c *BrowserColumn) Update(fields *BrowserFields) error {
	return nil
}

func (c *BrowserColumn) ToGraphQLColumnWrapper() *model.ColumnWrapper2 {
	return nil
}
