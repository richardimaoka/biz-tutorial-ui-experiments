package model

import "encoding/json"

func (c TerminalCommand) MarshalJSON() ([]byte, error) {
	extendedCommand := struct {
		ContentType     string  `json:"contentType"`
		BeforeExecution *bool   `json:"beforeExecution"`
		Command         *string `json:"command"`
	}{
		"TerminalCommand",
		c.BeforeExecution,
		c.Command,
	}

	return json.Marshal(extendedCommand)
}
