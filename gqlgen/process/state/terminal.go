package state

import (
	"fmt"
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
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

func (t *Terminal) WriteCommand(command string) {
	t.elements = append(t.elements, &TerminalCommand{Command: command})
}

func (t *Terminal) WriteOutput(output string) {
	t.elements = append(t.elements, &TerminalOutput{Output: output})
}

func (t *Terminal) ChangeCurrentDirectory(dirPath string) {
	t.currentDirectory = dirPath
}

func (t *Terminal) ToGraphQLTerminal() *model.Terminal {
	// copy to avoid mutation effect afterwards
	currentDirectory := internal.StringRef(t.currentDirectory)
	terminalName := internal.StringRef(t.terminalName)

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
