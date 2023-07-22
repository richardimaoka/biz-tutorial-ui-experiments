package model

import "encoding/json"

func (this TerminalColumn) MarshalJSON() ([]byte, error) {
	extendedOutput := struct {
		TypeName string    `json:"__typename"`
		Terminal *Terminal `json:"terminal,omitempty"`
	}{
		TypeName: "TerminalColumn",
		Terminal: this.Terminal,
	}

	return json.Marshal(extendedOutput)
}
