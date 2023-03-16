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

type GitDiff struct {
	Added   []FileAdd
	Updated []FileUpdate
	Deleted []FileDelete
}

func findDuplicate(operations []FileAdd) map[string]int {
	var found = make(map[string]int)
	for _, o := range operations {
		if count, ok := found[o.FilePath]; ok {
			found[o.FilePath] = count + 1
		} else {
			found[o.FilePath] = 1
		}
	}

	var duplicate = make(map[string]int)
	for k, v := range found {
		if v > 1 {
			duplicate[k] = v
		}
	}

	return duplicate
}
