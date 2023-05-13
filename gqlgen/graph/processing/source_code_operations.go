package processing

type SourceCodeOperation struct {
	FileOps []FileSystemOperation
}

type SourceCodeGitOperation struct {
	CommitHash string
}

type FileSystemOperation interface {
	IsFileSystemOperation()
}

type DirectoryAdd struct {
	FilePath string
}

type DirectoryDelete struct {
	FilePath string
}

type FileAdd struct {
	FilePath      string
	Content       string
	IsFullContent bool
}

type FileUpdate struct {
	FilePath string
	Content  string
}

type FileDelete struct {
	FilePath string
}

type FileUpsert struct {
	FilePath      string
	Content       string
	IsFullContent bool
}

func (o DirectoryAdd) IsFileSystemOperation()    {}
func (o DirectoryDelete) IsFileSystemOperation() {}
func (o FileAdd) IsFileSystemOperation()         {}
func (o FileUpdate) IsFileSystemOperation()      {}
func (o FileDelete) IsFileSystemOperation()      {}
