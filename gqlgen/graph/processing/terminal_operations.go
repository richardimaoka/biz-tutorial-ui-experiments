package processing

type TerminalOperation interface {
	IsTerminalOperation()
}

type TerminalCommand struct {
	Command string
}

type TerminalCommandWithOutput struct {
	Command string
	Output  string
}

type TerminalCommandWithCd struct {
	Command          string
	CurrentDirectory string
}

type TerminalCommandWithOutputCd struct {
	Command          string
	Output           string
	CurrentDirectory string
}

func (o TerminalCommand) IsTerminalOperation()             {}
func (o TerminalCommandWithOutput) IsTerminalOperation()   {}
func (o TerminalCommandWithCd) IsTerminalOperation()       {}
func (o TerminalCommandWithOutputCd) IsTerminalOperation() {}
