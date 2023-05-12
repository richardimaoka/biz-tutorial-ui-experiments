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
	t.writeCommandWithPrompt(defaultPromptExpression, defaultPromptSymbol, command)
}

func (t *TerminalProcessor) writeCommandWithPrompt(promptExpression string, promptSymbol rune, command string) {
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

func (t *TerminalProcessor) Transition(nextStep string, effect TerminalEffect) {
	t.writeCommand(effect.Command)
	if effect.Output != nil {
		t.writeOutput(*effect.Output)
	}
	if effect.CurrentDirectory != nil {
		t.changeCurrentDirectory(*effect.CurrentDirectory)
	}

	t.step = nextStep
}

func (t *TerminalProcessor) TransitionWithOperation(nextStep string, op TerminalOperation) {
	switch v := op.(type) {
	case TerminalCommand:
		t.writeCommand(v.Command)
	case TerminalCommandWithOutput:
		t.writeCommand(v.Command)
		t.writeOutput(v.Output)
	case TerminalCommandWithCd:
		t.writeCommand(v.Command)
		t.changeCurrentDirectory(v.CurrentDirectory)
	case TerminalCommandWithOutputCd:
		t.writeCommand(v.Command)
		t.writeOutput(v.Output)
		t.changeCurrentDirectory(v.CurrentDirectory)
	}

	t.step = nextStep
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
