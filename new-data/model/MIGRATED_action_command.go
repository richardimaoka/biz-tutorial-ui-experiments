package model

import (
	"encoding/json"
	"fmt"
)

func (c ActionCommand) MarshalJSON() ([]byte, error) {
	fmt.Println("ActionCommand MarshalJSON")
	extendedCommand := struct {
		ActionType       string
		Command          string
		TerminalName     string
		UpdateTerminal   UpdateTerminal
		UpdateSourceCode UpdateSourceCode
	}{
		"ActionCommand",
		c.Command,
		c.TerminalName,
		c.UpdateTerminal,
		c.UpdateSourceCode,
	}

	return json.Marshal(extendedCommand)
}
