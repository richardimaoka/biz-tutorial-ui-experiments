package model

import (
	"encoding/json"
	"reflect"
	"testing"
)

func prettyString(m map[string]interface{}) string {
	jsonString, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(jsonString)
}

func compareJsonBytes(t *testing.T, expectedBytes, resultBytes []byte) {
	var resultMap map[string]interface{}
	err := json.Unmarshal(resultBytes, &resultMap)
	if err != nil {
		t.Errorf("failed to unmarshal result json")
		return
	}

	var expectedMap map[string]interface{}
	err = json.Unmarshal(expectedBytes, &expectedMap)
	if err != nil {
		t.Errorf("failed to unmarshal expected json")
		return
	}

	if !reflect.DeepEqual(expectedMap, resultMap) {
		t.Errorf("expected\n%v\nbut got\n%v", prettyString(expectedMap), prettyString(resultMap))
	}
}

func compareAfterMarshal(t *testing.T, expectedBytes []byte, result interface{}) {
	resultBytes, err := json.Marshal(result)
	if err != nil {
		t.Errorf("failed to marshal result")
		return
	}

	compareJsonBytes(t, expectedBytes, resultBytes)
}

func Test_NewPageState(t *testing.T) {
	result := NewPageState()
	expectedBytes := []byte(`{
		"step":     "000",
		"nextStep": "001",
		"prevStep": null,
		"terminals": [
			{
				"currentDirectory": null,
				"currentDirectoryPath": null,
				"name": "default", 
				"nodes" : null
			}
		],
		"sourceCode": null
	}`)

	compareAfterMarshal(t, expectedBytes, result)
}

func Test_typeInCommand(t *testing.T) {
	cmd := ActionCommand{Command: "mkdir abc", TerminalName: "default"}
	result := NewPageState()
	result.typeInTerminalCommand(&cmd)

	expectedBytes := []byte(`{
		"step":     "001",
		"nextStep": "002",
		"prevStep": "000",
		"terminals": [
			{
				"currentDirectory": null,
				"currentDirectoryPath": null,
				"name": "default", 
				"nodes" : [
					{
						"content": {
							"contentType": "TerminalCommand",
							"beforeExecution": true,
							"command": "mkdir abc"
						}
     		  }
				]
			}
		],
		"sourceCode": null
	}`)

	compareAfterMarshal(t, expectedBytes, result)
}

func Test_calcNextStep(t *testing.T) {
	cases := []struct {
		Description string
		CurrentStep string
		Expected    string
	}{
		{Description: "next to 000 is 001", CurrentStep: "000", Expected: "001"},
		{Description: "next to 001 is 002", CurrentStep: "001", Expected: "002"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			result, err := calcNextStep(test.CurrentStep)
			if err != nil {
				t.Errorf("failed to calc next step for %s, %s", result, err)
			}

			if result != test.Expected {
				t.Errorf("expected %s, but result %s", test.Expected, result)
			}
		})
	}
}

func Test_InitPage(t *testing.T) {
	command := ActionCommand{
		TerminalName: "default",
		Command:      "mkdir workspace",
	}

	page := InitPage(&command)

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

	expectedBytes := []byte(`{
		"step":     "000",
		"nextStep": "001",
		"prevStep": null,
		"terminals": [
			{
				"currentDirectory": null,
				"currentDirectoryPath": null,
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
	err = json.Unmarshal(expectedBytes, &expectedMap)
	if err != nil {
		t.Errorf("failed to unmarshal expected json")
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
	want := `{"stepNum":1,"sourceCode":null,"terminals":[{"name":"default","currentDirectory":null,"currentDirectoryPath":null,"nodes":null}],"nextStepNum":2,"nextAction":null}`
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
