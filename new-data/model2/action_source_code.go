package model2

type FileSystemOperation interface {
	IsFileSystemOperation()
}

type DirectoryAdd struct {
	FilePath string `json:"filePath"`
}

func (o DirectoryAdd) IsFileSystemOperation() {}

type DirectoryDelete struct {
	FilePath string `json:"filePath"`
}

func (o DirectoryDelete) IsFileSystemOperation() {}

type FileAdd struct {
	FilePath      string `json:"filePath"`
	Content       string `json:"content"`
	IsFullContent bool   `json:"isFullContent"`
}

func (o FileAdd) IsFileSystemOperation() {}

type FileUpdate struct {
	FilePath string `json:"filePath"`
	Content  string `json:"content"`
}

func (o FileUpdate) IsFileSystemOperation() {}

type FileDelete struct {
	FilePath string `json:"filePath"`
}

func (o FileDelete) IsFileSystemOperation() {}

type SourceCodeEffect struct {
	DirectoriesToAdd    []DirectoryAdd
	DirectoriesToDelete []DirectoryDelete
	FilesToAdd          []FileAdd
	FilesToUpdate       []FileUpdate
	FilesToDelete       []FileDelete
}
