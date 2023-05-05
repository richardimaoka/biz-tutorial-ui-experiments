package processing

type TerminalEffect struct {
	TerminalName     string  `json:"terminalName"`
	Command          string  `json:"command"`
	Output           *string `json:"output"`           //if zero value, no output after execution
	CurrentDirectory *string `json:"currentDirectory"` //if zero value, current directory is not changed after execution
}

type SourceCodeEffect struct {
	Diff                Diff    `json:"diff"`
	DefaultOpenFilePath *string `json:"defaultOpenFilePath"`
}

// TODO: later optimization
// type SourceCodeGitEffect struct {
// 	commitHash string
// }

type Action2 struct {
	Step             string
	SourceCodeEffect *SourceCodeEffect `json:"sourceCodeEffect"`
	TerminalEffect   *TerminalEffect   `json:"terminalEffect"`
}
