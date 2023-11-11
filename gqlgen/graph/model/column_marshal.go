package model

import "encoding/json"

func (this TerminalColumn2) MarshalJSON() ([]byte, error) {
	extendedOutput := struct {
		TypeName    string       `json:"__typename"`
		Placeholder *string      `json:"_placeholder"`
		Terminals   []*Terminal2 `json:"terminals"`
	}{
		TypeName:    "TerminalColumn2",
		Placeholder: this.Placeholder,
		Terminals:   this.Terminals,
	}

	return json.Marshal(extendedOutput)
}

func (this SourceCodeColumn2) MarshalJSON() ([]byte, error) {
	extendedOutput := struct {
		TypeName    string       `json:"__typename"`
		Placeholder *string      `json:"_placeholder"`
		SourceCode  *SourceCode2 `json:"sourceCode"`
	}{
		TypeName:    "TerminalColumn2",
		Placeholder: this.Placeholder,
		SourceCode:  this.SourceCode,
	}

	return json.Marshal(extendedOutput)
}
