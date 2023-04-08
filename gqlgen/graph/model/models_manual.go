package model

type SourceCode struct {
	FileTree     []*FileNode         `json:"fileTree"`
	FileContents map[string]OpenFile `json:"fileContents"`
}
