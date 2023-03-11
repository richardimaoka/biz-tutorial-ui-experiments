package model2

func (n *TerminalNode) markCommandExecuted(command string) {
	falseValue := false
	n.Content = TerminalCommand{Command: &command, BeforeExecution: &falseValue}
}
