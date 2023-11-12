package model

import "encoding/json"

func (this TerminalColumn) MarshalJSON() ([]byte, error) {
	extendedOutput := struct {
		TypeName    string      `json:"__typename"`
		Placeholder *string     `json:"_placeholder"`
		Terminals   []*Terminal `json:"terminals"`
	}{
		TypeName:    "TerminalColumn",
		Placeholder: this.Placeholder,
		Terminals:   this.Terminals,
	}

	return json.Marshal(extendedOutput)
}

func (this SourceCodeColumn2) MarshalJSON() ([]byte, error) {
	extendedOutput := struct {
		TypeName    string      `json:"__typename"`
		Placeholder *string     `json:"_placeholder"`
		SourceCode  *SourceCode `json:"sourceCode"`
	}{
		TypeName:    "SourceCodeColumn2",
		Placeholder: this.Placeholder,
		SourceCode:  this.SourceCode,
	}

	return json.Marshal(extendedOutput)
}

func (this BrowserColumn2) MarshalJSON() ([]byte, error) {
	extendedOutput := struct {
		TypeName    string   `json:"__typename"`
		Placeholder *string  `json:"_placeholder"`
		Browser     *Browser `json:"browser"`
	}{
		TypeName:    "BrowserColumn2",
		Placeholder: this.Placeholder,
		Browser:     this.Browser,
	}

	return json.Marshal(extendedOutput)
}
