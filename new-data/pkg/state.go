package pkg

type State struct {
	SourceCode SourceCode
	Terminal   Terminal
}

func StatesFromCommand(cmd *Command, state *State) (*State, *State) {
	stateBefore := State{
		SourceCode{FileTree: state.SourceCode.FileTree},
		Terminal{Elements: state.Terminal.Elements},
	}

	stateAfter := State{
		SourceCode{FileTree: state.SourceCode.FileTree},
		Terminal{Elements: state.Terminal.Elements},
	}

	return &stateBefore, &stateAfter
}
