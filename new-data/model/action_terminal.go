package model

type ActionTerminal struct {
	Command          string
	TerminalName     string
	Output           string
	CurrentDirectory string
}

//no pre-condition required, always succeed
func (t *Terminal) typeIn(action ActionTerminal) {
	trueValue := true
	node := TerminalNode{
		Content: TerminalCommand{
			Command:         &action.Command,
			BeforeExecution: &trueValue,
		},
	}

	// works even if Nodes is nil
	t.Nodes = append(t.Nodes, &node)
}

//no pre-condition required, always succeed
func (t *Terminal) changeDirectory(action ActionTerminal) {
	t.CurrentDirectoryPath = &action.CurrentDirectory
}

//no pre-condition required, always succeed
func (t *Terminal) writeOutput2(action ActionTerminal) {
	node := TerminalNode{
		Content: TerminalOutput{
			Output: &action.Output,
		},
	}

	// works even if Nodes is nil
	t.Nodes = append(t.Nodes, &node)
}

// assuming typeIn() is called earlier
func (t *Terminal) execute(action ActionTerminal) {
	//action.validate()
	//t.validate(action)

	if action.Output != "" {
		t.writeOutput2(action)
	}

	if action.CurrentDirectory != "" {
		t.changeDirectory(action)
	}
}
