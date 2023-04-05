package model

import "encoding/json"

func (o TerminalOutput) MarshalJSON() ([]byte, error) {
	extendedOutput := struct {
		ContentType string  `json:"contentType"`
		Output      *string `json:"output"`
	}{
		"TerminalOutput",
		o.Output,
	}

	return json.Marshal(extendedOutput)
}
