package processing

type SourceCodeOperation interface {
	IsSourceCodeOperation()
}

type SourceCodeFileOperation struct {
	FileOps []FileOperation
}

type SourceCodeGitOperation struct {
	CommitHash string
}

func (o SourceCodeFileOperation) IsSourceCodeOperation() {}
func (o SourceCodeGitOperation) IsSourceCodeOperation()  {}
