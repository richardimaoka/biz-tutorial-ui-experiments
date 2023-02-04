package model

import (
	"encoding/json"
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
