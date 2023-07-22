package model

import "encoding/json"

func (this TerminalColumn) MarshalJSON() ([]byte, error) {
	extendedOutput := struct {
		TypeName    string    `json:"__typename"`
		Placeholder *string   `json:"_placeholder,omitempty"`
		Terminal    *Terminal `json:"terminal,omitempty"`
	}{
		TypeName:    "TerminalColumn",
		Placeholder: this.Placeholder,
		Terminal:    this.Terminal,
	}

	return json.Marshal(extendedOutput)
}
