package model

import "encoding/json"

func (a NextActionTerminal) MarshalJSON() ([]byte, error) {
	extended := struct {
		ContentType string  `json:"contentType"`
		Command     *string `json:"command"`
	}{
		"NextActionTerminal",
		a.Command,
	}

	return json.Marshal(extended)
}
