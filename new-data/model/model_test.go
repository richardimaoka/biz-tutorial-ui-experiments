package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

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
	fmt.Println(string(m))
}

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

	err := step.TypeInTerminalCommand(&ActionCommand{
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
