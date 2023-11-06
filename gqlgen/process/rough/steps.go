package rough

type ResultStep struct {
	Step         string `json:"step"`
	Phase        string `json:"phase"`
	Type         string `json:"type"`
	Instruction  string `json:"instruction"`
	Instruction2 string `json:"instruction2"`
	Instruction3 string `json:"instruction3"`
	ModalText    string `json:"modalText"`
	Tooltip      string `json:"tooltip"`
	Commit       string `json:"commit"`
	Comment      string `json:"comment"`
}

type TerminalCommandStep struct {
	Comment string
	Command string
	Commit  string
	Output  string
	Tooltip string
}

type TerminalOutputStep struct {
	Comment string
	Output  string
	Tooltip string
}

type SourceCodeCommitStep struct {
	Comment string
	Commit  string
	Tooltip string
}

type BrowserStep struct {
	Comment string
	Commit  string
	Tooltip string
}

func toTerminalCommandStep(roughStep ResultStep) TerminalCommandStep {
	return TerminalCommandStep{
		Comment: roughStep.Comment,
		Command: roughStep.Instruction,
		Commit:  roughStep.Commit,
		Output:  roughStep.Instruction2,
		Tooltip: roughStep.Tooltip,
	}
}
