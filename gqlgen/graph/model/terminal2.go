package model

type TerminalElement2 interface {
	String() string
}

type Terminal2 struct {
	terminalName     string
	currentDirectory string
	elements         []TerminalElement2
}

func NewTerminal2(terminalName string) *Terminal2 {
	return &Terminal2{
		terminalName:     terminalName,
		currentDirectory: "",
		elements:         []TerminalElement2{},
	}
}

func (t *Terminal2) Clone() *Terminal2 {
	return nil
}

func (t *Terminal2) ToTerminal() *Terminal {
	var currentDirectory *string
	if t.currentDirectory != "" {
		currentDirectory = &t.currentDirectory
	}

	return &Terminal{
		Name:             &t.terminalName,
		CurrentDirectory: currentDirectory,
		Nodes:            nil,
	}
}

func (t *Terminal2) WriteCommand(commandPrefix, command string) *Terminal2 {
	defaultPromptExpression := ""
	defaultPromptSymbol := '$'
	t.WriteCommandWithPrompt(defaultPromptExpression, defaultPromptSymbol, command)
	return nil
}

func (t *Terminal2) WriteCommandWithPrompt(promptExpression string, promptSymbol rune, command string) *Terminal2 {
	return nil
}

func (t *Terminal2) WriteCommandOutput(output string) error {
	return nil
}

func (t *Terminal2) ChangeCurrentDirectory(dir string) error {
	return nil
}
