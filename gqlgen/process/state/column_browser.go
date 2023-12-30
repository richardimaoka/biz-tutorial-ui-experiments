package state

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type BrowserColumn struct {
	tutorial string
	browser  *Browser
	modal    *Modal
}

func NewBrowserColumn(tutorial string) *BrowserColumn {
	return &BrowserColumn{
		tutorial: tutorial,
		browser:  NewBrowser(),
	}
}

func (c *BrowserColumn) Open(fields *BrowserFields, modalFields *ModalFields) error {
	err := c.browser.SetImage(
		c.tutorial,
		fields.BrowserImagePath,
		fields.BrowserImageWidth,
		fields.BrowserImageHeight,
	)
	if err != nil {
		return fmt.Errorf("BrowserColumn Open() failed, %s", err)
	}

	if modalFields.ModalContents == "" {
		// clean up
		c.modal = nil
	} else {
		c.modal = &Modal{
			markdownBody: modalFields.ModalContents,
			position:     modalFields.ModalPosition,
		}
	}

	return nil
}

func (c *BrowserColumn) Update(fields *BrowserFields, modalFields *ModalFields) error {
	errorPrefix := fmt.Errorf("Update() failed")

	var err error
	switch fields.BrowserStepType {
	case BrowserOpen:
		err = c.Open(fields, modalFields)
	case BrowserMove:
		// no update is needed, just changing FocusColumn is fine
	default:
		err = fmt.Errorf("browser step type = '%s' is not implemented yet", fields.BrowserStepType)
	}
	// check if error happend
	if err != nil {
		return fmt.Errorf("%s failed, %s", errorPrefix, err)
	}

	return nil
}

func (c *BrowserColumn) ToGraphQL() *model.BrowserColumn {
	var browserModel *model.Browser
	if c.browser != nil {
		browserModel = c.browser.ToGraphQL()
	}

	return &model.BrowserColumn{
		Browser: browserModel,
	}
}

func (c *BrowserColumn) ToGraphQLColumnWrapper() *model.ColumnWrapper {
	var modal *model.Modal
	if c.modal != nil {
		modal = c.modal.ToGraphQL()
	}

	return &model.ColumnWrapper{
		Column:            c.ToGraphQL(),
		ColumnName:        "Browser",
		ColumnDisplayName: stringRef("BrowserColumn"),
		Modal:             modal,
	}
}
