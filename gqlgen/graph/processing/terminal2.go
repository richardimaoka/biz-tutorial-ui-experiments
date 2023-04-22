package processing

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type Terminal2 struct {
	terminalName     string
	currentDirectory string
	elements         []terminalElement
}

func NewTerminal2(terminalName string) *Terminal2 {
	return &Terminal2{
		terminalName:     terminalName,
		currentDirectory: "",
		elements:         []terminalElement{},
	}
}

func (t *Terminal2) Clone() *Terminal2 {
	return &Terminal2{
		terminalName:     t.terminalName,
		currentDirectory: t.currentDirectory,
		elements:         t.elements,
	}
}

func (t *Terminal2) ToTerminal() *model.Terminal {
	var currentDirectory *string
	if t.currentDirectory != "" {
		currentDirectory = &t.currentDirectory
	}

	var nodes []*model.TerminalNode
	for _, e := range t.elements {
		nodes = append(nodes, &model.TerminalNode{
			Content: e.ToTerminalElement(),
		})
	}

	return &model.Terminal{
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
		//TODO: reflect promptExpression and promptSymbol
		command: command,
	})
}

func (t *Terminal2) WriteOutput(output string) {
	t.elements = append(t.elements, &terminalOutput{
		output: output,
	})
}

func (t *Terminal2) ChangeCurrentDirectory(dir string) {
	t.currentDirectory = dir
}
