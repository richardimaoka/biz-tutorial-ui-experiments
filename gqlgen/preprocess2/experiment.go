package preprocess2

type SourceCodeTransition interface {
}

//differentiate git->git transition from manual->git transition? highlighting would be affected
type GitTransition struct {
}

// type ManualTransition struct {
// }
