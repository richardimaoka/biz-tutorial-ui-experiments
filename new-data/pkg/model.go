package pkg

type File struct {
	TypeName    string `json:"__typename"`
	FilePath    []string
	FileContent string
	Offset      int
}

type Terminal struct {
	elements []interface{}
}

type State struct {
	SourceCode interface{}
	Terminal   interface{}
}

type ActionInfo struct {
	Action  Action
	Results []Result
}
