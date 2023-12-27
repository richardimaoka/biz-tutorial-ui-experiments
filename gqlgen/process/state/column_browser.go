package state

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type BrowserColumn struct {
	imageFilePath   string
	imageFileWidth  int
	imageFileHeight int
}

func NewBrowserColumn(fields BrowserFields) *BrowserColumn {
	return &BrowserColumn{
		fields.BrowserImagePath,
		fields.BrowserImageWidth,
		fields.BrowserImageHeight,
	}
}

func (c *BrowserColumn) Update(fields *BrowserFields) error {
	return nil
}

func (c *BrowserColumn) ToGraphQL() *model.BrowserColumn {
	return nil //&model.BrowserColumn2{}
}

func (c *BrowserColumn) ToGraphQLColumnWrapper() *model.ColumnWrapper {
	return &model.ColumnWrapper{
		Column:            c.ToGraphQL(),
		ColumnName:        "Browser",
		ColumnDisplayName: stringRef("BrowserColumn"),
	}
}
