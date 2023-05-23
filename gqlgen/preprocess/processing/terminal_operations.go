package processing

type TerminalOperation interface {
	GetTerminalName() string
	GetCommand() string
}

// TODO : same name as model.TerminalCommand. it's ok as you can avoid conflict by using package name, but a different name is better
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

func (o TerminalCommand) GetCommand() string             { return o.Command }
func (o TerminalCommandWithOutput) GetCommand() string   { return o.Command }
func (o TerminalCommandWithCd) GetCommand() string       { return o.Command }
func (o TerminalCommandWithOutputCd) GetCommand() string { return o.Command }
