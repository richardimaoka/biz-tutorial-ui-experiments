package state

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type BrowserColumn struct {
	browser *Browser
}

func NewBrowserColumn(fields BrowserFields, tutorial string) *BrowserColumn {
	return &BrowserColumn{
		browser: NewBrowser(
			tutorial,
			fields.BrowserImagePath,
			fields.BrowserImageWidth,
			fields.BrowserImageHeight,
		),
	}
}

func (c *BrowserColumn) Update(fields *BrowserFields) error {
	c = &BrowserColumn{
		browser: NewBrowser(
			c.browser.image.tutorial,
			fields.BrowserImagePath,
			fields.BrowserImageWidth,
			fields.BrowserImageHeight,
		),
	}

	return nil
}

func (c *BrowserColumn) ToGraphQL() *model.BrowserColumn {
	return &model.BrowserColumn{
		Browser: c.browser.ToGraphQL(),
	}
}

func (c *BrowserColumn) ToGraphQLColumnWrapper() *model.ColumnWrapper {
	return &model.ColumnWrapper{
		Column:            c.ToGraphQL(),
		ColumnName:        "Browser",
		ColumnDisplayName: stringRef("BrowserColumn"),
	}
}
