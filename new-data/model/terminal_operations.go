package model

type TerminalOperation interface {
	IsTerminalOperation()
}

type ChangeDirectory struct {
	FilePath string `json:"filePath"`
}

func (o ChangeDirectory) IsTerminalOperation() {}

type TypeInCommand struct {
	Command string `json:"Command"`
}

func (o TypeInCommand) IsTerminalOperation() {}

type MarkLastCommandExecuted struct {
	Command string `json:"Command"`
}

func (o MarkLastCommandExecuted) IsTerminalOperation() {}

type WriteOutput struct {
	Output string `json:"Output"`
}

func (o WriteOutput) IsTerminalOperation() {}

type ExecuteCommand struct {
	Command          string `json:"command"`
	Output           string `json:"output"`
	CurrentDirectory string `json:"currentDirectory"`
}

func (o ExecuteCommand) IsTerminalOperation() {}
