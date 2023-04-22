package model

type terminalElement2 interface {
	String() string
	ToTerminalElement() TerminalElement
}

type terminalCommand struct {
	promptExpression string
	promptSymbol     string
	command          string
}

func (t *terminalCommand) String() string {
	return t.command
}

func (t *terminalCommand) ToTerminalElement() TerminalElement {
	falseValue := false
	return &TerminalCommand{
		BeforeExecution: &falseValue,
		Command:         &t.command,
	}
}

type Terminal2 struct {
	terminalName     string
	currentDirectory string
	elements         []terminalElement2
}

func NewTerminal2(terminalName string) *Terminal2 {
	return &Terminal2{
		terminalName:     terminalName,
		currentDirectory: "",
		elements:         []terminalElement2{},
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

	var nodes []*TerminalNode
	for _, e := range t.elements {
		nodes = append(nodes, &TerminalNode{
			Content: e.ToTerminalElement(),
		})
	}

	return &Terminal{
		Name:             &t.terminalName,
		CurrentDirectory: currentDirectory,
		Nodes:            nodes,
	}
}

func (t *Terminal2) WriteCommand(command string) {
	defaultPromptExpression := ""
	defaultPromptSymbol := '$'
	t.WriteCommandWithPrompt(defaultPromptExpression, defaultPromptSymbol, command)
}

func (t *Terminal2) WriteCommandWithPrompt(promptExpression string, promptSymbol rune, command string) {
	t.elements = append(t.elements, &terminalCommand{
		command: command,
	})
}

func (t *Terminal2) WriteCommandOutput(output string) error {
	return nil
}

func (t *Terminal2) ChangeCurrentDirectory(dir string) error {
	return nil
}
