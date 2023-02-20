package model

func (c TerminalCommand) toExecutedCommand() TerminalCommand {
	falseValue := false
	c.BeforeExecution = &falseValue

	return c //return copy
}
