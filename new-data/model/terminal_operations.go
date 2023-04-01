package model

type TerminalOperation interface {
	IsTerminalOperation()
}

type TypeInCommand struct {
	Command string `json:"Command"`
}

func (o TypeInCommand) IsTerminalOperation() {}

type ExecuteCommand struct {
	Command          string  `json:"command"`
	Output           *string `json:"output"`
	CurrentDirectory *string `json:"currentDirectory"`
}

func (o ExecuteCommand) IsTerminalOperation() {}
