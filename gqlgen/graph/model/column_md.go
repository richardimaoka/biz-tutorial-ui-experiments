package model

import "encoding/json"

func (this MarkdownColumn) MarshalJSON() ([]byte, error) {
	extendedOutput := struct {
		TypeName    string    `json:"__typename"`
		Description *Markdown `json:"description,omitempty"`
	}{
		TypeName:    "MarkdownColumn",
		Description: this.Description,
	}

	return json.Marshal(extendedOutput)
}
