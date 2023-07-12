package process

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/read"
)

// Other packages don't use this, so unexported (lowercase) struct
type StepEntry struct {
	// Uppercase fields to allow json dump for testing
	Step                   string                       `json:"step"`
	NColumns               int                          `json:"nColumns"`
	PrevStep               string                       `json:"prevStep,omitempty"`
	NextStep               string                       `json:"nextStep,omitempty"`
	BackgroundImageColumn  *read.BackgroundImageColumn  `json:"backgroundImageColumn,omitempty"`
	ImageDescriptionColumn *read.ImageDescriptionColumn `json:"imageDescriptionColumn,omitempty"`
	MarkdownColumn         *read.MarkdownColumn         `json:"markdownColumn,omitempty"`
}

func (this StepEntry) ToGraphQLColumns() []*model.ColumnWrapper {
	var colWrappers []*model.ColumnWrapper
	for i := 0; i < this.NColumns; i++ {

		if this.BackgroundImageColumn != nil && this.BackgroundImageColumn.Column == i {
			column := ToGraphQLBgImgCol(this.BackgroundImageColumn)
			colWrappers = append(colWrappers, &model.ColumnWrapper{Column: column})
		}

		if this.ImageDescriptionColumn != nil && this.ImageDescriptionColumn.Column == i {
			column := ToGraphQLImgDescCol(this.ImageDescriptionColumn)
			colWrappers = append(colWrappers, &model.ColumnWrapper{Column: column})
		}

		if this.MarkdownColumn != nil && this.MarkdownColumn.Column == i {
			column := ToGraphQLMarkdownColumn(this.MarkdownColumn)
			colWrappers = append(colWrappers, &model.ColumnWrapper{Column: column})
		}
	}

	return colWrappers
}
