package state

type Column interface {
	IsColumn()
}

type Columns []Column

func (c *BackgroundImageColumn) IsColumn()  {}
func (c *ImageDescriptionColumn) IsColumn() {}
