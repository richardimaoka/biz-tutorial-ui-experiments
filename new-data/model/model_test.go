package model

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_NewPageState(t *testing.T) {
	page := NewPageState()

	// There must be a default terminal
	if len(page.Terminals) != 1 {
		t.Errorf("terminal size = %d, not 1", len(page.Terminals))
		return
	}

	if len(page.Terminals[0].Nodes) != 0 {
		t.Errorf("terminal nodes size = %d, not 0", len(page.Terminals[0].Nodes))
		return
	}

	page.Terminals[0].Nodes = append(page.Terminals[0].Nodes, &TerminalNode{})
}

func prettyString(m map[string]interface{}) string {
	jsonString, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(jsonString)
}

func Test_InitPage(t *testing.T) {
	command := ActionCommand{
		TerminalName: "default",
		Command:      "mkdir workspace",
	}
	page := InitPage(&command)

	expectedBytes := []byte(`{
		"step":     "000",
		"nextStep": "001",
		"prevStep": null,
		"terminals": [
			{
				"currentDirectory": null,
				"name": "default",
				"nodes": [
					{
						"content": {
							"contentType": "TerminalCommand",
							"beforeExecution": true,
							"command": "mkdir workspace"
						}
					}
				]
			}
		],
		"sourceCode": null
	}`)
	var expectedMap map[string]interface{}
	err := json.Unmarshal(expectedBytes, &expectedMap)
	if err != nil {
		t.Errorf("failed to unmarshal expected json")
		return
	}

	resultBytes, err := json.Marshal(page)
	if err != nil {
		t.Errorf("failed to marshal page")
	}
	var resultMap map[string]interface{}
	err = json.Unmarshal(resultBytes, &resultMap)
	if err != nil {
		t.Errorf("failed to unmarshal result json")
		return
	}

	if !reflect.DeepEqual(expectedMap, resultMap) {
		t.Errorf("expected\n%v\nbut got\n%v", prettyString(expectedMap), prettyString(resultMap))
	}
}

func Test_NewStep(t *testing.T) {
	step := NewStep()

	if len(step.Terminals) != 1 {
		t.Errorf("terminal size = %d, not 1", len(step.Terminals))
	}

	if len(step.Terminals[0].Nodes) != 0 {
		t.Errorf("terminal nodes size = %d, not 0", len(step.Terminals[0].Nodes))
	}

	step.Terminals[0].Nodes = append(step.Terminals[0].Nodes, &TerminalNode{})
}

func Test_MarshalStep(t *testing.T) {
	stepNum := 1
	nextStepNum := 2
	terminalName := "default"

	step := Step{
		StepNum:     &stepNum,
		NextStepNum: &nextStepNum,
		Terminals: []*Terminal{
			{
				Name: &terminalName,
			},
		},
	}

	m, err := json.Marshal(step)
	if err != nil {
		t.Error(err)
	}

	got := string(m)
	want := `{"stepNum":1,"sourceCode":null,"terminals":[{"name":"default","currentDirectory":null,"nodes":null}],"nextStepNum":2,"nextAction":null}`
	if got != want {
		t.Errorf("got %s but want %s", got, want)

	}
}

// func Test_TypeInTerminalCommand2(t *testing.T) {
// 	page := NewPageState()
// 	command := ActionCommand{
// 		Command:      "mkdir protoc-go-experiments",
// 		TerminalName: "default",
// 	}

// 	err := page.typeInTerminalCommand(&command)

// 	if err != nil {
// 		t.Errorf("%s", err)
// 	}

// 	if *step.StepNum != 2 {
// 		t.Errorf("StepNum = %d, which is not incremented", *step.StepNum)
// 	}

// 	if *step.NextStepNum != 3 {
// 		t.Errorf("NextStepNum = %d, which is not incremented", *step.StepNum)
// 	}
// }

func Test_TypeInTerminalCommand(t *testing.T) {
	stepNum := 1
	nextStepNum := 2
	terminalName := "default"

	step := Step{
		StepNum:     &stepNum,
		NextStepNum: &nextStepNum,
		Terminals: []*Terminal{
			{
				Name: &terminalName,
			},
		},
	}

	err := step.typeInTerminalCommand(&ActionCommand{
		Command:      "mkdir protoc-go-experiments",
		TerminalName: "default",
	})

	if err != nil {
		t.Errorf("%s", err)
	}

	if *step.StepNum != 2 {
		t.Errorf("StepNum = %d, which is not incremented", *step.StepNum)
	}

	if *step.NextStepNum != 3 {
		t.Errorf("NextStepNum = %d, which is not incremented", *step.StepNum)
	}
}
