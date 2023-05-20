package processing

type TerminalOperation interface {
	GetTerminalName() string
	IsTerminalOperation()
}

type TerminalCommand struct {
	TerminalName string
	Command      string
}

type TerminalCommandWithOutput struct {
	TerminalName string
	Command      string
	Output       string
}

type TerminalCommandWithCd struct {
	TerminalName     string
	Command          string
	CurrentDirectory string
}

type TerminalCommandWithOutputCd struct {
	TerminalName     string
	Command          string
	Output           string
	CurrentDirectory string
}

func (o TerminalCommand) GetTerminalName() string             { return o.TerminalName }
func (o TerminalCommandWithOutput) GetTerminalName() string   { return o.TerminalName }
func (o TerminalCommandWithCd) GetTerminalName() string       { return o.TerminalName }
func (o TerminalCommandWithOutputCd) GetTerminalName() string { return o.TerminalName }

func (o TerminalCommand) IsTerminalOperation()             {}
func (o TerminalCommandWithOutput) IsTerminalOperation()   {}
func (o TerminalCommandWithCd) IsTerminalOperation()       {}
func (o TerminalCommandWithOutputCd) IsTerminalOperation() {}
