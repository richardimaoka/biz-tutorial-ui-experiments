package preprocess2

type SourceCodeStep struct {
	ID         string
	commitHash string
}

type SourceCodeTransition interface {
}

//differentiate git->git transition from manual->git transition? highlighting would be affected
type GitTransition struct {
}

type ManualTransition struct {
}

type DirectoryNode struct {
	dirs     []DirectoryNode
	files    []FileNode
	treeHash string //assuming every transition is git
}

type FileNode struct {
	updatedStepID string
	blobHash      string //blobHash is equivalent to content, assuming every transition is git
	// ref        FileContentRef //when manual edit is allowed, blobHash is not enought to indicate the contnet
}

func (n *FileNode) IsUpdated(step SourceCodeStep) bool {
	return step.ID == n.updatedStepID
}

func (n *FileNode) Diff(from SourceCodeStep) string {
	return ""
}

type SourceCode struct {
	RootDir DirectoryNode
}
