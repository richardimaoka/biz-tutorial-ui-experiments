package model

import "testing"

func TestHello(t *testing.T) {

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
