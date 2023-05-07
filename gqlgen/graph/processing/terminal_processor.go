package processing

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type TerminalProcessor struct {
	step             string
	terminalName     string
	currentDirectory string
	elements         []terminalElementProcessor
}

func NewTerminalProcessor(terminalName string) *TerminalProcessor {
	return &TerminalProcessor{
		step:             "",
		terminalName:     terminalName,
		currentDirectory: "",
		elements:         []terminalElementProcessor{},
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

func (t *TerminalProcessor) Transition(nextStep string, effect TerminalEffect) error {
	t.WriteCommand(effect.Command)
	if effect.Output != nil {
		t.WriteOutput(*effect.Output)
	}
	if effect.CurrentDirectory != nil {
		t.ChangeCurrentDirectory(*effect.CurrentDirectory)
	}
	return nil
}

func (t *TerminalProcessor) ToGraphQLTerminal() *model.Terminal {
	var currentDirectory *string
	if t.currentDirectory != "" {
		copied := t.currentDirectory // copy to avoid receiver's mutation effect afterwards
		currentDirectory = &copied
	}

	var terminalName *string
	if t.terminalName != "" {
		copied := t.terminalName // copy to avoid receiver's mutation effect afterwards
		terminalName = &copied
	}

	// clone to avoid receiver's mutation effect afterwards
	var nodes []*model.TerminalNode
	for _, e := range t.elements {
		nodes = append(nodes, &model.TerminalNode{
			Content: e.ToGraphQLModel(),
		})
	}

	return &model.Terminal{
		Name:             terminalName,
		CurrentDirectory: currentDirectory,
		Nodes:            nodes,
	}
}

func (t *TerminalProcessor) Clone() *TerminalProcessor {
	// clone to avoid receiver's mutation effect afterwards
	clonedElements := make([]terminalElementProcessor, 0)
	for _, e := range t.elements {
		clonedElements = append(clonedElements, e) //element is effectively immutable, so fine to use elsewhere
	}

	return &TerminalProcessor{
		terminalName:     t.terminalName,     // copy to avoid receiver's mutation effect afterwards
		currentDirectory: t.currentDirectory, // copy to avoid receiver's mutation effect afterwards
		elements:         clonedElements,
	}
}
