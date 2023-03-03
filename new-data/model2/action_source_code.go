package model2

type AddDirectory struct {
	FilePath string `json:"filePath"`
}

type DeleteDirectory struct {
	FilePath string `json:"filePath"`
}

type AddFile struct {
	FilePath      string `json:"filePath"`
	Content       string `json:"content"`
	IsFullContent bool   `json:"isFullContent"`
}

type UpdateFile struct {
	FilePath string `json:"filePath"`
	Content  string `json:"content"`
}

type DeleteFile struct {
	FilePath string `json:"filePath"`
}
