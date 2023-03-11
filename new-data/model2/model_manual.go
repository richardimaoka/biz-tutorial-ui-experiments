package model2

type SourceCode struct {
	FileTree     []*FileNode         `json:"fileTree"`
	OpenFile     *OpenFile           `json:"openFile"`
	FileContents map[string]OpenFile `json:"fileContents"`
}
