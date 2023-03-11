package model2

type DirectoryAdd struct {
	FilePath string `json:"filePath"`
}

type DirectoryDelete struct {
	FilePath string `json:"filePath"`
}

type FileAdd struct {
	FilePath      string `json:"filePath"`
	Content       string `json:"content"`
	IsFullContent bool   `json:"isFullContent"`
}

type FileUpdate struct {
	FilePath string `json:"filePath"`
	Content  string `json:"content"`
}

type FileDelete struct {
	FilePath string `json:"filePath"`
}

type SourceCodeEffect struct {
	DirectoriesToAdd    []DirectoryAdd
	DirectoriesToDelete []DirectoryDelete
	FilesToAdd          []FileAdd
	FilesToUpdate       []FileUpdate
	FilesToDelete       []FileDelete
}
