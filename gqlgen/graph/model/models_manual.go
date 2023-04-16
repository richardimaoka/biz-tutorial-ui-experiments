package model

type SourceCode struct {
	Step            *string             `json:"step"` //This is not exposed in the schema, so that SourceCode needs to be a manual type
	DefaultOpenFile *OpenFile           `json:"defaultOpenFile"`
	FileTree        []*FileNode         `json:"fileTree"`
	FileContents    map[string]OpenFile `json:"fileContents"` //This is not exposed in the schema, so that SourceCode needs to be a manual type
}
