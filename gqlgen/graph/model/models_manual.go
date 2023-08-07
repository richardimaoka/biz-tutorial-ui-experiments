package model

type SourceCode struct {
	Step           string      `json:"step"` //This is not exposed in the schema, so that SourceCode needs to be a manual type
	FileTree       []*FileNode `json:"fileTree"`
	ProjectDir     string      `json:"projectDir"`
	IsFoldFileTree bool        `json:"isFoldFileTree"`

	// These are not exposed in the schema, so that SourceCode needs to be a manual type
	Tutorial            string              `json:"tutorial"`
	FileContents        map[string]OpenFile `json:"fileContents"`
	DefaultOpenFilePath string              `json:"defaultOpenFilePath"`
}
