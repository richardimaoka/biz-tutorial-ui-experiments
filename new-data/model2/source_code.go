package model2

type SourceCodeExtended struct {
	SourceCode
	FileContents map[string]OpenFile `json:"fileContents"`
}

func newSourceCode() *SourceCodeExtended {
	return &SourceCodeExtended{}
}
