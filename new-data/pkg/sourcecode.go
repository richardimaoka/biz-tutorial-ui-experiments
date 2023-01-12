package pkg

type File struct {
	TypeName    string `json:"__typename"`
	FilePath    []string
	FileContent string
	Offset      int
}

type FileTreeNode struct {
	TypeName            string `json:"__typename"`
	FilePath            []string
	Offset              int
	CurrentDirTerminals []string
}

type SourceCode struct {
	FileTree []FileTreeNode
}
