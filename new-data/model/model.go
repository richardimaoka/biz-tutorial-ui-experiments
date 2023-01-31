package model

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type AddDirectory struct {
	FilePath []string
}

type DeleteDirectory struct {
	FilePath []string
}

type AddFile struct {
	FilePath []string
	Content  string
}

type UpdateFile struct {
	FilePath []string
	Content  string
}

type DeleteFile struct {
	FilePath []string
}

type UpdateSourceCode struct {
	AddDirectories    []AddDirectory
	DeleteDirectories []DeleteDirectory
	AddFiles          []AddFile
	UpdateFiles       []UpdateFile
	DeleteFiles       []DeleteFile
}

type UpdateTerminal struct {
	Output           string
	CurrentDirectory []string
}

type ActionCommand struct {
	ActionType       string
	Command          string
	TerminalName     string
	UpdateTerminal   UpdateTerminal
	UpdateSourceCode UpdateSourceCode
}

func extractTypeName(bytes []byte, fromField string) (string, error) {
	var unmarshaled map[string]interface{}
	if err := json.Unmarshal(bytes, &unmarshaled); err != nil {
		return "", err
	}

	typeNameRaw, ok := unmarshaled[fromField]
	if !ok {
		return "", fmt.Errorf("\"%s\" does not exist in JSON", fromField)
	}

	typeName, ok := typeNameRaw.(string)
	if !ok {
		return "", fmt.Errorf("\"%s\" is not a string, but found in type = %v", fromField, reflect.TypeOf(typeNameRaw))
	}

	return typeName, nil
}

func readAction(filePath string) (*ActionCommand, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("readAction() failed, %s", err)
	}

	typeName, err := extractTypeName(bytes, "actionType")
	if err != nil {
		return nil, fmt.Errorf("readAction() failed to actionType from %s, %s", filePath, err)
	}

	switch typeName {
	case "ActionCommand":
		var action ActionCommand
		err := json.Unmarshal(bytes, &action)
		if err != nil {
			return nil, err
		}
		return &action, nil
	default:
		return nil, fmt.Errorf("readAction() found invalid typeName = %s in file = %s", typeName, filePath)
	}
}

func (node TerminalNode) MarshalJSON() ([]byte, error) {
	type TerminalNodeExtended struct {
		ContentType string `json:"contentType"`
		TerminalNode
	}

	switch content := node.Content.(type) {
	case TerminalCommand:
		typedNode := struct {
			ContentType string
			Content     TerminalCommand
		}{
			"TerminalCommand",
			content,
		}
		return json.Marshal(typedNode)
	default:
		fmt.Println("(t TerminalNode) MarshalJSON()")
		return nil, fmt.Errorf("default is error")
	}
}

func (step *Step) TypeInTerminalCommand(command *ActionCommand) error {
	for _, v := range step.Terminals {
		if *v.Name == command.TerminalName {
			falseValue := false
			v.Nodes = append(v.Nodes, &TerminalNode{
				Content: TerminalCommand{
					Command:         &command.Command,
					BeforeExecution: &falseValue,
				},
			})
			*step.StepNum++
			*step.NextStepNum++
			return nil
		}
	}

	return fmt.Errorf("TypeInTerminalCommand() failed, terminal with name = %s not found", command.TerminalName)
}

func (step *Step) RunTerminalCommand(command *ActionCommand) error {
	for _, v := range step.Terminals {
		if v.Name == &command.TerminalName {
			v.Nodes = append(v.Nodes, &TerminalNode{
				Content: TerminalCommand{
					Command: &command.Command,
				},
			})
			*step.StepNum++
			*step.NextStepNum++
			return nil
		}
	}

	return fmt.Errorf("TypeInTerminalCommand() failed, terminal with name = %s not found", command.TerminalName)
}

func convertActionCommand(command *ActionCommand, step int) (int, error) {
	nextStep := step + 1
	nextNextStep := step + 1
	terminalName := "default"
	trueValue := true
	falseValue := false

	stepBeforeStruct := Step{
		StepNum:     &step,
		NextStepNum: &nextStep,
		Terminals: []*Terminal{
			{
				Name:             &terminalName,
				CurrentDirectory: []*string{&terminalName},
				Nodes: []*TerminalNode{
					{
						Content: TerminalCommand{
							Command:         &command.Command,
							BeforeExecution: &trueValue,
						},
					},
				},
			},
		},
	}

	stepAfterStruct := Step{
		StepNum:     &nextStep,
		NextStepNum: &nextNextStep,
		Terminals: []*Terminal{
			{
				Name:             &terminalName,
				CurrentDirectory: []*string{&terminalName},
				Nodes: []*TerminalNode{
					{
						// ContentType: &terminalCommandType,
						Content: TerminalCommand{
							Command:         &command.Command,
							BeforeExecution: &falseValue,
						},
					},
				},
			},
		},
		SourceCode: &SourceCode{},
	}

	fmt.Println(stepBeforeStruct.StepNum)
	fmt.Println(stepAfterStruct.StepNum)
	s, err := json.Marshal(stepAfterStruct)
	if err != nil {
		return 0, err
	}
	fmt.Println(string(s))

	return nextNextStep, nil
}

func Process() error {
	var step = 0
	var actionFile = fmt.Sprintf("data2/action%02d.json", step)

	action, err := readAction(actionFile)
	if err != nil {
		return err
	}
	convertActionCommand(action, step)
	return nil
}
