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
func (c TerminalCommand) toExecutedCommand() TerminalCommand {
	falseValue := false
	c.BeforeExecution = &falseValue

	return c //return copy
}
