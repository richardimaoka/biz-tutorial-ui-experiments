package state

import (
	"fmt"
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type Terminal struct {
	terminalName     string
	currentDirectory string
	elements         []TerminalElement
}

type TerminalElementType string

const (
	TerminalTypeCommand TerminalElementType = "COMMAND"
	TerminalTypeOutput  TerminalElementType = "OUTPUT"
)

func ToTerminalElementType(t string) (TerminalElementType, error) {
	switch strings.ToUpper(t) {
	case "COMMAND":
		return TerminalTypeCommand, nil
	case "OUTPUT":
		return TerminalTypeOutput, nil
	default:
		return "", fmt.Errorf("'%s' is unknown TerminalElementType", t)
	}
}

func NewTerminal() *Terminal {
	return &Terminal{}
}

func (t *Terminal) MarkAllExecuted() {
	for _, e := range t.elements {
		if c, ok := e.(*TerminalCommand); ok {
			c.BeforeExecution = false
			c.Tooltip = ""
		}
		if o, ok := e.(*TerminalOutput); ok {
			o.Tooltip = ""
		}
	}
}

func (t *Terminal) WriteCommand(command string) {
	t.MarkAllExecuted()
	t.elements = append(t.elements, &TerminalCommand{Command: command, BeforeExecution: true})
}

func (t *Terminal) WriteOutput(output string) {
	t.MarkAllExecuted()
	t.elements = append(t.elements, &TerminalOutput{Output: output})
}

func (t *Terminal) ChangeCurrentDirectory(dirPath string) {
	t.currentDirectory = dirPath
}

func (t *Terminal) ToGraphQLTerminal() *model.Terminal {
	// copy to avoid mutation effect afterwards
	currentDirectory := stringRef(t.currentDirectory)
	terminalName := stringRef(t.terminalName)

	var nodes []*model.TerminalNode
	for _, e := range t.elements {
		nodes = append(nodes, &model.TerminalNode{
			Content: e.ToGraphQLTerminalElement(),
		})
	}

	return &model.Terminal{
		Name:             terminalName,
		CurrentDirectory: currentDirectory,
		Nodes:            nodes,
	}
}
