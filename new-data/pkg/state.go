package pkg

type State struct {
	SourceCode SourceCode
	Terminal   Terminal
}

func InitialState() State {
	return State{InitialSourceCode(), InitialTerminal()}
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

// func transition(prevState *State, action *Action) (*State, *State) {
// 	stateBeforeAction := action.stateBeforeAction(prevState)
// 	stateAfterAction := action.stateAfterAction(stateBeforeAction)

// 	return stateBeforeAction, stateAfterAction
// }
