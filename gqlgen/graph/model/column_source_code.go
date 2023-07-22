package model

import "encoding/json"

func (this SourceCodeColumn) MarshalJSON() ([]byte, error) {
	extendedOutput := struct {
		TypeName    string      `json:"__typename"`
		Placeholder *string     `json:"_placeholder,omitempty"`
		SourceCode  *SourceCode `json:"sourceCode,omitempty"`
	}{
		TypeName:    "SourceCodeColumn",
		Placeholder: this.Placeholder,
		SourceCode:  this.SourceCode,
	}

	return json.Marshal(extendedOutput)
}
