package model

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
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

func extractTypeName(jsonBytes []byte, fromField string) (string, error) {
	var unmarshaled map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &unmarshaled); err != nil {
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

func readActionFromBytes(bytes []byte) (*ActionCommand, error) {
	typeName, err := extractTypeName(bytes, "actionType")
	if err != nil {
		return nil, fmt.Errorf("readActionFromBytes() failed to extract actionType %s", err)
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
		return nil, fmt.Errorf("readActionFromBytes() found invalid typeName = %s", typeName)
	}
}

func (node TerminalNode) MarshalJSON() ([]byte, error) {
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
		return nil, fmt.Errorf("default is error")
	}
}

func (ut UpdateTerminal) MarshalJSON() ([]byte, error) {
	fmt.Println("UpdateTerminal MarshalJSON")
	m := make(map[string]interface{})

	if ut.Output != "" {
		m["output"] = ut.Output
	}

	if len(ut.CurrentDirectory) != 0 {
		m["currentDirectory"] = ut.CurrentDirectory
	}

	return json.Marshal(m)

}

func NewStep() *Step {
	stepNum := 0
	nextStepNum := 1
	terminalName := "default"

	return &Step{
		StepNum:     &stepNum,
		NextStepNum: &nextStepNum,
		Terminals: []*Terminal{
			{
				Name: &terminalName,
			},
		},
		SourceCode: &SourceCode{},
	}
}

func NewPageState() *PageState {
	step := "000"
	nextStep := "001"

	return &PageState{
		Step:     &step,
		NextStep: &nextStep,
	}
}

func (step *Step) typeInTerminalCommand(command *ActionCommand) error {
	for _, t := range step.Terminals {
		if *t.Name == command.TerminalName {
			*step.StepNum++
			*step.NextStepNum++

			falseValue := false
			t.Nodes = append(t.Nodes, &TerminalNode{
				Content: TerminalCommand{
					Command:         &command.Command,
					BeforeExecution: &falseValue,
				},
			})

			return nil
		}
	}

	return fmt.Errorf("typeInTerminalCommand() failed, terminal with name = %s not found", command.TerminalName)
}

func (step *Step) runTerminalCommand(command *ActionCommand) error {
	for _, t := range step.Terminals {
		if *t.Name == command.TerminalName {
			*step.StepNum++
			*step.NextStepNum++

			// Process UpdateTerminal
			if command.UpdateTerminal.Output != "" {
				t.Nodes = append(t.Nodes, &TerminalNode{
					Content: TerminalOutput{
						Output: &command.Command,
					},
				})
			}

			if len(command.UpdateTerminal.CurrentDirectory) > 0 {
				t.CurrentDirectory = []*string{}
				for _, d := range command.UpdateTerminal.CurrentDirectory {
					t.CurrentDirectory = append(t.CurrentDirectory, &d)
				}
			}

			// Process UpdateSourceCode
			if len(command.UpdateSourceCode.AddDirectories) > 0 {
				for i, d := range command.UpdateSourceCode.AddDirectories {
					if len(d.FilePath) == 0 {
						return fmt.Errorf("AddDirectories has %s element with empty filePath", ordinal(i))
					}

					dType := FileNodeTypeDirectory
					offset := len(d.FilePath)
					trueValue := true
					dirName := d.FilePath[len(d.FilePath)-1]

					filePath := []*string{}
					for _, p := range d.FilePath {
						filePath = append(filePath, &p)
					}

					step.SourceCode.FileTree = append(
						step.SourceCode.FileTree,
						&FileNode{
							NodeType:  &dType,
							Name:      &dirName,
							FilePath:  filePath,
							Offset:    &offset,
							IsUpdated: &trueValue,
						},
					)
				}
			}

			//TODO: sort FileTree

			return nil
		}
	}

	return fmt.Errorf("runTerminalCommand() failed, terminal with name = %s not found", command.TerminalName)
}

func (step *Step) ProcessActionCommand(command *ActionCommand, filePrefix string) error {
	//step 1 type in terminal command
	err := step.typeInTerminalCommand(command)
	if err != nil {
		return err
	}

	err = step.writeFile(filePrefix)
	if err != nil {
		return fmt.Errorf("ProcessActionCommand() failed, %s", err)
	}

	//step 2 run terminal command
	err = step.runTerminalCommand(command)
	if err != nil {
		return err
	}

	err = step.writeFile(filePrefix)
	if err != nil {
		return fmt.Errorf("ProcessActionCommand() failed, %s", err)
	}

	return nil
}

func (s *Step) writeFile(filePrefix string) error {
	bytes, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("writeFile() failed, %s", err)
	}

	fileName := fmt.Sprintf("%s%03d.json", filePrefix, *s.StepNum)
	if os.WriteFile(fileName, bytes, 0644) != nil {
		return fmt.Errorf("writeFile() failed, %s", err)
	}

	return nil
}

func Process() error {
	filePrefix := "data2/step"

	step := NewStep()
	step.writeFile(filePrefix)

	for i := 0; i <= 100; i++ {
		actionFile := fmt.Sprintf("data2/action%03d.json", i)
		bytes, err := os.ReadFile(actionFile)
		if err != nil {
			return nil //no file found, end of processing
		}

		cmd, err := readActionFromBytes(bytes)
		if err != nil {
			return fmt.Errorf("failure in %s, %s", actionFile, err)
		}

		err = step.ProcessActionCommand(cmd, filePrefix)
		if err != nil {
			return err
		}
	}

	return fmt.Errorf("loop count 100, too big!!")
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
