package model

type SourceCode struct {
	DefaultOpenFile *OpenFile           `json:"defaultOpenFile"`
	FileTree        []*FileNode         `json:"fileTree"`
	FileContents    map[string]OpenFile `json:"fileContents"`
}
