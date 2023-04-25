package processing

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type TerminalProcessor struct {
	terminalName     string
	currentDirectory string
	elements         []terminalElementProcessor
}

func NewTerminalProcessor(terminalName string) *TerminalProcessor {
	return &TerminalProcessor{
		terminalName:     terminalName,
		currentDirectory: "",
		elements:         []terminalElementProcessor{},
	}
}

func (t *TerminalProcessor) Clone() *TerminalProcessor {
	return &TerminalProcessor{
		terminalName:     t.terminalName,
		currentDirectory: t.currentDirectory,
		elements:         t.elements,
	}
}

func (t *TerminalProcessor) WriteCommand(command string) {
	defaultPromptExpression := ""
	defaultPromptSymbol := '$'
	t.WriteCommandWithPrompt(defaultPromptExpression, defaultPromptSymbol, command)
}

func (t *TerminalProcessor) WriteCommandWithPrompt(promptExpression string, promptSymbol rune, command string) {
	t.elements = append(t.elements, &terminalCommandProcessor{
		promptExpression: promptExpression,
		promptSymbol:     promptSymbol,
		command:          command,
	})
}

func (t *TerminalProcessor) WriteOutput(output string) {
	t.elements = append(t.elements, &terminalOutputProcessor{
		output: output,
	})
}

func (t *TerminalProcessor) ChangeCurrentDirectory(dir string) {
	t.currentDirectory = dir
}

func (t *TerminalProcessor) ToTerminal() *model.Terminal {
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
