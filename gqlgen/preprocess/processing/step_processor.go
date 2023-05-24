package processing

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

func NewStepProcessor() *stepProcessor {
	return &stepProcessor{
		currentStep: "_initial",
		nextStep:    "",
		prevStep:    "",
	}
}

// increment step to next
func (p *stepProcessor) IncrementStep(nextNextStep string) {
	p.prevStep = p.currentStep
	p.currentStep = p.nextStep
	p.nextStep = nextNextStep
}

func (p *stepProcessor) Clone() *stepProcessor {
	return &stepProcessor{
		currentStep: p.currentStep, // copy to avoid receiver's mutation effect afterwards
		nextStep:    p.nextStep,    // copy to avoid receiver's mutation effect afterwards
		prevStep:    p.prevStep,    // copy to avoid receiver's mutation effect afterwards
	}
}
