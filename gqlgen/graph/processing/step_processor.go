package processing

import (
	"fmt"
	"strconv"
)

type stepProcessor struct {
	currentStep string
	nextStep    string
	prevStep    string
}

func (p *stepProcessor) CurrentStep() string {
	return p.currentStep
}

func (p *stepProcessor) NextStep() string {
	return p.nextStep
}

func (p *stepProcessor) PrevStep() string {
	return p.prevStep
}

// increment step to next
func (p *stepProcessor) IncrementStep(nextNextStep string) {
	p.prevStep = p.currentStep
	p.currentStep = p.nextStep
	p.nextStep = nextNextStep
}

func (p *stepProcessor) AutoIncrementStep() error {
	errorPreceding := "AutoIncrementStep failed"
	stepNum, err := strconv.Atoi(p.NextStep())
	if err != nil {
		return fmt.Errorf("%s, failed, as step %s is not in number format", errorPreceding, p.NextStep())
	}

	formatted := fmt.Sprintf("%03d", stepNum)
	if p.NextStep() != formatted {
		return fmt.Errorf("%s, as step %s is expected 3-digit number format %s", errorPreceding, p.NextStep(), formatted)
	}

	p.IncrementStep(formatted)

	return nil
}

func (p *stepProcessor) Clone() *stepProcessor {
	return &stepProcessor{
		currentStep: p.currentStep, // copy to avoid receiver's mutation effect afterwards
		nextStep:    p.nextStep,    // copy to avoid receiver's mutation effect afterwards
		prevStep:    p.prevStep,    // copy to avoid receiver's mutation effect afterwards
	}
}
