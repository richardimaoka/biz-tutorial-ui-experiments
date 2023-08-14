package model

import "encoding/json"

func (this MarkdownColumn) MarshalJSON() ([]byte, error) {
	extendedOutput := struct {
		TypeName         string                  `json:"__typename"`
		Placeholder      *string                 `json:"_placeholder,omitempty"`
		Description      *Markdown               `json:"description,omitempty"`
		ContentsPosition *ColumnVerticalPosition `json:"contentsPosition,omitempty"`
	}{
		TypeName:         "MarkdownColumn",
		Placeholder:      this.Placeholder,
		Description:      this.Description,
		ContentsPosition: this.ContentsPosition,
	}

	return json.Marshal(extendedOutput)
}
