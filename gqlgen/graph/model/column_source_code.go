package model

import "encoding/json"

func (this SourceCodeColumn) MarshalJSON() ([]byte, error) {
	extendedOutput := struct {
		TypeName   string      `json:"__typename"`
		SourceCode *SourceCode `json:"sourceCode,omitempty"`
	}{
		TypeName:   "SourceCodeColumn",
		SourceCode: this.SourceCode,
	}

	return json.Marshal(extendedOutput)
}
