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

func (t *TerminalProcessor) writeCommand(command string) {
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

func (t *TerminalProcessor) writeOutput(output string) {
	t.elements = append(t.elements, &terminalOutputProcessor{
		output: output,
	})
}

func (t *TerminalProcessor) changeCurrentDirectory(dir string) {
	t.currentDirectory = dir
}

//--------------------------------------------------------------------------------------------
// public methods below
//--------------------------------------------------------------------------------------------

func NewTerminalProcessor(terminalName string) *TerminalProcessor {
	return &TerminalProcessor{
		step:             "",
		terminalName:     terminalName,
		currentDirectory: "",
		elements:         []terminalElementProcessor{},
	}
}

func (t *TerminalProcessor) Transition(nextStep string, effect TerminalEffect) error {
	t.writeCommand(effect.Command)
	if effect.Output != nil {
		t.writeOutput(*effect.Output)
	}
	if effect.CurrentDirectory != nil {
		t.changeCurrentDirectory(*effect.CurrentDirectory)
	}
	return nil
}

func (t *TerminalProcessor) ToGraphQLTerminal() *model.Terminal {
	var currentDirectory *string
	if t.currentDirectory != "" {
		copied := t.currentDirectory // copy to avoid mutation effect afterwards
		currentDirectory = &copied
	}

	var terminalName *string
	if t.terminalName != "" {
		copied := t.terminalName // copy to avoid mutation effect afterwards
		terminalName = &copied
	}

	// clone to avoid mutation effect afterwards
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
	// clone to avoid mutation effect afterwards
	clonedElements := make([]terminalElementProcessor, 0)
	for _, e := range t.elements {
		clonedElements = append(clonedElements, e) //element is effectively immutable, so fine to use elsewhere
	}

	return &TerminalProcessor{
		terminalName:     t.terminalName,
		currentDirectory: t.currentDirectory,
		elements:         clonedElements,
	}
}
