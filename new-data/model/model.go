package model

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type AddDirectory struct {
	FilePath       []string
	FilePathString string
}

type DeleteDirectory struct {
	FilePath       []string
	FilePathString string
}

type AddFile struct {
	FilePath       []string
	FilePathString string
	Content        string
}

type UpdateFile struct {
	FilePath       []string
	FilePathString string
	Content        string
}

type DeleteFile struct {
	FilePathString string
	FilePath       []string
}

type UpdateSourceCode struct {
	AddDirectories    []AddDirectory
	DeleteDirectories []DeleteDirectory
	AddFiles          []AddFile
	UpdateFiles       []UpdateFile
	DeleteFiles       []DeleteFile
}

type UpdateTerminal struct {
	Output               string //from spreadsheet
	CurrentDirectoryPath string //from spreadsheet
	CurrentDirectory     []string
}

type Action interface {
	IsAction()
}

type ActionCommand struct {
	Command          string //from spreadsheet
	TerminalName     string //from spreadsheet
	UpdateTerminal   UpdateTerminal
	UpdateSourceCode UpdateSourceCode
}

func (c *ActionCommand) IsAction() {}

type ManualUpdate struct {
	UpdateSourceCode UpdateSourceCode
}

func (c *ManualUpdate) IsAction() {}

func (o TerminalOutput) MarshalJSON() ([]byte, error) {
	extendedOutput := struct {
		ContentType string  `json:"contentType"`
		Output      *string `json:"output"`
	}{
		"TerminalOutput",
		o.Output,
	}

	return json.Marshal(extendedOutput)
}

func (ut UpdateTerminal) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})

	if ut.Output != "" {
		m["output"] = ut.Output
	}

	if len(ut.CurrentDirectory) != 0 {
		m["currentDirectory"] = ut.CurrentDirectory
	}

	return json.Marshal(m)

}

func NewPageState() *PageState {
	step := "000"
	nextStep := "001"

	//There must be a default terminal
	terminals := []*Terminal{newTerminal("default")}

	return &PageState{
		Step:       &step,
		NextStep:   &nextStep,
		Terminals:  terminals,
		SourceCode: &SourceCode{},
	}
}

func calcNextStep(stepNumString string) (string, error) {
	stepNum, err := strconv.Atoi(stepNumString)
	if err != nil {
		return "", fmt.Errorf("next step calc failed, as step %s is not number format", stepNumString)
	}

	formatted := fmt.Sprintf("%03d", stepNum)
	if stepNumString != formatted {
		return "", fmt.Errorf("next step calc failed, as step %s is expected 3-digit number format %s", stepNumString, formatted)
	}

	return fmt.Sprintf("%03d", stepNum+1), nil
}

func (p *PageState) gotoNextStep(nextNextStep string) {
	p.PrevStep = p.Step
	p.Step = p.NextStep
	p.NextStep = &nextNextStep
}

func (p *PageState) getTerminal(terminalName string) *Terminal {
	var terminal *Terminal // nil as zero value
	for _, t := range p.Terminals {
		if *t.Name == terminalName {
			terminal = t
		}
	}
	return terminal
}

func (p *PageState) typeInTerminalCommand(command *ActionCommand) error {
	// pre-condition - next step calculation
	nextNextStep, err := calcNextStep(*p.NextStep)
	if err != nil {
		return fmt.Errorf("failed to type in command, %s", err)
	}

	// pre-condition - find command's target terminal
	terminal := p.getTerminal(command.TerminalName)
	if terminal == nil {
		return fmt.Errorf("failed to type in command, terminal with name = %s not found", command.TerminalName)
	}

	// type in command
	terminal.typeInCommand(command)

	// update step
	p.gotoNextStep(nextNextStep)

	return nil
}

