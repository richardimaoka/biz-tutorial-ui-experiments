package model

//no pre-condition required, always succeed
func (t *Terminal) typeInCommand(command *ActionCommand) {
	// append terminal node
	trueValue := true
	node := TerminalNode{
		Content: TerminalCommand{
			Command:         &command.Command,
			BeforeExecution: &trueValue,
		},
	}

	// works even if Nodes is nil
	t.Nodes = append(t.Nodes, &node)
}
