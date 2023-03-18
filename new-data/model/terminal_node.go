package model

func (n *TerminalNode) markCommandExecuted(command string) {
	falseValue := false
	n.Content = TerminalCommand{Command: &command, BeforeExecution: &falseValue}
}
