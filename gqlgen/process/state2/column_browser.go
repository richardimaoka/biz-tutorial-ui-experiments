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

func (c *BrowserColumn) ToGraphQL() *model.BrowserColumn2 {
	return nil //&model.BrowserColumn2{}
}

func (c *BrowserColumn) ToGraphQLColumnWrapper() *model.ColumnWrapper2 {
	return &model.ColumnWrapper2{
		Column:            c.ToGraphQL(),
		ColumnName:        "Browser",
		ColumnDisplayName: stringRef("BrowserColumn"),
	}
}