func (p *PageState) runTerminalCommand(command *ActionCommand) error {
	// 1.1 pre-conditions for next step

	// pre-condition - next step calculation
	nextNextStep, err := calcNextStep(*p.NextStep)
	if err != nil {
		return fmt.Errorf("failed to run command, %s", err)
	}

	// 1.2 pre-conditions for terminal

	// pre-condition - find command's target terminal
	terminal := p.getTerminal(command.TerminalName)
	if terminal == nil {
		return fmt.Errorf("failed run command, terminal with name = %s not found", command.TerminalName)
	}

	// pre-condition -terminal's last command executable
	if err := terminal.isLastCommandExecutable(); err != nil {
		return fmt.Errorf("failed run command, %s", err)
	}

	// 1.3 pre-conditions for TerminalCommand.UpdateSourceCode

	// pre-condition UpdateSourceCode.AddDirectories
	if len(command.UpdateSourceCode.AddDirectories) > 0 {
		for i, v := range command.UpdateSourceCode.AddDirectories {
			if v.FilePathString == "" {
				return fmt.Errorf("AddDirectories has %s element with empty FilePathString", ordinal(i))
			}
		}
	}

	// pre-condition UpdateSourceCode.DeleteDirectories
	if len(command.UpdateSourceCode.DeleteDirectories) > 0 {
		for i, v := range command.UpdateSourceCode.DeleteDirectories {
			if v.FilePathString == "" {
				return fmt.Errorf("DeleteDirectories has %s element with empty FilePathString", ordinal(i))
			}
		}
	}

	// pre-condition UpdateSourceCode.AddFiles
	if len(command.UpdateSourceCode.AddFiles) > 0 {
		for i, v := range command.UpdateSourceCode.AddFiles {
			if v.FilePathString == "" {
				return fmt.Errorf("AddFiles has %s element with empty FilePathString", ordinal(i))
			}
		}
	}

	// pre-condition UpdateSourceCode.UpdateFiles
	if len(command.UpdateSourceCode.UpdateFiles) > 0 {
		for i, v := range command.UpdateSourceCode.UpdateFiles {
			if v.FilePathString == "" {
				return fmt.Errorf("UpdateFiles has %s element with empty FilePathString", ordinal(i))
			}
		}
	}

	// pre-condition UpdateSourceCode.DeleteFiles
	if len(command.UpdateSourceCode.DeleteFiles) > 0 {
		for i, v := range command.UpdateSourceCode.DeleteFiles {
			if v.FilePathString == "" {
				return fmt.Errorf("UpdateFiles has %s element with empty FilePathString", ordinal(i))
			}
		}
	}

	// pre-condition AddFiles does not have a matching node in fileTree, and the parent dir
	// pre-condition UpdateFiles has matching node in fileTree
	// pre-condition DeleteFiles has matching node in fileTree

	// 2.1 Terminal update
	// p.updateTerminal(command.UpdateTerminal)

	//execute command!
	// if err := terminal.executeEffect(); err != nil {
	if err := terminal.markLastCommandExecuted(); err != nil {
		return err
	}

	// Process UpdateTerminal.Output
	if command.UpdateTerminal.Output != "" {
		terminal.Nodes = append(terminal.Nodes, &TerminalNode{
			Content: TerminalOutput{
				Output: &command.UpdateTerminal.Output,
			},
		})
	}

	// 2.2 SourceCode update

	// if len(command.UpdateTerminal.CurrentDirectory) > 0 {
	// 	terminal.CurrentDirectory = []*string{}
	// 	for _, d := range command.UpdateTerminal.CurrentDirectory {
	// 		terminal.CurrentDirectory = append(terminal.CurrentDirectory, &d)
	// 	}
	// }

	// Process UpdateSourceCode.AddDirectories
	if len(command.UpdateSourceCode.AddDirectories) > 0 {
		for _, dir := range command.UpdateSourceCode.AddDirectories {
			fileNode := dir.toFileNode()
			p.SourceCode.FileTree = append(p.SourceCode.FileTree, fileNode)
		}
	}

	//TODO: sort FileTree

	// update step
	p.gotoNextStep(nextNextStep)

	// return fmt.Errorf("runTerminalCommand() failed, terminal with name = %s not found", command.TerminalName)
	return nil
}

func (a AddDirectory) toFileNode() *FileNode {
	dType := FileNodeTypeDirectory

	split := strings.Split(a.FilePathString, "/")
	name := split[len(split)-1]

	var filePath []*string
	for _, v := range split {
		filePath = append(filePath, &v)
	}

	offset := strings.Count(a.FilePathString, "/")

	trueValue := true

	fileNode := FileNode{
		NodeType:  &dType,
		Name:      &name,
		FilePath:  filePath,
		Offset:    &offset,
		IsUpdated: &trueValue,
	}
	return &fileNode
}

func ordinal(x int) string {
	suffix := "th"
	switch x % 10 {
	case 1:
		if x%100 != 11 {
			suffix = "st"
		}
	case 2:
		if x%100 != 12 {
			suffix = "nd"
		}
	case 3:
		if x%100 != 13 {
			suffix = "rd"
		}
	}
	return strconv.Itoa(x) + suffix
}
