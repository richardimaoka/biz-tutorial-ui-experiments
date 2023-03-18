package model

type SourceCode struct {
	FileTree     []*FileNode         `json:"fileTree"`
	OpenFile     *OpenFile           `json:"openFile"`
	FileContents map[string]OpenFile `json:"fileContents"`
}
